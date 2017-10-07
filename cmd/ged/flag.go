// Copyright 2015 The go-earthdollar Authors
// This file is part of go-earthdollar.
//
// go-earthdollar is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-earthdollar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-earthdollar. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
<<<<<<< HEAD:cmd/ged/flag.go
	"math"
=======
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

<<<<<<< HEAD:cmd/ged/flag.go
	"github.com/Tzunami/edhash"
	"github.com/Tzunami/go-earthdollar/accounts"
	"github.com/Tzunami/go-earthdollar/common"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/core/state"
	"github.com/Tzunami/go-earthdollar/core/types"
	"github.com/Tzunami/go-earthdollar/crypto"
	"github.com/Tzunami/go-earthdollar/ed"
	"github.com/Tzunami/go-earthdollar/eddb"
	"github.com/Tzunami/go-earthdollar/event"
	"github.com/Tzunami/go-earthdollar/logger"
	"github.com/Tzunami/go-earthdollar/logger/glog"
	"github.com/Tzunami/go-earthdollar/miner"
	"github.com/Tzunami/go-earthdollar/node"
	"github.com/Tzunami/go-earthdollar/p2p/discover"
	"github.com/Tzunami/go-earthdollar/p2p/nat"
	"github.com/Tzunami/go-earthdollar/pow"
	"github.com/Tzunami/go-earthdollar/rpc"
	"github.com/Tzunami/go-earthdollar/whisper"
=======
	"errors"

	"time"

	"github.com/ethereumproject/ethash"
	"github.com/ethereumproject/go-ethereum/accounts"
	"github.com/ethereumproject/go-ethereum/common"
	"github.com/ethereumproject/go-ethereum/core"
	"github.com/ethereumproject/go-ethereum/core/state"
	"github.com/ethereumproject/go-ethereum/core/types"
	"github.com/ethereumproject/go-ethereum/crypto"
	"github.com/ethereumproject/go-ethereum/eth"
	"github.com/ethereumproject/go-ethereum/ethdb"
	"github.com/ethereumproject/go-ethereum/event"
	"github.com/ethereumproject/go-ethereum/logger"
	"github.com/ethereumproject/go-ethereum/logger/glog"
	"github.com/ethereumproject/go-ethereum/miner"
	"github.com/ethereumproject/go-ethereum/node"
	"github.com/ethereumproject/go-ethereum/p2p/discover"
	"github.com/ethereumproject/go-ethereum/p2p/nat"
	"github.com/ethereumproject/go-ethereum/pow"
	"github.com/ethereumproject/go-ethereum/whisper"
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	"gopkg.in/urfave/cli.v1"
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	cli.CommandHelpTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} [arguments...]
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
	{{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
	{{end}}{{end}}{{if .Flags}}
OPTIONS:
	{{range .Flags}}{{.}}
	{{end}}{{end}}
`
}

var (
<<<<<<< HEAD:cmd/ged/flag.go
	// General settings
	DataDirFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: DirectoryString{common.DefaultDataDir()},
	}
	KeyStoreDirFlag = DirectoryFlag{
		Name:  "keystore",
		Usage: "Directory for the keystore (default = inside the datadir)",
	}
	NetworkIdFlag = cli.IntFlag{
		Name:  "networkid",
		Usage: "Network identifier (integer, 0=Olympic, 1=Homestead, 2=Morden)",
		Value: ed.NetworkId,
	}
	OlympicFlag = cli.BoolFlag{ 
		Name:  "olympic",
		Usage: "Olympic network: pre-configured pre-release test network",
	}
	TestNetFlag = cli.BoolFlag{
		Name:  "testnet",
		Usage: "Morden network: pre-configured test network with modified starting nonces (replay protection)",
	}
	DevModeFlag = cli.BoolFlag{
		Name:  "dev",
		Usage: "Developer mode: pre-configured private network with several debugging flags",
	}
	IdentityFlag = cli.StringFlag{
		Name:  "identity",
		Usage: "Custom node name",
	}
	NatspecEnabledFlag = cli.BoolFlag{
		Name:  "natspec",
		Usage: "Enable NatSpec confirmation notice",
	}
	DocRootFlag = DirectoryFlag{
		Name:  "docroot",
		Usage: "Document Root for HTTPClient file scheme",
		Value: DirectoryString{common.HomeDir()},
	}
	CacheFlag = cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching (min 16MB / database forced)",
		Value: 128,
	}
	BlockchainVersionFlag = cli.IntFlag{
		Name:  "blockchainversion",
		Usage: "Blockchain version (integer)",
		Value: core.BlockChainVersion,
	}
	FastSyncFlag = cli.BoolFlag{
		Name:  "fast",
		Usage: "Enable fast syncing through state downloads",
	}
	LightKDFFlag = cli.BoolFlag{
		Name:  "lightkdf",
		Usage: "Reduce key-derivation RAM & CPU usage at some expense of KDF strength",
	}
	// Network Split settings
	ETFChain = cli.BoolFlag{
		Name:  "etf",
		Usage: "Updates the chain rules to use the ETF hard-fork blockchain",
	}
	// Miner settings
	// TODO Refactor CPU vs GPU mining flags
	MiningEnabledFlag = cli.BoolFlag{
		Name:  "mine",
		Usage: "Enable mining",
	}
	MinerThreadsFlag = cli.IntFlag{
		Name:  "minerthreads",
		Usage: "Number of CPU threads to use for mining",
		Value: runtime.NumCPU(),
	}
	MiningGPUFlag = cli.StringFlag{
		Name:  "minergpus",
		Usage: "List of GPUs to use for mining (e.g. '0,1' will use the first two GPUs found)",
	}
	TargetGasLimitFlag = cli.StringFlag{
		Name:  "targetgaslimit",
		Usage: "Target gas limit sets the artificial target gas floor for the blocks to mine",
		Value: core.TargetGasLimit.String(),
	}
	AutoDAGFlag = cli.BoolFlag{
		Name:  "autodag",
		Usage: "Enable automatic DAG pregeneration",
	}
	EarthbaseFlag = cli.StringFlag{
		Name:  "earthbase",
		Usage: "Public address for block mining rewards (default = first account created)",
		Value: "0",
	}
	GasPriceFlag = cli.StringFlag{
		Name:  "gasprice",
		Usage: "Minimal gas price to accept for mining a transactions",
		Value: new(big.Int).Mul(big.NewInt(20), common.Chief).String(),
	}
	ExtraDataFlag = cli.StringFlag{
		Name:  "extradata",
		Usage: "Freeform header field set by the miner",
	}
	// Account settings
	UnlockedAccountFlag = cli.StringFlag{
		Name:  "unlock",
		Usage: "Comma separated list of accounts to unlock",
		Value: "",
	}
	PasswordFileFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password file to use for non-inteactive password input",
		Value: "",
	}

	// logging and debug settings
	VerbosityFlag = cli.GenericFlag{
		Name:  "verbosity",
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=core, 5=debug, 6=detail",
		Value: glog.GetVerbosity(),
	}
	VModuleFlag = cli.GenericFlag{
		Name:  "vmodule",
		Usage: "Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. ed/*=6,p2p=5)",
		Value: glog.GetVModule(),
	}
	BacktraceAtFlag = cli.GenericFlag{
		Name:  "backtrace",
		Usage: "Request a stack trace at a specific logging statement (e.g. \"block.go:271\")",
		Value: glog.GetTraceLocation(),
	}
	MetricsFlag = cli.StringFlag{
		Name:  "metrics",
		Usage: "Enables metrics reporting. When the value is a path, either relative or absolute, then a log is written to the respective file.",
	}
	FakePoWFlag = cli.BoolFlag{
		Name:  "fakepow",
		Usage: "Disables proof-of-work verification",
	}

	// RPC settings
	RPCEnabledFlag = cli.BoolFlag{
		Name:  "rpc",
		Usage: "Enable the HTTP-RPC server",
	}
	RPCListenAddrFlag = cli.StringFlag{
		Name:  "rpcaddr",
		Usage: "HTTP-RPC server listening interface",
		Value: common.DefaultHTTPHost,
	}
	RPCPortFlag = cli.IntFlag{
		Name:  "rpcport",
		Usage: "HTTP-RPC server listening port",
		Value: common.DefaultHTTPPort,
	}
	RPCCORSDomainFlag = cli.StringFlag{
		Name:  "rpccorsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	RPCApiFlag = cli.StringFlag{
		Name:  "rpcapi",
		Usage: "API's offered over the HTTP-RPC interface",
		Value: rpc.DefaultHTTPApis,
	}
	IPCDisabledFlag = cli.BoolFlag{
		Name:  "ipcdisable",
		Usage: "Disable the IPC-RPC server",
	}
	IPCApiFlag = cli.StringFlag{
		Name:  "ipcapi",
		Usage: "API's offered over the IPC-RPC interface",
		Value: rpc.DefaultIPCApis,
	}
	IPCPathFlag = DirectoryFlag{
		Name:  "ipcpath",
		Usage: "Filename for IPC socket/pipe within the datadir (explicit paths escape it)",
		Value: DirectoryString{common.DefaultIPCSocket},
	}
	WSEnabledFlag = cli.BoolFlag{
		Name:  "ws",
		Usage: "Enable the WS-RPC server",
	}
	WSListenAddrFlag = cli.StringFlag{
		Name:  "wsaddr",
		Usage: "WS-RPC server listening interface",
		Value: common.DefaultWSHost,
	}
	WSPortFlag = cli.IntFlag{
		Name:  "wsport",
		Usage: "WS-RPC server listening port",
		Value: common.DefaultWSPort,
	}
	WSApiFlag = cli.StringFlag{
		Name:  "wsapi",
		Usage: "API's offered over the WS-RPC interface",
		Value: rpc.DefaultHTTPApis,
	}
	WSAllowedOriginsFlag = cli.StringFlag{
		Name:  "wsorigins",
		Usage: "Origins from which to accept websockets requests",
		Value: "",
	}
	ExecFlag = cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement (only in combination with console/attach)",
	}
	PreloadJSFlag = cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}

	// Network Settings
	MaxPeersFlag = cli.IntFlag{
		Name:  "maxpeers",
		Usage: "Maximum number of network peers (network disabled if set to 0)",
		Value: 25,
	}
	MaxPendingPeersFlag = cli.IntFlag{
		Name:  "maxpendpeers",
		Usage: "Maximum number of pending connection attempts (defaults used if set to 0)",
		Value: 0,
	}
	ListenPortFlag = cli.IntFlag{
		Name:  "port",
		Usage: "Network listening port",
		Value: 20203,
	}
	BootnodesFlag = cli.StringFlag{
		Name:  "bootnodes",
		Usage: "Comma separated enode URLs for P2P discovery bootstrap",
		Value: "",
	}
	NodeKeyFileFlag = cli.StringFlag{
		Name:  "nodekey",
		Usage: "P2P node key file",
	}
	NodeKeyHexFlag = cli.StringFlag{
		Name:  "nodekeyhex",
		Usage: "P2P node key as hex (for testing)",
	}
	NATFlag = cli.StringFlag{
		Name:  "nat",
		Usage: "NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
		Value: "any",
	}
	NoDiscoverFlag = cli.BoolFlag{
		Name:  "nodiscover",
		Usage: "Disables the peer discovery mechanism (manual peer addition)",
	}
	WhisperEnabledFlag = cli.BoolFlag{
		Name:  "shh",
		Usage: "Enable Whisper",
	}
	// ATM the url is left to the user and deployment to
	JSpathFlag = cli.StringFlag{
		Name:  "jspath",
		Usage: "JavaScript root path for `loadScript` and document root for `admin.httpGet`",
		Value: ".",
	}
	SolcPathFlag = cli.StringFlag{
		Name:  "solc",
		Usage: "Solidity compiler command to be used",
		Value: "solc",
	}

	// Gas price oracle settings
	GpoMinGasPriceFlag = cli.StringFlag{
		Name:  "gpomin",
		Usage: "Minimum suggested gas price",
		Value: new(big.Int).Mul(big.NewInt(20), common.Chief).String(),
	}
	GpoMaxGasPriceFlag = cli.StringFlag{
		Name:  "gpomax",
		Usage: "Maximum suggested gas price",
		Value: new(big.Int).Mul(big.NewInt(500), common.Chief).String(),
	}
	GpoFullBlockRatioFlag = cli.IntFlag{
		Name:  "gpofull",
		Usage: "Full block threshold for gas price calculation (%)",
		Value: 80,
	}
	GpobaseStepDownFlag = cli.IntFlag{
		Name:  "gpobasedown",
		Usage: "Suggested gas price base step down ratio (1/1000)",
		Value: 10,
	}
	GpobaseStepUpFlag = cli.IntFlag{
		Name:  "gpobaseup",
		Usage: "Suggested gas price base step up ratio (1/1000)",
		Value: 100,
	}
	GpobaseCorrectionFactorFlag = cli.IntFlag{
		Name:  "gpobasecf",
		Usage: "Suggested gas price base correction factor (%)",
		Value: 110,
	}
	Unused1 = cli.BoolFlag{
		Name:  "oppose-dao-fork",
		Usage: "Use classic blockchain (always set, flag is unused and exists for compatibility only)",
=======
	// Errors.
	ErrInvalidFlag        = errors.New("invalid flag or context value")
	ErrInvalidChainID     = errors.New("invalid chainID")
	ErrDirectoryStructure = errors.New("error in directory structure")
	ErrStackFail          = errors.New("error in stack protocol")

	// Chain identities.
	chainIdentitiesBlacklist = map[string]bool{
		"chaindata": true,
		"dapp":      true,
		"keystore":  true,
		"nodekey":   true,
		"nodes":     true,
	}
	chainIdentitiesMain = map[string]bool{
		"main":    true,
		"mainnet": true,
	}
	chainIdentitiesMorden = map[string]bool{
		"morden":  true,
		"testnet": true,
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

	devModeDataDirPath = filepath.Join(os.TempDir(), "/ethereum_dev_mode")

	cacheChainIdentity string
	cacheChainConfig *core.SufficientChainConfig
)

// chainIsMorden allows either
// '--testnet' (legacy), or
// '--chain=morden|testnet'
func chainIsMorden(ctx *cli.Context) bool {
	if ctx.GlobalIsSet(aliasableName(TestNetFlag.Name, ctx)) && ctx.GlobalIsSet(aliasableName(ChainIdentityFlag.Name, ctx)) {
		glog.Fatalf(`%v: used redundant/conflicting flags: %v, %v
		Please use one or the other, but not both.`, ErrInvalidFlag, aliasableName(TestNetFlag.Name, ctx), aliasableName(ChainIdentityFlag.Name, ctx))
		return false
	}
	return ctx.GlobalBool(aliasableName(TestNetFlag.Name, ctx)) || chainIdentitiesMorden[ctx.GlobalString(aliasableName(ChainIdentityFlag.Name, ctx))]
}

// if argument to --chain is a path and is a valid configuration file, copy it to
// identity/chain.json. It will overwrite an existing configuration file with same identity.
// This allows specification of a chain config by filename and subsequently by configured identity as well.
func copyChainConfigFileToChainDataDir(ctx *cli.Context, identity, configFilePath string) error {
	// Ensure directory path exists.
	identityDirPath := common.EnsurePathAbsoluteOrRelativeTo(mustMakeDataDir(ctx), identity)
	if e := os.MkdirAll(identityDirPath, os.ModePerm); e != nil {
		return e
	}

	glog.V(logger.Debug).Infof("Copying %v to %v/chain.json", configFilePath, identityDirPath)
	b, e := ioutil.ReadFile(configFilePath)
	if e != nil {
		return e
	}

	if e := ioutil.WriteFile(filepath.Join(identityDirPath, "chain.json"), b, os.ModePerm); e != nil {
		return e
	}
	return nil
}

// getChainIdentity parses --chain and --testnet (legacy) flags.
// It will fatal if finds notok value.
// It returns one of valid strings: ["mainnet", "morden", or --chain="flaggedCustom"]
func mustMakeChainIdentity(ctx *cli.Context) (identity string) {

	if cacheChainIdentity != "" {
		return cacheChainIdentity
	}
	defer func () {
		cacheChainIdentity = identity
	}()

	if chainIsMorden(ctx) {
		identity = core.DefaultConfigMorden.Identity // this makes '--testnet', '--chain=testnet', and '--chain=morden' all use the same /morden subdirectory, if --chain isn't specified
		return identity
	}
	// If --chain is in use.
	if chainFlagVal := ctx.GlobalString(aliasableName(ChainIdentityFlag.Name, ctx)); chainFlagVal != "" {
		if chainIdentitiesMain[chainFlagVal] {
			identity = core.DefaultConfigMainnet.Identity
			return identity
		}
		// Check for unallowed values.
		if chainIdentitiesBlacklist[chainFlagVal] {
			glog.Fatalf(`%v: %v: reserved word
					reserved words for --chain flag include: 'chaindata', 'dapp', 'keystore', 'nodekey', 'nodes',
					please use a different identifier`, ErrInvalidFlag, ErrInvalidChainID)
			identity = ""
			return identity
		}

		// Check if passed arg exists as a path to a valid config file.
		if fstat, ferr := os.Stat(filepath.Clean(chainFlagVal)); ferr == nil && !fstat.IsDir() {
			glog.V(logger.Debug).Infof("Found existing file at --%v: %v", aliasableName(ChainIdentityFlag.Name, ctx), chainFlagVal)
			c, configurationError := core.ReadExternalChainConfigFromFile(filepath.Clean(chainFlagVal))
			if configurationError == nil {
				glog.V(logger.Debug).Infof("OK: Valid chain configuration. Chain identity: %v", c.Identity)
				if e := copyChainConfigFileToChainDataDir(ctx, c.Identity, filepath.Clean(chainFlagVal)); e != nil {
					glog.Fatalf("Could not copy chain configuration: %v", e)
				}
				// In edge case of using a config file for default configuration (decided by 'identity'),
				// set global context and override config file.
				if chainIdentitiesMorden[c.Identity] || chainIdentitiesMain[c.Identity] {
					if e := ctx.Set(aliasableName(ChainIdentityFlag.Name, ctx), c.Identity); e != nil {
						glog.Fatalf("Could not set global context chain identity to morden, error: %v", e)
					}
				}
				identity = c.Identity
				return identity
			}
			glog.Fatalf("Invalid chain config file at --%v: '%v': %v \nAssuming literal identity argument.",
				aliasableName(ChainIdentityFlag.Name, ctx), chainFlagVal, configurationError)
		}
		glog.V(logger.Debug).Infof("No existing file at --%v: '%v'. Using literal chain identity.", aliasableName(ChainIdentityFlag.Name, ctx), chainFlagVal)
		identity = chainFlagVal
		return identity
	} else if ctx.GlobalIsSet(aliasableName(ChainIdentityFlag.Name, ctx)) {
		glog.Fatalf("%v: %v: chainID empty", ErrInvalidFlag, ErrInvalidChainID)
		identity = ""
		return identity
	}
	// If no relevant flag is set, return default mainnet.
	identity = core.DefaultConfigMainnet.Identity
	return identity
}

// mustMakeChainConfigNameDefaulty gets mainnet or testnet defaults if in use.
// _If a custom net is in use, it echoes the name of the ChainConfigID._
// It is intended to be a human-readable name for a chain configuration.
// - It should only be called in reference to default configuration (name will be configured
// separately through external JSON config otherwise).
func mustMakeChainConfigNameDefaulty(ctx *cli.Context) string {
	if chainIsMorden(ctx) {
		return core.DefaultConfigMorden.Name
	}
	return core.DefaultConfigMainnet.Name
}

// mustMakeDataDir retrieves the currently requested data directory, terminating
// if none (or the empty string) is specified.
// --> <home>/<EthereumClassic>(defaulty) or --datadir
func mustMakeDataDir(ctx *cli.Context) string {
	if !ctx.GlobalIsSet(aliasableName(DataDirFlag.Name, ctx)) {
		if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
			return devModeDataDirPath
		}
	}
	if path := ctx.GlobalString(aliasableName(DataDirFlag.Name, ctx)); path != "" {
		return path
	}

	glog.Fatalf("%v: cannot determine data directory, please set manually (--%v)", ErrDirectoryStructure, DataDirFlag.Name)
	return ""
}

// MustMakeChainDataDir retrieves the currently requested data directory including chain-specific subdirectory.
// A subdir of the datadir is used for each chain configuration ("/mainnet", "/testnet", "/my-custom-net").
// --> <home>/<EthereumClassic>/<mainnet|testnet|custom-net>, per --chain
func MustMakeChainDataDir(ctx *cli.Context) string {
	rp := common.EnsurePathAbsoluteOrRelativeTo(mustMakeDataDir(ctx), mustMakeChainIdentity(ctx))
	if !filepath.IsAbs(rp) {
		af, e := filepath.Abs(rp)
		if e != nil {
			glog.Fatalf("cannot make absolute path for chain data dir: %v: %v", rp, e)
		}
		rp = af
	}
	return rp
}

// MakeIPCPath creates an IPC path configuration from the set command line flags,
// returning an empty string if IPC was explicitly disabled, or the set path.
func MakeIPCPath(ctx *cli.Context) string {
	if ctx.GlobalBool(aliasableName(IPCDisabledFlag.Name, ctx)) {
		return ""
	}
	return ctx.GlobalString(aliasableName(IPCPathFlag.Name, ctx))
}

// MakeNodeKey creates a node key from set command line flags, either loading it
// from a file or as a specified hex value. If neither flags were provided, this
// method returns nil and an emphemeral key is to be generated.
func MakeNodeKey(ctx *cli.Context) *ecdsa.PrivateKey {
	var (
		hex  = ctx.GlobalString(aliasableName(NodeKeyHexFlag.Name, ctx))
		file = ctx.GlobalString(aliasableName(NodeKeyFileFlag.Name, ctx))

		key *ecdsa.PrivateKey
		err error
	)
	switch {
	case file != "" && hex != "":
		log.Fatalf("Options %q and %q are mutually exclusive", aliasableName(NodeKeyFileFlag.Name, ctx), aliasableName(NodeKeyHexFlag.Name, ctx))

	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			log.Fatalf("Option %q: %v", aliasableName(NodeKeyFileFlag.Name, ctx), err)
		}

	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			log.Fatalf("Option %q: %v", aliasableName(NodeKeyHexFlag.Name, ctx), err)
		}
	}
	return key
}

// MakeBootstrapNodesFromContext creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func MakeBootstrapNodesFromContext(ctx *cli.Context) []*discover.Node {
	// Return pre-configured nodes if none were manually requested
	if !ctx.GlobalIsSet(aliasableName(BootnodesFlag.Name, ctx)) {

		// --testnet/--chain=morden flag overrides --config flag
		if chainIsMorden(ctx) {
			return core.DefaultConfigMorden.ParsedBootstrap
		}
		return core.DefaultConfigMainnet.ParsedBootstrap
	}
	return core.ParseBootstrapNodeStrings(strings.Split(ctx.GlobalString(aliasableName(BootnodesFlag.Name, ctx)), ","))
}

// MakeListenAddress creates a TCP listening address string from set command
// line flags.
func MakeListenAddress(ctx *cli.Context) string {
	return fmt.Sprintf(":%d", ctx.GlobalInt(aliasableName(ListenPortFlag.Name, ctx)))
}

// MakeNAT creates a port mapper from set command line flags.
func MakeNAT(ctx *cli.Context) nat.Interface {
	natif, err := nat.Parse(ctx.GlobalString(aliasableName(NATFlag.Name, ctx)))
	if err != nil {
		log.Fatalf("Option %s: %v", aliasableName(NATFlag.Name, ctx), err)
	}
	return natif
}

// MakeRPCModules splits input separated by a comma and trims excessive white
// space from the substrings.
func MakeRPCModules(input string) []string {
	result := strings.Split(input, ",")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// MakeHTTPRpcHost creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func MakeHTTPRpcHost(ctx *cli.Context) string {
	if !ctx.GlobalBool(aliasableName(RPCEnabledFlag.Name, ctx)) {
		return ""
	}
	return ctx.GlobalString(aliasableName(RPCListenAddrFlag.Name, ctx))
}

// MakeWSRpcHost creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func MakeWSRpcHost(ctx *cli.Context) string {
	if !ctx.GlobalBool(aliasableName(WSEnabledFlag.Name, ctx)) {
		return ""
	}
	return ctx.GlobalString(aliasableName(WSListenAddrFlag.Name, ctx))
}

// MakeDatabaseHandles raises out the number of allowed file handles per process
// for Ged and returns half of the allowance to assign to the database.
func MakeDatabaseHandles() int {
	if err := raiseFdLimit(2048); err != nil {
		glog.V(logger.Warn).Info("Failed to raise file descriptor allowance: ", err)
	}
	limit, err := getFdLimit()
	if err != nil {
		glog.V(logger.Warn).Info("Failed to retrieve file descriptor allowance: ", err)
	}
	if limit > 2048 { // cap database file descriptors even if more is available
		limit = 2048
	}
	return limit / 2 // Leave half for networking and other stuff
}

// MakeAccountManager creates an account manager from set command line flags.
func MakeAccountManager(ctx *cli.Context) *accounts.Manager {
	// Create the keystore crypto primitive, light if requested
	scryptN := accounts.StandardScryptN
	scryptP := accounts.StandardScryptP
	if ctx.GlobalBool(aliasableName(LightKDFFlag.Name, ctx)) {
		scryptN = accounts.LightScryptN
		scryptP = accounts.LightScryptP
	}

	datadir := MustMakeChainDataDir(ctx)

	keydir := filepath.Join(datadir, "keystore")
	if path := ctx.GlobalString(aliasableName(KeyStoreDirFlag.Name, ctx)); path != "" {
		af, e := filepath.Abs(path)
		if e != nil {
			glog.V(logger.Error).Infof("keydir path could not be made absolute: %v: %v", path, e)
		} else {
			keydir = af
		}
	}

	m, err := accounts.NewManager(keydir, scryptN, scryptP, ctx.GlobalBool(aliasableName(AccountsIndexFlag.Name, ctx)))
	if err != nil {
		glog.Fatalf("init account manager at %q: %s", keydir, err)
	}
	return m
}

// MakeAddress converts an account specified directly as a hex encoded string or
// a key index in the key store to an internal account representation.
func MakeAddress(accman *accounts.Manager, account string) (accounts.Account, error) {
	// If the specified account is a valid address, return it
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}
	// Otherwise try to interpret the account as a keystore index
	index, err := strconv.Atoi(account)
	if err != nil {
		return accounts.Account{}, fmt.Errorf("invalid account address or index: %q", account)
	}
	return accman.AccountByIndex(index)
}

// MakeEarthbase retrieves the earthbase either from the directly specified
// command line flags or from the keystore if CLI indexed.
func MakeEarthbase(accman *accounts.Manager, ctx *cli.Context) common.Address {
	accounts := accman.Accounts()
<<<<<<< HEAD:cmd/ged/flag.go
	if !ctx.GlobalIsSet(EarthbaseFlag.Name) && len(accounts) == 0 {
		glog.V(logger.Error).Infoln("WARNING: No earthbase set and no accounts found as default")
		return common.Address{}
	}
	earthbase := ctx.GlobalString(EarthbaseFlag.Name)
	if earthbase == "" {
=======
	if !ctx.GlobalIsSet(aliasableName(EtherbaseFlag.Name, ctx)) && len(accounts) == 0 {
		glog.V(logger.Warn).Infoln("WARNING: No etherbase set and no accounts found as default")
		return common.Address{}
	}
	etherbase := ctx.GlobalString(aliasableName(EtherbaseFlag.Name, ctx))
	if etherbase == "" {
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
		return common.Address{}
	}
	// If the specified earthbase is a valid address, return it
	account, err := MakeAddress(accman, earthbase)
	if err != nil {
<<<<<<< HEAD:cmd/ged/flag.go
		log.Fatalf("Option %q: %v", EarthbaseFlag.Name, err)
=======
		log.Fatalf("Option %q: %v", aliasableName(EtherbaseFlag.Name, ctx), err)
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}
	return account.Address
}

// MakePasswordList reads password lines from the file specified by --password.
func MakePasswordList(ctx *cli.Context) []string {
	path := ctx.GlobalString(aliasableName(PasswordFileFlag.Name, ctx))
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Fatal("Failed to read password file: ", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

<<<<<<< HEAD:cmd/ged/flag.go
// MakeSystemNode sets up a local node, configures the services to launch and
// assembles the P2P protocol stack.
func MakeSystemNode(version string, ctx *cli.Context) *node.Node {
	name := fmt.Sprintf("Ged/%s/%s/%s", version, runtime.GOOS, runtime.Version())
	if identity := ctx.GlobalString(IdentityFlag.Name); len(identity) > 0 {
=======
// makeName makes the node name, which can be (in part) customized by the NodeNameFlag
func makeNodeName(version string, ctx *cli.Context) string {
	name := fmt.Sprintf("Geth/%s/%s/%s", version, runtime.GOOS, runtime.Version())
	if identity := ctx.GlobalString(aliasableName(NodeNameFlag.Name, ctx)); len(identity) > 0 {
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
		name += "/" + identity
	}
	return name
}

func mustMakeMLogDir(ctx *cli.Context) string {
	if ctx.GlobalIsSet(MLogDirFlag.Name) {
		p := ctx.GlobalString(MLogDirFlag.Name)
		if p == "" {
			glog.Fatalf("Flag %v requires a non-empty argument", MLogDirFlag.Name)
			return ""
		}
		if filepath.IsAbs(p) {
			return p
		}
		ap, e := filepath.Abs(p)
		if e != nil {
			glog.Fatalf("could not establish absolute path for mlog dir: %v", e)
		}
		return ap
	}

	return filepath.Join(MustMakeChainDataDir(ctx), "mlogs")
}

func makeMLogFileLogger(ctx *cli.Context) (string, error) {
	now := time.Now()

	mlogdir := mustMakeMLogDir(ctx)
	logger.SetMLogDir(mlogdir)

	_, filename, err := logger.CreateMLogFile(now)
	if err != nil {
		return "", err
	}
	// withTs toggles custom timestamp ISO8601 prefix
	// logger print without timestamp header prefix if json
	withTs := true
	if f := ctx.GlobalString(MLogFlag.Name); logger.MLogStringToFormat[f] == logger.MLOGJSON {
		withTs = false
	}
	logger.BuildNewMLogSystem(mlogdir, filename, 1, 0, withTs) // flags: 0 disables automatic log package time prefix
	return filename, nil
}

func mustRegisterMLogsFromContext(ctx *cli.Context) {
	if e := logger.MLogRegisterComponentsFromContext(ctx.GlobalString(MLogComponentsFlag.Name)); e != nil {
		// print documentation if user enters unavailable mlog component
		var components []string
		for k := range logger.MLogRegistryAvailable {
			components = append(components, string(k))
		}
		glog.V(logger.Error).Infof("Error: %s", e)
		glog.V(logger.Error).Infof("Available machine log components: %v", components)
		os.Exit(1)
	}
	// Set the global logger mlog format from context
	if e := logger.SetMLogFormatFromString(ctx.GlobalString(MLogFlag.Name)); e != nil {
		glog.Fatalf("Error setting mlog format: %v, value was: %v", e, ctx.GlobalString(MLogFlag.Name))
	}
	fname, e := makeMLogFileLogger(ctx)
	if e != nil {
		glog.Fatalf("Failed to start machine log: %v", e)
	}
	glog.V(logger.Info).Infof("Machine logs file: %v", fname)
}

// MakeSystemNode sets up a local node, configures the services to launch and
// assembles the P2P protocol stack.
func MakeSystemNode(version string, ctx *cli.Context) *node.Node {

	// global settings

	if ctx.GlobalIsSet(aliasableName(ExtraDataFlag.Name, ctx)) {
		s := ctx.GlobalString(aliasableName(ExtraDataFlag.Name, ctx))
		if len(s) > types.HeaderExtraMax {
			log.Fatalf("%s flag %q exceeds size limit of %d", aliasableName(ExtraDataFlag.Name, ctx), s, types.HeaderExtraMax)
		}
		miner.HeaderExtra = []byte(s)
	}

	// Data migrations if should.
	if shouldAttemptDirMigration(ctx) {
		// Rename existing default datadir <home>/<Ethereum>/ to <home>/<EthereumClassic>.
		// Only do this if --datadir flag is not specified AND <home>/<EthereumClassic> does NOT already exist (only migrate once and only for defaulty).
		// If it finds an 'Ethereum' directory, it will check if it contains default ETC or ETHF chain data.
		// If it contains ETC data, it will rename the dir. If ETHF data, if will do nothing.
		if migrationError := migrateExistingDirToClassicNamingScheme(ctx); migrationError != nil {
			glog.Fatalf("%v: failed to migrate existing Classic database: %v", ErrDirectoryStructure, migrationError)
		}

		// Move existing mainnet data to pertinent chain-named subdir scheme (ie ethereum-classic/mainnet).
		// This should only happen if the given (newly defined in this protocol) subdir doesn't exist,
		// and the dirs&files (nodekey, dapp, keystore, chaindata, nodes) do exist,
		if subdirMigrateErr := migrateToChainSubdirIfNecessary(ctx); subdirMigrateErr != nil {
			glog.Fatalf("%v: failed to migrate existing data to chain-specific subdir: %v", ErrDirectoryStructure, subdirMigrateErr)
		}
	}

	// Makes sufficient configuration from JSON file or DB pending flags.
	// Delegates flag usage.
	config := mustMakeSufficientChainConfig(ctx)
	logChainConfiguration(ctx, config)

	// Configure the Ethereum service
	ethConf := mustMakeEthConf(ctx, config)

	// Configure node's service container.
	name := makeNodeName(version, ctx)
	stackConf, shhEnable := mustMakeStackConf(ctx, name, config)

	// Assemble and return the protocol stack
	stack, err := node.New(stackConf)
	if err != nil {
		glog.Fatalf("%v: failed to create the protocol stack: ", ErrStackFail, err)
	}
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return eth.New(ctx, ethConf)
	}); err != nil {
		glog.Fatalf("%v: failed to register the Ethereum service: ", ErrStackFail, err)
	}
	if shhEnable {
		if err := stack.Register(func(*node.ServiceContext) (node.Service, error) { return whisper.New(), nil }); err != nil {
			glog.Fatalf("%v: failed to register the Whisper service: ", ErrStackFail, err)
		}
	}

	// If --mlog enabled, configure and create mlog dir and file
	if ctx.GlobalString(MLogFlag.Name) != "off" {
		mustRegisterMLogsFromContext(ctx)
	} else {
		// Just demonstrative code.
		if b := logger.SetMlogEnabled(false); b == false && logger.MlogEnabled() == false {
			glog.V(logger.Warn).Infof("Machine logs: disabled")
		}
	}

	if ctx.GlobalBool(Unused1.Name) {
		glog.V(logger.Info).Infoln(fmt.Sprintf("Geth started with --%s flag, which is unused by Geth Classic and can be omitted", Unused1.Name))
	}

	return stack
}

// shouldAttemptDirMigration decides based on flags if
// should attempt to migration from old (<=3.3) directory schema to new.
func shouldAttemptDirMigration(ctx *cli.Context) bool {
	if !ctx.GlobalIsSet(aliasableName(DataDirFlag.Name, ctx)) {
		if chainVal := mustMakeChainIdentity(ctx); chainIdentitiesMain[chainVal] || chainIdentitiesMorden[chainVal] {
			return true
		}
	}
	return false
}

func mustMakeStackConf(ctx *cli.Context, name string, config *core.SufficientChainConfig) (stackConf *node.Config, shhEnable bool) {
	// Configure the node's service container
	stackConf = &node.Config{
		DataDir:         MustMakeChainDataDir(ctx),
		PrivateKey:      MakeNodeKey(ctx),
		Name:            name,
		NoDiscovery:     ctx.GlobalBool(aliasableName(NoDiscoverFlag.Name, ctx)),
		BootstrapNodes:  config.ParsedBootstrap,
		ListenAddr:      MakeListenAddress(ctx),
		NAT:             MakeNAT(ctx),
		MaxPeers:        ctx.GlobalInt(aliasableName(MaxPeersFlag.Name, ctx)),
		MaxPendingPeers: ctx.GlobalInt(aliasableName(MaxPendingPeersFlag.Name, ctx)),
		IPCPath:         MakeIPCPath(ctx),
		HTTPHost:        MakeHTTPRpcHost(ctx),
		HTTPPort:        ctx.GlobalInt(aliasableName(RPCPortFlag.Name, ctx)),
		HTTPCors:        ctx.GlobalString(aliasableName(RPCCORSDomainFlag.Name, ctx)),
		HTTPModules:     MakeRPCModules(ctx.GlobalString(aliasableName(RPCApiFlag.Name, ctx))),
		WSHost:          MakeWSRpcHost(ctx),
		WSPort:          ctx.GlobalInt(aliasableName(WSPortFlag.Name, ctx)),
		WSOrigins:       ctx.GlobalString(aliasableName(WSAllowedOriginsFlag.Name, ctx)),
		WSModules:       MakeRPCModules(ctx.GlobalString(aliasableName(WSApiFlag.Name, ctx))),
	}

<<<<<<< HEAD:cmd/ged/flag.go
	// Configure the Earthdollar service
	accman := MakeAccountManager(ctx)
glog.V(logger.Info).Infoln(fmt.Sprintf("-------------ctx:", ctx))
	edConf := &ed.Config{
		ChainConfig:             MustMakeChainConfig(ctx),
		FastSync:                ctx.GlobalBool(FastSyncFlag.Name),
		BlockChainVersion:       ctx.GlobalInt(BlockchainVersionFlag.Name),
		DatabaseCache:           ctx.GlobalInt(CacheFlag.Name),
		DatabaseHandles:         MakeDatabaseHandles(),
		NetworkId:               ctx.GlobalInt(NetworkIdFlag.Name), 
		AccountManager:          accman,
		Earthbase:               MakeEarthbase(accman, ctx),
		MinerThreads:            ctx.GlobalInt(MinerThreadsFlag.Name),
		NatSpec:                 ctx.GlobalBool(NatspecEnabledFlag.Name),
		DocRoot:                 ctx.GlobalString(DocRootFlag.Name),
		GasPrice:                common.String2Big(ctx.GlobalString(GasPriceFlag.Name)),
		GpoMinGasPrice:          common.String2Big(ctx.GlobalString(GpoMinGasPriceFlag.Name)),
		GpoMaxGasPrice:          common.String2Big(ctx.GlobalString(GpoMaxGasPriceFlag.Name)),
		GpoFullBlockRatio:       ctx.GlobalInt(GpoFullBlockRatioFlag.Name),
		GpobaseStepDown:         ctx.GlobalInt(GpobaseStepDownFlag.Name),
		GpobaseStepUp:           ctx.GlobalInt(GpobaseStepUpFlag.Name),
		GpobaseCorrectionFactor: ctx.GlobalInt(GpobaseCorrectionFactorFlag.Name),
		SolcPath:                ctx.GlobalString(SolcPathFlag.Name),
		AutoDAG:                 ctx.GlobalBool(AutoDAGFlag.Name) || ctx.GlobalBool(MiningEnabledFlag.Name),
=======
	// Configure the Whisper service
	shhEnable = ctx.GlobalBool(aliasableName(WhisperEnabledFlag.Name, ctx))

	// Override any default configs in dev mode
	if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
		if !ctx.GlobalIsSet(aliasableName(MaxPeersFlag.Name, ctx)) {
			stackConf.MaxPeers = 0
		}
		// From p2p/server.go:
		// If the port is zero, the operating system will pick a port. The
		// ListenAddr field will be updated with the actual address when
		// the server is started.
		if !ctx.GlobalIsSet(aliasableName(ListenPortFlag.Name, ctx)) {
			stackConf.ListenAddr = ":0"
		}
		if !ctx.GlobalIsSet(aliasableName(WhisperEnabledFlag.Name, ctx)) {
			shhEnable = true
		}
	}

	return stackConf, shhEnable
}

func mustMakeEthConf(ctx *cli.Context, sconf *core.SufficientChainConfig) *eth.Config {

	accman := MakeAccountManager(ctx)
	passwords := MakePasswordList(ctx)

	accounts := strings.Split(ctx.GlobalString(aliasableName(UnlockedAccountFlag.Name, ctx)), ",")
	for i, account := range accounts {
		if trimmed := strings.TrimSpace(account); trimmed != "" {
			unlockAccount(ctx, accman, trimmed, i, passwords)
		}
	}

	ethConf := &eth.Config{
		ChainConfig:             sconf.ChainConfig,
		Genesis:                 sconf.Genesis,
		FastSync:                ctx.GlobalBool(aliasableName(FastSyncFlag.Name, ctx)),
		BlockChainVersion:       ctx.GlobalInt(aliasableName(BlockchainVersionFlag.Name, ctx)),
		DatabaseCache:           ctx.GlobalInt(aliasableName(CacheFlag.Name, ctx)),
		DatabaseHandles:         MakeDatabaseHandles(),
		NetworkId:               sconf.Network,
		AccountManager:          accman,
		Etherbase:               MakeEtherbase(accman, ctx),
		MinerThreads:            ctx.GlobalInt(aliasableName(MinerThreadsFlag.Name, ctx)),
		NatSpec:                 ctx.GlobalBool(aliasableName(NatspecEnabledFlag.Name, ctx)),
		DocRoot:                 ctx.GlobalString(aliasableName(DocRootFlag.Name, ctx)),
		GasPrice:                new(big.Int),
		GpoMinGasPrice:          new(big.Int),
		GpoMaxGasPrice:          new(big.Int),
		GpoFullBlockRatio:       ctx.GlobalInt(aliasableName(GpoFullBlockRatioFlag.Name, ctx)),
		GpobaseStepDown:         ctx.GlobalInt(aliasableName(GpobaseStepDownFlag.Name, ctx)),
		GpobaseStepUp:           ctx.GlobalInt(aliasableName(GpobaseStepUpFlag.Name, ctx)),
		GpobaseCorrectionFactor: ctx.GlobalInt(aliasableName(GpobaseCorrectionFactorFlag.Name, ctx)),
		SolcPath:                ctx.GlobalString(aliasableName(SolcPathFlag.Name, ctx)),
		AutoDAG:                 ctx.GlobalBool(aliasableName(AutoDAGFlag.Name, ctx)) || ctx.GlobalBool(aliasableName(MiningEnabledFlag.Name, ctx)),
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

<<<<<<< HEAD:cmd/ged/flag.go
	// Override any default configs in dev mode or the test net
	switch {
	case ctx.GlobalBool(OlympicFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			edConf.NetworkId = 1
		}
		edConf.Genesis = core.OlympicGenesis

	case ctx.GlobalBool(TestNetFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			edConf.NetworkId = 2
		}
		edConf.Genesis = core.TestNetGenesis
		state.StartingNonce = 1048576 // (2**20)

	case ctx.GlobalBool(DevModeFlag.Name):
		// Override the base network stack configs
		if !ctx.GlobalIsSet(DataDirFlag.Name) {
			stackConf.DataDir = filepath.Join(os.TempDir(), "/earthdollar_dev_mode")
=======
	if _, ok := ethConf.GasPrice.SetString(ctx.GlobalString(aliasableName(GasPriceFlag.Name, ctx)), 0); !ok {
		log.Fatalf("malformed %s flag value %q", aliasableName(GasPriceFlag.Name, ctx), ctx.GlobalString(aliasableName(GasPriceFlag.Name, ctx)))
	}
	if _, ok := ethConf.GpoMinGasPrice.SetString(ctx.GlobalString(aliasableName(GpoMinGasPriceFlag.Name, ctx)), 0); !ok {
		log.Fatalf("malformed %s flag value %q", aliasableName(GpoMinGasPriceFlag.Name, ctx), ctx.GlobalString(aliasableName(GpoMinGasPriceFlag.Name, ctx)))
	}
	if _, ok := ethConf.GpoMaxGasPrice.SetString(ctx.GlobalString(aliasableName(GpoMaxGasPriceFlag.Name, ctx)), 0); !ok {
		log.Fatalf("malformed %s flag value %q", aliasableName(GpoMaxGasPriceFlag.Name, ctx), ctx.GlobalString(aliasableName(GpoMaxGasPriceFlag.Name, ctx)))
	}

	switch sconf.Consensus {
	case "ethash-test":
		ethConf.PowTest = true
	}

	// Override any default configs in dev mode
	if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
		// Override the Ethereum protocol configs
		if !ctx.GlobalIsSet(aliasableName(GasPriceFlag.Name, ctx)) {
			ethConf.GasPrice = new(big.Int)
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
		}
	}

	return ethConf
}

// mustMakeSufficientChainConfig makes a sufficent chain configuration (id, chainconfig, nodes,...)
// based on --chain or defaults or fails hard.
// - User must provide a full and complete config file if any is specified located at /custom/chain.json
// - Note: Function reads custom config file each time it is called; this could be altered if desired, but I (whilei) felt
// reading a file a couple of times was more efficient than storing a global.
func mustMakeSufficientChainConfig(ctx *cli.Context) *core.SufficientChainConfig {

	if cacheChainConfig != nil {
		return cacheChainConfig
	}

	config := &core.SufficientChainConfig{}
	defer func() {
		// Allow flags to override external config file.
		if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
			config.Consensus = "ethash-test"
		}
		if ctx.GlobalIsSet(aliasableName(BootnodesFlag.Name, ctx)) {
			config.ParsedBootstrap = MakeBootstrapNodesFromContext(ctx)
			glog.V(logger.Warn).Infof(`WARNING: overwriting external bootnodes configuration with those from --%s flag. Value set from flag: %v`, aliasableName(BootnodesFlag.Name, ctx), config.ParsedBootstrap)
		}
<<<<<<< HEAD:cmd/ged/flag.go
		// Override the Earthdollar protocol configs
		edConf.Genesis = core.OlympicGenesis
		if !ctx.GlobalIsSet(GasPriceFlag.Name) {
			edConf.GasPrice = new(big.Int)
=======
		if ctx.GlobalIsSet(aliasableName(NetworkIdFlag.Name, ctx)) {
			i := ctx.GlobalInt(aliasableName(NetworkIdFlag.Name, ctx))
			glog.V(logger.Warn).Infof(`WARNING: overwriting external network id configuration with that from --%s flag. Value set from flag: %d`, aliasableName(NetworkIdFlag.Name, ctx), i)
			if i < 1 {
				glog.Fatalf("Network ID cannot be less than 1. Got: %d", i)
			}
			config.Network = i
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
		}
		cacheChainConfig = config
	}()

	chainIdentity := mustMakeChainIdentity(ctx)

	// If chain identity is either of defaults (via config file or flag), use defaults.
	if chainIdentitiesMain[chainIdentity] || chainIdentitiesMorden[chainIdentity] {
		// Initialise chain configuration before handling migrations or setting up node.
		config.Identity = chainIdentity
		config.Name = mustMakeChainConfigNameDefaulty(ctx)
		config.Network = eth.NetworkId // 1, default mainnet
		config.Consensus = "ethash"
		config.Genesis = core.DefaultConfigMainnet.Genesis
		config.ChainConfig = MustMakeChainConfigFromDefaults(ctx).SortForks()
		config.ParsedBootstrap = MakeBootstrapNodesFromContext(ctx)
		if chainIsMorden(ctx) {
			config.Network = 2
			config.Genesis = core.DefaultConfigMorden.Genesis
			state.StartingNonce = state.DefaultTestnetStartingNonce // (2**20)
		}
<<<<<<< HEAD:cmd/ged/flag.go
		edConf.PowTest = true
=======
		return config
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

	// Returns surely valid suff chain config.
	chainDir := MustMakeChainDataDir(ctx)
	defaultChainConfigPath := filepath.Join(chainDir, "chain.json")
	if _, de := os.Stat(defaultChainConfigPath); de != nil && os.IsNotExist(de) {
		glog.Fatalf(`%v: %v
		It looks like you haven't set up your custom chain yet...
		Here's a possible workflow for that:

		$ geth --chain morden dump-chain-config %v/chain.json
		$ sed -i.bak s/morden/%v/ %v/chain.json
		$ vi %v/chain.json # <- make your customizations
		`, core.ErrChainConfigNotFound, defaultChainConfigPath,
			chainDir, chainIdentity, chainDir, chainDir)
	}
<<<<<<< HEAD:cmd/ged/flag.go
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return ed.New(ctx, edConf)
	}); err != nil {
		log.Fatal("Failed to register the Earthdollar service: ", err)
=======
	config, err := core.ReadExternalChainConfigFromFile(defaultChainConfigPath)
	if err != nil {
		glog.Fatalf(`invalid external configuration JSON: '%v': %v
		Valid custom configuration JSON file must be named 'chain.json' and be located in respective chain subdir.`, defaultChainConfigPath, err)
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

	// Ensure JSON 'id' value matches name of parent chain subdir.
	if config.Identity != chainIdentity {
		glog.Fatalf(`%v: JSON 'id' value in external config file (%v) must match name of parent subdir (%v)`, ErrInvalidChainID, config.Identity, chainIdentity)
	}

<<<<<<< HEAD:cmd/ged/flag.go
	if ctx.GlobalBool(Unused1.Name) {
		glog.V(logger.Info).Infoln(fmt.Sprintf("Ged started with --%s flag, which is unused by Ged and can be omitted", Unused1.Name))
=======
	// Set statedb StartingNonce from external config, if specified (is optional)
	if config.State != nil {
		if sn := config.State.StartingNonce; sn != 0 {
			state.StartingNonce = sn
		}
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

	return config
}

<<<<<<< HEAD:cmd/ged/flag.go
// SetupNetwork configures the system for either the main net or some test network.
func SetupNetwork(ctx *cli.Context) {
	switch {
	case ctx.GlobalBool(OlympicFlag.Name):
		core.DurationLimit = big.NewInt(8)
		core.MinGasLimit = big.NewInt(125000)
		types.HeaderExtraMax = 1024
		NetworkIdFlag.Value = 0
		core.BlockReward = big.NewInt(1.5e+18)
		core.ExpDiffPeriod = big.NewInt(math.MaxInt64)
	}
	core.TargetGasLimit = common.String2Big(ctx.GlobalString(TargetGasLimitFlag.Name))
}

// MustMakeChainConfig reads the chain configuration from the database in ctx.Datadir.
func MustMakeChainConfig(ctx *cli.Context) *core.ChainConfig {  
 	db := MakeChainDatabase(ctx)
	defer db.Close()
=======
func logChainConfiguration(ctx *cli.Context, config *core.SufficientChainConfig) {

	chainIdentity := mustMakeChainIdentity(ctx)
	chainIsCustom := !(chainIdentitiesMain[chainIdentity] || chainIdentitiesMorden[chainIdentity])
	if chainIsCustom {
		glog.V(logger.Info).Infof("Using custom chain configuration: \x1b[32m%s\x1b[39m", chainIdentity)
	}

	glog.V(logger.Info).Info(glog.Separator("-"))
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go

	glog.V(logger.Info).Infof("Starting Geth Classic \x1b[32m%s\x1b[39m", ctx.App.Version)
	glog.V(logger.Info).Infof("Geth is configured to use ETC blockchain: \x1b[32m%v\x1b[39m", config.Name)
	glog.V(logger.Info).Infof("Using chain database at: \x1b[32m%s\x1b[39m", MustMakeChainDataDir(ctx)+"/chaindata")

<<<<<<< HEAD:cmd/ged/flag.go
// MustMakeChainConfigFromDb reads the chain configuration from the given database.
func MustMakeChainConfigFromDb(ctx *cli.Context, db eddb.Database) *core.ChainConfig {  
	c := core.DefaultConfig
	if ctx.GlobalBool(TestNetFlag.Name) {
		c = core.TestConfig
	}

	/*for i := range c.Forks {  //earthdollar, remove
		// Force override any existing configs if explicitly requested
		if c.Forks[i].Name == "ETF" {
			if ctx.GlobalBool(ETFChain.Name) {
				c.Forks[i].Support = true
=======
	glog.V(logger.Info).Infof("%v blockchain upgrades associated with this configuration:", len(config.ChainConfig.Forks))

	for i := range config.ChainConfig.Forks {
		glog.V(logger.Info).Infof(" %7v %v", config.ChainConfig.Forks[i].Block, config.ChainConfig.Forks[i].Name)
		if !config.ChainConfig.Forks[i].RequiredHash.IsEmpty() {
			glog.V(logger.Info).Infof("         with block %v", config.ChainConfig.Forks[i].RequiredHash.Hex())
		}
		for _, feat := range config.ChainConfig.Forks[i].Features {
			glog.V(logger.Debug).Infof("    id: %v", feat.ID)
			for k, v := range feat.Options {
				glog.V(logger.Debug).Infof("        %v: %v", k, v)
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
			}
		}
	}*/

<<<<<<< HEAD:cmd/ged/flag.go
	separator := strings.Repeat("-", 110)
	glog.V(logger.Warn).Info(separator)
	glog.V(logger.Warn).Info(fmt.Sprintf("Starting Ged \x1b[32m%s\x1b[39m", ctx.App.Version))

	genesis := core.GetBlock(db, core.GetCanonicalHash(db, 0))

	genesisHash := ""
	if genesis != nil {
		genesisHash = genesis.Hash().Hex()
=======
	if chainIsCustom {
		glog.V(logger.Info).Infof("State starting nonce: %v", colorGreen(state.StartingNonce))
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

<<<<<<< HEAD:cmd/ged/flag.go
	//netsplitChoice := ""
	/*for i := range c.Forks {
		if c.Forks[i].NetworkSplit {
			netsplitChoice = fmt.Sprintf("resulted in a network split (support: %t)", c.Forks[i].Support)
		} else {
			netsplitChoice = ""
		}
		glog.V(logger.Warn).Info(fmt.Sprintf(" %7v %v hard-fork %v", c.Forks[i].Block, c.Forks[i].Name, netsplitChoice))
	}*/

	if ctx.GlobalBool(TestNetFlag.Name) {
		glog.V(logger.Warn).Info("Ged is configured to use the \x1b[33mEarthdollar Testnet\x1b[39m blockchain!")
	} else {
		glog.V(logger.Warn).Info("Ged is configured to use the \x1b[32mEarthdollar\x1b[39m blockchain!")
=======
	glog.V(logger.Info).Info(glog.Separator("-"))
}

// MustMakeChainConfigFromDefaults reads the chain configuration from hardcode.
func MustMakeChainConfigFromDefaults(ctx *cli.Context) *core.ChainConfig {
	c := core.DefaultConfigMainnet.ChainConfig
	if chainIsMorden(ctx) {
		c = core.DefaultConfigMorden.ChainConfig
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}
	return c
}

// MakeChainDatabase open an LevelDB using the flags passed to the client and will hard crash if it fails.
func MakeChainDatabase(ctx *cli.Context) eddb.Database {
	var (
		datadir = MustMakeChainDataDir(ctx)
		cache   = ctx.GlobalInt(aliasableName(CacheFlag.Name, ctx))
		handles = MakeDatabaseHandles()
	)
	chainDb, err := eddb.NewLDBDatabase(filepath.Join(datadir, "chaindata"), cache, handles)
	if err != nil {
		glog.Fatal("Could not open database: ", err)
	}
	return chainDb
}

// MakeChain creates a chain manager from set command line flags.
func MakeChain(ctx *cli.Context) (chain *core.BlockChain, chainDb eddb.Database) {
	var err error
	sconf := mustMakeSufficientChainConfig(ctx)
	chainDb = MakeChainDatabase(ctx)

<<<<<<< HEAD:cmd/ged/flag.go
	if ctx.GlobalBool(OlympicFlag.Name) {
		_, err := core.WriteGenesisBlock(chainDb, core.OlympicGenesis)
		if err != nil {
			log.Fatal(err)
		}
	}
	chainConfig := MustMakeChainConfigFromDb(ctx, chainDb)

	pow := pow.PoW(core.FakePow{})
	if !ctx.GlobalBool(FakePoWFlag.Name) {
		pow = edhash.New()
=======
	pow := pow.PoW(core.FakePow{})
	if !ctx.GlobalBool(aliasableName(FakePoWFlag.Name, ctx)) {
		pow = ethash.New()
	} else {
		glog.V(logger.Info).Info("Consensus: fake")
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/flag.go
	}

	chain, err = core.NewBlockChain(chainDb, sconf.ChainConfig, pow, new(event.TypeMux))
	if err != nil {
		glog.Fatal("Could not start chainmanager: ", err)
	}
	return chain, chainDb
}

// MakeConsolePreloads retrieves the absolute paths for the console JavaScript
// scripts to preload before starting.
func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.GlobalString(aliasableName(PreloadJSFlag.Name, ctx)) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	preloads := []string{}

	assets := ctx.GlobalString(aliasableName(JSpathFlag.Name, ctx))
	for _, file := range strings.Split(ctx.GlobalString(aliasableName(PreloadJSFlag.Name, ctx)), ",") {
		preloads = append(preloads, common.EnsurePathAbsoluteOrRelativeTo(assets, strings.TrimSpace(file)))
	}
	return preloads
}
