// Copyright 2014 The go-earthdollar Authors
// This file is part of the go-earthdollar library.
//
// The go-earthdollar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as publisheddby
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-earthdollar library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-earthdollar library. If not, see <http://www.gnu.org/licenses/>.

// Package ed implements the Earthdollar protocol.
package ed

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Tzunami/edhash"
	"github.com/Tzunami/go-earthdollar/accounts"
	"github.com/Tzunami/go-earthdollar/common"
	"github.com/Tzunami/go-earthdollar/common/compiler"
	"github.com/Tzunami/go-earthdollar/common/httpclient"
	"github.com/Tzunami/go-earthdollar/common/registrar/edreg"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/core/types"
	"github.com/Tzunami/go-earthdollar/ed/downloader"
	"github.com/Tzunami/go-earthdollar/ed/filters"
	"github.com/Tzunami/go-earthdollar/eddb"
	"github.com/Tzunami/go-earthdollar/event"
	"github.com/Tzunami/go-earthdollar/logger"
	"github.com/Tzunami/go-earthdollar/logger/glog"
	"github.com/Tzunami/go-earthdollar/miner"
	"github.com/Tzunami/go-earthdollar/node"
	"github.com/Tzunami/go-earthdollar/p2p"
	"github.com/Tzunami/go-earthdollar/rlp"
	"github.com/Tzunami/go-earthdollar/rpc"
)

const (
	epochLength    = 30000
	edhashRevision = 23

	autoDAGcheckInterval = 10 * time.Hour
	autoDAGepochHeight   = epochLength / 2
)

type Config struct {
	ChainConfig *core.ChainConfig // chain configuration

	NetworkId int // Network ID to use for selecting peers to connect to
	Genesis   *core.GenesisDump
	FastSync  bool // Enables the state download based fast synchronisation algorithm

	BlockChainVersion  int
	SkipBcVersionCheck bool // e.g. blockchain export
	DatabaseCache      int
	DatabaseHandles    int

	NatSpec   bool
	DocRoot   string
	AutoDAG   bool
	PowTest   bool
	PowShareddbool

	AccountManager *accounts.Manager
	Earthbase      common.Address
	GasPrice       *big.Int
	MinerThreads   int
	SolcPath       string

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	TestGenesisBlock *types.Block   // Genesis block to seed the chain database with (testing only!)
	TestGenesisState eddb.Database // Genesis state to seed the database with (testing only!)
}

type Earthdollar struct {
	chainConfig *core.ChainConfig
	// Channel for shutting down the earthdollar
	shutdownChan chan bool

	// DB interfaces
	chainDb eddb.Database // Block chain database
	dappDb  eddb.Database // Dapp database

	// Handlers
	txPool          *core.TxPool
	txMu            sync.Mutex
	blockchain      *core.BlockChain
	accountManager  *accounts.Manager
	pow             *edhash.Edhash
	protocolManager *ProtocolManager
	SolcPath        string
	solc            *compiler.Solidity
	gpo             *GasPriceOracle

	GpoMinGasPrice          *big.Int
	GpoMaxGasPrice          *big.Int
	GpoFullBlockRatio       int
	GpobaseStepDown         int
	GpobaseStepUp           int
	GpobaseCorrectionFactor int

	httpclient *httpclient.HTTPClient

	eventMux *event.TypeMux
	miner    *miner.Miner

	Mining        bool
	MinerThreads  int
	NatSpec       bool
	AutoDAG       bool
	PowTest       bool
	autodagquit   chan bool
	earthbase     common.Address
	netVersionId  int
	netRPCService *PublicNetAPI
}

func New(ctx *node.ServiceContext, config *Config) (*Earthdollar, error) {
	// Open the chain database and perform any upgrades needed
	chainDb, err := ctx.OpenDatabase("chaindata", config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}
	if err := upgradeChainDatabase(chainDb); err != nil {
		return nil, err
	}
	if err := addMipmapBloomBins(chainDb); err != nil {
		return nil, err
	}

	dappDb, err := ctx.OpenDatabase("dapp", config.DatabaseCache, config.DatabaseHandles)
	if err != nil {
		return nil, err
	}

	glog.V(logger.Info).Infof("Protocol Versions: %v, Network Id: %v, Chain Id: %v", ProtocolVersions, config.NetworkId, config.ChainConfig.GetChainID())

	// Load up any custom genesis block if requested
	if config.Genesis != nil {
		_, err := core.WriteGenesisBlock(chainDb, config.Genesis)
		if err != nil {
			return nil, err
		}
	}

	// Load up a test setup if directly injected
	if config.TestGenesisState != nil {
		chainDb = config.TestGenesisState
	}
	if config.TestGenesisBlock != nil {
		core.WriteTd(chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.Difficulty())
		core.WriteBlock(chainDb, config.TestGenesisBlock)
		core.WriteCanonicalHash(chainDb, config.TestGenesisBlock.Hash(), config.TestGenesisBlock.NumberU64())
		core.WriteHeadBlockHash(chainDb, config.TestGenesisBlock.Hash())
	}

	if !config.SkipBcVersionCheck {
		bcVersion := core.GetBlockChainVersion(chainDb)
		if bcVersion != config.BlockChainVersion && bcVersion != 0 {
			return nil, fmt.Errorf("Blockchain DB version mismatch (%d / %d). Run ged upgradedb.\n", bcVersion, config.BlockChainVersion)
		}
		core.WriteBlockChainVersion(chainDb, config.BlockChainVersion)
	}
	glog.V(logger.Info).Infof("Blockchain DB Version: %d", config.BlockChainVersion)

	ed := &Earthdollar{
		shutdownChan:            make(chan bool),
		chainDb:                 chainDb,
		dappDb:                  dappDb,
		eventMux:                ctx.EventMux,
		accountManager:          config.AccountManager,
		earthbase:               config.Earthbase,
		netVersionId:            config.NetworkId,
		NatSpec:                 config.NatSpec,
		MinerThreads:            config.MinerThreads,
		SolcPath:                config.SolcPath,
		AutoDAG:                 config.AutoDAG,
		PowTest:                 config.PowTest,
		GpoMinGasPrice:          config.GpoMinGasPrice,
		GpoMaxGasPrice:          config.GpoMaxGasPrice,
		GpoFullBlockRatio:       config.GpoFullBlockRatio,
		GpobaseStepDown:         config.GpobaseStepDown,
		GpobaseStepUp:           config.GpobaseStepUp,
		GpobaseCorrectionFactor: config.GpobaseCorrectionFactor,
		httpclient:              httpclient.New(config.DocRoot),
	}
	switch {
	case config.PowTest:
        glog.V(logger.Info).Infof("Consensus: edhash used in test mode")
		ed.pow, err = edhash.NewForTesting()
		if err != nil {
			return nil, err
		}
	case config.PowShared:
        glog.V(logger.Info).Infof("Consensus: edhash used in shared mode")
		ed.pow = edhash.NewShared()
	default:
		ed.pow = edhash.New()
	}

	// load the genesis block or write a new one if no genesis
	// block is prenent in the database.
	genesis := core.GetBlock(chainDb, core.GetCanonicalHash(chainDb, 0))
	if genesis == nil {
		genesis, err = core.WriteGenesisBlock(chainDb, core.DefaultConfigMainnet.Genesis)
		if err != nil {
			return nil, err
		}
		glog.V(logger.Info).Infof("Successfully wrote default.earthdollar mainnet genesis block: %s", genesis.Hash().Hex())
	}

	// Log genesis block information. // earthdollar ERROR: replace d4e5 with actual ed hash.
	if fmt.Sprintf("%x", genesis.Hash()) == "0cd786a2425d16f152c658316c423e6ce1181e15c3295826d7c9904cba9ce303" {
		glog.V(logger.Info).Infof("Successfully established morden testnet genesis block: \x1b[36m%s\x1b[39m", genesis.Hash().Hex())
	} else if fmt.Sprintf("%x", genesis.Hash()) == "d4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3" {
		glog.V(logger.Info).Infof("Successfully established mainnet genesis block: \x1b[36m%s\x1b[39m", genesis.Hash().Hex())
	} else {
		glog.V(logger.Info).Infof("Successfully established custom genesis block: \x1b[36m%s\x1b[39m", genesis.Hash().Hex())
	}

	if config.ChainConfig == nil {
		return nil, errors.New("missing chain config")
	}

	ed.chainConfig = config.ChainConfig

	eddblockchain, err = core.NewBlockChain(chainDb, ed.chainConfig, ed.pow, ed.EventMux())
	if err != nil {
		if err == core.ErrNoGenesis {
			return nil, fmt.Errorf(`No chain found. Please initialise a new chain using the "init" subcommand.`)
		}
		return nil, err
	}
	ed.gpo = NewGasPriceOracle(ed)

	newPool := core.NewTxPool(ed.chainConfig, ed.EventMux(), eddblockchain.State, eddblockchain.GasLimit)
	ed.txPool = newPool

	if ed.protocolManager, err = NewProtocolManager(ed.chainConfig, config.FastSync, config.NetworkId, ed.eventMux, ed.txPool, ed.pow, eddblockchain, chainDb); err != nil {
		return nil, err
	}
	ed.miner = miner.New(ed, ed.chainConfig, ed.EventMux(), ed.pow)
	if err = ed.miner.SetGasPrice(config.GasPrice); err != nil {
		return nil, err
	}
	return ed, nil
}

// APIs returns the collection of RPC services the earthdollar package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Earthdollar) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "ed",
			Version:   "1.0",
			Service:   NewPublicEarthdollarAPI(s),
			Public:    true,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(s.accountManager),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(s),
			Public:    false,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(s.chainConfig, s.blockchain, s.miner, s.chainDb, s.gpo, s.eventMux, s.accountManager),
			Public:    true,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(s),
			Public:    true,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   NewPublicMinerAPI(s),
			Public:    true,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(s),
			Public:    true,
		}, {
			Namespace: "ed",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.chainDb, s.eventMux),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   edreg.NewPrivateRegistarAPI(s.chainConfig, s.blockchain, s.chainDb, s.txPool, s.accountManager),
		},
	}
}

func (s *Earthdollar) ResetWithGenesisBlock(gb *types.Block) {
	s.blockchain.ResetWithGenesisBlock(gb)
}

func (s *Earthdollar) Earthbase() (eb common.Address, err error) {
	eb = s.earthbase
	if (eb == common.Address{}) {
		firstAccount, err := s.AccountManager().AccountByIndex(0)
		eb = firstAccount.Address
		if err != nil {
			return eb, fmt.Errorf("earthbase address must be explicitly specified")
		}
	}
	return eb, nil
}

// set in js console via admin interface or wrapper from cli flags
func (self *Earthdollar) SetEarthbase(earthbase common.Address) {
	self.earthbase = earthbase
	self.miner.SetEarthbase(earthbase)
}

func (s *Earthdollar) StopMining()         { s.miner.Stop() }
func (s *Earthdollar) IsMining() bool      { return s.miner.Mining() }
func (s *Earthdollar) Miner() *miner.Miner { return s.miner }

func (s *Earthdollar) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Earthdollar) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Earthdollar) TxPool() *core.TxPool               { return s.txPool }
func (s *Earthdollar) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Earthdollar) ChainDb() eddb.Database             { return s.chainDb }
func (s *Earthdollar) DappDb() eddb.Database              { return s.dappDb }
func (s *Earthdollar) IsListening() bool                  { return true } // Always listening
func (s *Earthdollar) EdVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Earthdollar) NetVersion() int                    { return s.netVersionId }
func (s *Earthdollar) Downloader() *downloader.Downloader { return s.protocolManager.downloader }

// Protocols implements node.Service, returning all the currently configured
// network protocols to start.
func (s *Earthdollar) Protocols() []p2p.Protocol {
	return s.protocolManager.SubProtocols
}

// Start implements node.Service, starting all internal goroutines neededdby the
// Earthdollar protocol implementation.
func (s *Earthdollar) Start(srvr *p2p.Server) error {
	if s.AutoDAG {
		s.StartAutoDAG()
	}
	s.protocolManager.Start()
	s.netRPCService = NewPublicNetAPI(srvr, s.NetVersion())
	return nil
}

// Stop implements node.Service, terminating all internal goroutines useddby the
// Earthdollar protocol.
func (s *Earthdollar) Stop() error {
	s.blockchain.Stop()
	s.protocolManager.Stop()
	s.txPool.Stop()
	s.miner.Stop()
	s.eventMux.Stop()

	s.StopAutoDAG()

	s.chainDb.Close()
	s.dappDb.Close()
	close(s.shutdownChan)

	return nil
}

// This function will wait for a shutdown and resumes main thread execution
func (s *Earthdollar) WaitForShutdown() {
	<-s.shutdownChan
}

// StartAutoDAG() spawns a go routine that checks the DAG every autoDAGcheckInterval
// by default that is 10 times per epoch
// in epoch n, if we past autoDAGepochHeight within-epoch blocks,
// it calls edhash.MakeDAG  to pregenerate the DAG for the next epoch n+1
// if it does not exist yet as well as remove the DAG for epoch n-1
// the loop quits if autodagquit channel is closed, it can safely restart and
// stop any number of times.
// For any more sophisticated pattern of DAG generation, use CLI subcommand
// makedag
func (self *Earthdollar) StartAutoDAG() {
	if self.autodagquit != nil {
		return // already started
	}
	go func() {
		glog.V(logger.Info).Infof("Automatic pregeneration of edhash DAG ON (edhash dir: %s)", edhash.DefaultDir)
		var nextEpoch uint64
		timer := time.After(0)
		self.autodagquit = make(chan bool)
		for {
			select {
			case <-timer:
				glog.V(logger.Info).Infof("checking DAG (edhash dir: %s)", edhash.DefaultDir)
				currentBlock := self.BlockChain().CurrentBlock().NumberU64()
				thisEpoch := currentBlock / epochLength
				if nextEpoch <= thisEpoch {
					if currentBlock%epochLength > autoDAGepochHeight {
						if thisEpoch > 0 {
							previousDag, previousDagFull := dagFiles(thisEpoch - 1)
							os.Remove(filepath.Join(edhash.DefaultDir, previousDag))
							os.Remove(filepath.Join(edhash.DefaultDir, previousDagFull))
							glog.V(logger.Info).Infof("removed DAG for epoch %d (%s)", thisEpoch-1, previousDag)
						}
						nextEpoch = thisEpoch + 1
						dag, _ := dagFiles(nextEpoch)
						if _, err := os.Stat(dag); os.IsNotExist(err) {
							glog.V(logger.Info).Infof("Pregenerating DAG for epoch %d (%s)", nextEpoch, dag)
							err := edhash.MakeDAG(nextEpoch*epochLength, "") // "" -> edhash.DefaultDir
							if err != nil {
								glog.V(logger.Error).Infof("Error generating DAG for epoch %d (%s)", nextEpoch, dag)
								return
							}
						} else {
							glog.V(logger.Error).Infof("DAG for epoch %d (%s)", nextEpoch, dag)
						}
					}
				}
				timer = time.After(autoDAGcheckInterval)
			case <-self.autodagquit:
				return
			}
		}
	}()
}

// stopAutoDAG stops automatic DAG pregeneration by quitting the loop
func (self *Earthdollar) StopAutoDAG() {
	if self.autodagquit != nil {
		close(self.autodagquit)
		self.autodagquit = nil
	}
	glog.V(logger.Info).Infof("Automatic pregeneration of edhash DAG OFF (edhash dir: %s)", edhash.DefaultDir)
}

// HTTPClient returns the light http client used for fetching offchain docs
// (natspec, source for verification)
func (self *Earthdollar) HTTPClient() *httpclient.HTTPClient {
	return self.httpclient
}

func (self *Earthdollar) Solc() (*compiler.Solidity, error) {
	var err error
	if self.solc == nil {
		self.solc, err = compiler.New(self.SolcPath)
	}
	return self.solc, err
}

// set in js console via admin interface or wrapper from cli flags
func (self *Earthdollar) SetSolc(solcPath string) (*compiler.Solidity, error) {
	self.SolcPath = solcPath
	self.solc = nil
	return self.Solc()
}

// dagFiles(epoch) returns the two alternative DAG filenames (not a path)
// 1) <revision>-<hex(seedhash[8])> 2) full-R<revision>-<hex(seedhash[8])>
func dagFiles(epoch uint64) (string, string) {
	seedHash, _ := edhash.GetSeedHash(epoch * epochLength)
	dag := fmt.Sprintf("full-R%d-%x", edhashRevision, seedHash[:8])
	return dag, "full-R" + dag
}

// upgradeChainDatabase ensures that the chain database stores block split into
// separate header and body entries.
func upgradeChainDatabase(db eddb.Database) error {
	// Short circuit if the head block is stored already as separate header and body
	data, err := db.Get([]byte("LastBlock"))
	if err != nil {
		return nil
	}
	head := common.BytesToHash(data)

	if block := core.GetBlockByHashOld(db, head); block == nil {
		return nil
	}
	// At least some of the database is still the old format, upgrade (skip the head block!)
	glog.V(logger.Info).Info("Old database detected, upgrading...")

	if db, ok := db.(*eddb.LDBDatabase); ok {
		blockPrefix := []byte("block-hash-")
		for it := db.NewIterator(); it.Next(); {
			// Skip anything other than a combineddblock
			if !bytes.HasPrefix(it.Key(), blockPrefix) {
				continue
			}
			// Skip the head block (merge last to signal upgrade completion)
			if bytes.HasSuffix(it.Key(), head.Bytes()) {
				continue
			}
			// Load the block, split and serialize (order!)
			block := core.GetBlockByHashOld(db, common.BytesToHash(bytes.TrimPrefix(it.Key(), blockPrefix)))

			if err := core.WriteTd(db, block.Hash(), block.DeprecatedTd()); err != nil {
				return err
			}
			if err := core.WriteBody(db, block.Hash(), block.Body()); err != nil {
				return err
			}
			if err := core.WriteHeader(db, block.Header()); err != nil {
				return err
			}
			if err := db.Delete(it.Key()); err != nil {
				return err
			}
		}
		// Lastly, upgrade the head block, disabling the upgrade mechanism
		current := core.GetBlockByHashOld(db, head)

		if err := core.WriteTd(db, current.Hash(), current.DeprecatedTd()); err != nil {
			return err
		}
		if err := core.WriteBody(db, current.Hash(), current.Body()); err != nil {
			return err
		}
		if err := core.WriteHeader(db, current.Header()); err != nil {
			return err
		}
	}
	return nil
}

func addMipmapBloomBins(db eddb.Database) (err error) {
	const mipmapVersion uint = 2

	// check if the version is set. We ignore data for now since there's
	// only one version so we can easily ignore it for now
	var data []byte
	data, _ = db.Get([]byte("setting-mipmap-version"))
	if len(data) > 0 {
		var version uint
		if err := rlp.DecodeBytes(data, &version); err == nil && version == mipmapVersion {
			return nil
		}
	}

	defer func() {
		if err == nil {
			var val []byte
			val, err = rlp.EncodeToBytes(mipmapVersion)
			if err == nil {
				err = db.Put([]byte("setting-mipmap-version"), val)
			}
			return
		}
	}()
	latestBlock := core.GetBlock(db, core.GetHeadBlockHash(db))
	if latestBlock == nil { // clean database
		return
	}

	tstart := time.Now()
	glog.V(logger.Info).Infoln("upgrading db log bloom bins")
	for i := uint64(0); i <= latestBlock.NumberU64(); i++ {
		hash := core.GetCanonicalHash(db, i)
		if (hash == common.Hash{}) {
			return fmt.Errorf("chain db corrupted. Could not find block %d.", i)
		}
		core.WriteMipmapBloom(db, i, core.GetBlockReceipts(db, hash))
	}
	glog.V(logger.Info).Infoln("upgrade completed in", time.Since(tstart))
	return nil
}
