// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// EthAPIBackend implements ethapi.Backend for full nodes
type EthAPIBackend struct {
	eth *Ethereum
	gpo *gasprice.Oracle

	// Quorum
	//
	// hex node id from node public key
	hexNodeId string
}

// ChainConfig returns the active chain configuration.
func (b *EthAPIBackend) ChainConfig() *params.ChainConfig {
	return b.eth.chainConfig
}

func (b *EthAPIBackend) CurrentBlock() *types.Block {
	return b.eth.blockchain.CurrentBlock()
}

func (b *EthAPIBackend) SetHead(number uint64) {
	b.eth.protocolManager.downloader.Cancel()
	b.eth.blockchain.SetHead(number)
}

func (b *EthAPIBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block := b.eth.miner.PendingBlock()
		return block.Header(), nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock().Header(), nil
	}
	return b.eth.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
}

func (b *EthAPIBackend) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return b.eth.blockchain.GetHeaderByHash(hash), nil
}

func (b *EthAPIBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		if b.eth.protocolManager.raftMode {
			// Use latest instead.
			return b.eth.blockchain.CurrentBlock(), nil
		}
		block := b.eth.miner.PendingBlock()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock(), nil
	}
	return b.eth.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EthAPIBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (vm.MinimalApiState, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		if b.eth.protocolManager.raftMode {
			// Use latest instead.
			header, err := b.HeaderByNumber(ctx, rpc.LatestBlockNumber)
			if header == nil || err != nil {
				return nil, nil, err
			}
			publicState, privateState, err := b.eth.BlockChain().StateAt(header.Root)
			return EthAPIState{publicState, privateState}, header, err
		}
		block, publicState, privateState := b.eth.miner.Pending()
		return EthAPIState{publicState, privateState}, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header, err := b.HeaderByNumber(ctx, blockNr)
	if header == nil || err != nil {
		return nil, nil, err
	}
	stateDb, privateState, err := b.eth.BlockChain().StateAt(header.Root)
	return EthAPIState{stateDb, privateState}, header, err
}

func (b *EthAPIBackend) GetBlock(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return b.eth.blockchain.GetBlockByHash(hash), nil
}

func (b *EthAPIBackend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	return b.eth.blockchain.GetReceiptsByHash(hash), nil
}

func (b *EthAPIBackend) GetLogs(ctx context.Context, hash common.Hash) ([][]*types.Log, error) {
	receipts := b.eth.blockchain.GetReceiptsByHash(hash)
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *EthAPIBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.eth.blockchain.GetTdByHash(blockHash)
}

func (b *EthAPIBackend) GetEVM(ctx context.Context, msg core.Message, state vm.MinimalApiState, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	statedb := state.(EthAPIState)
	from := statedb.state.GetOrNewStateObject(msg.From())
	from.SetBalance(math.MaxBig256)
	vmError := func() error { return nil }

	context := core.NewEVMContext(msg, header, b.eth.BlockChain(), nil)

	// Set the private state to public state if contract address is not present in the private state
	to := common.Address{}
	if msg.To() != nil {
		to = *msg.To()
	}

	privateState := statedb.privateState
	if !privateState.Exist(to) {
		privateState = statedb.state
	}

	return vm.NewEVM(context, statedb.state, privateState, b.eth.chainConfig, vmCfg), vmError, nil
}

func (b *EthAPIBackend) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeRemovedLogsEvent(ch)
}

func (b *EthAPIBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainEvent(ch)
}

func (b *EthAPIBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainHeadEvent(ch)
}

func (b *EthAPIBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	return b.eth.BlockChain().SubscribeChainSideEvent(ch)
}

func (b *EthAPIBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return b.eth.BlockChain().SubscribeLogsEvent(ch)
}

func (b *EthAPIBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	// validation for node need to happen here and cannot be done as a part of
	// validateTx in tx_pool.go as tx_pool validation will happen in every node
	if b.hexNodeId != "" && !types.ValidateNodeForTxn(b.hexNodeId, signedTx.From()) {
		return errors.New("cannot send transaction from this node")
	}
	return b.eth.txPool.AddLocal(signedTx)
}

func (b *EthAPIBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.eth.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EthAPIBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.eth.txPool.Get(hash)
}

func (b *EthAPIBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.eth.txPool.State().GetNonce(addr), nil
}

func (b *EthAPIBackend) Stats() (pending int, queued int) {
	return b.eth.txPool.Stats()
}

func (b *EthAPIBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.eth.TxPool().Content()
}

func (b *EthAPIBackend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.eth.TxPool().SubscribeNewTxsEvent(ch)
}

func (b *EthAPIBackend) Downloader() *downloader.Downloader {
	return b.eth.Downloader()
}

func (b *EthAPIBackend) ProtocolVersion() int {
	return b.eth.EthVersion()
}

func (b *EthAPIBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	if b.ChainConfig().EnableGasPrice { //Quorum
		return b.gpo.SuggestPrice(ctx)
	} else {
		return big.NewInt(0), nil
	}
}

func (b *EthAPIBackend) ChainDb() ethdb.Database {
	return b.eth.ChainDb()
}

func (b *EthAPIBackend) EventMux() *event.TypeMux {
	return b.eth.EventMux()
}

func (b *EthAPIBackend) AccountManager() *accounts.Manager {
	return b.eth.AccountManager()
}

func (b *EthAPIBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.eth.bloomIndexer.Sections()
	return params.BloomBitsBlocks, sections
}

func (b *EthAPIBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.eth.bloomRequests)
	}
}

// used by Quorum
type EthAPIState struct {
	state, privateState *state.StateDB
}

func (s EthAPIState) GetBalance(addr common.Address) *big.Int {
	if s.privateState.Exist(addr) {
		return s.privateState.GetBalance(addr)
	}
	return s.state.GetBalance(addr)
}

func (s EthAPIState) GetCode(addr common.Address) []byte {
	if s.privateState.Exist(addr) {
		return s.privateState.GetCode(addr)
	}
	return s.state.GetCode(addr)
}

func (s EthAPIState) GetState(a common.Address, b common.Hash) common.Hash {
	if s.privateState.Exist(a) {
		return s.privateState.GetState(a, b)
	}
	return s.state.GetState(a, b)
}

func (s EthAPIState) GetNonce(addr common.Address) uint64 {
	if s.privateState.Exist(addr) {
		return s.privateState.GetNonce(addr)
	}
	return s.state.GetNonce(addr)
}

func (s EthAPIState) GetProof(addr common.Address) ([][]byte, error) {
	if s.privateState.Exist(addr) {
		return s.privateState.GetProof(addr)
	}
	return s.state.GetProof(addr)
}

func (s EthAPIState) GetStorageProof(addr common.Address, h common.Hash) ([][]byte, error) {
	if s.privateState.Exist(addr) {
		return s.privateState.GetStorageProof(addr, h)
	}
	return s.state.GetStorageProof(addr, h)
}

func (s EthAPIState) StorageTrie(addr common.Address) state.Trie {
	if s.privateState.Exist(addr) {
		return s.privateState.StorageTrie(addr)
	}
	return s.state.StorageTrie(addr)
}

func (s EthAPIState) Error() error {
	if s.privateState.Error() != nil {
		return s.privateState.Error()
	}
	return s.state.Error()
}

func (s EthAPIState) GetCodeHash(addr common.Address) common.Hash {
	if s.privateState.Exist(addr) {
		return s.privateState.GetCodeHash(addr)
	}
	return s.state.GetCodeHash(addr)
}

//func (s MinimalApiState) Error
