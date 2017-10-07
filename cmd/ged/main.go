// Copyright 2014 The go-earthdollar Authors
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

// ged is the official command-line client for Earthdollar.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/urfave/cli.v1"

<<<<<<< HEAD:cmd/ged/main.go
	"github.com/Tzunami/edhash"
	"github.com/Tzunami/go-earthdollar/console"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/ed"
	"github.com/Tzunami/go-earthdollar/eddb"
	"github.com/Tzunami/go-earthdollar/logger"
	"github.com/Tzunami/go-earthdollar/logger/glog"
	"github.com/Tzunami/go-earthdollar/metrics"
	"github.com/Tzunami/go-earthdollar/node"
=======
	"github.com/ethereumproject/go-ethereum/console"
	"github.com/ethereumproject/go-ethereum/core"
	"github.com/ethereumproject/go-ethereum/eth"
	"github.com/ethereumproject/go-ethereum/logger"
	"github.com/ethereumproject/go-ethereum/logger/glog"
	"github.com/ethereumproject/go-ethereum/metrics"
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/main.go
)

// Version is the application revision identifier. It can be set with the linker
// as in: go build -ldflags "-X main.Version="`git describe --tags`
var Version = "source"

func makeCLIApp() (app *cli.App) {
	app = cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = Version
	app.Usage = "the go-earthdollar command line interface"
	app.Action = ged
	app.HideVersion = true // we have a command to print the version

	app.Commands = []cli.Command{
		importCommand,
		exportCommand,
		dumpChainConfigCommand,
		upgradedbCommand,
		removedbCommand,
		dumpCommand,
		rollbackCommand,
		monitorCommand,
		accountCommand,
		walletCommand,
		consoleCommand,
		attachCommand,
		javascriptCommand,
		statusCommand,
		{
<<<<<<< HEAD:cmd/ged/main.go
			Action: makedag,
			Name:   "makedag",
			Usage:  "generate edhash dag (for testing)",
=======
			Action:  makedag,
			Name:    "make-dag",
			Aliases: []string{"makedag"},
			Usage:   "Generate ethash dag (for testing)",
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/main.go
			Description: `
The makedag command generates an edhash DAG in /tmp/dag.

This command exists to support the system testing project.
Regular users do not need to execute it.
`,
		},
		{
			Action:  gpuinfo,
			Name:    "gpu-info",
			Aliases: []string{"gpuinfo"},
			Usage:   "GPU info",
			Description: `
Prints OpenCL device info for all found GPUs.
`,
		},
		{
			Action:  gpubench,
			Name:    "gpu-bench",
			Aliases: []string{"gpubench"},
			Usage:   "Benchmark GPU",
			Description: `
Runs quick benchmark on first GPU found.
`,
		},
		{
			Action: version,
			Name:   "version",
<<<<<<< HEAD:cmd/ged/main.go
			Usage:  "print earthdollar version numbers",
=======
			Usage:  "Print ethereum version numbers",
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/main.go
			Description: `
The output of this command is supposed to be machine-readable.
`,
		},
		{
			Action: makeMLogDocumentation,
			Name:   "mdoc",
			Usage:  "Generate mlog documentation",
			Description: `
Auto-generates documentation for all available mlog lines.
Use -md switch to toggle markdown output (eg. for wiki).
Arguments may be used to specify exclusive candidate components;
so 'geth mdoc -md discover' will generate markdown documentation only
for the 'discover' component.
`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "md",
					Usage: "Toggle markdown formatting",
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		NodeNameFlag,
		UnlockedAccountFlag,
		PasswordFileFlag,
		AccountsIndexFlag,
		BootnodesFlag,
		DataDirFlag,
		DocRootFlag,
		KeyStoreDirFlag,
		ChainIdentityFlag,
		BlockchainVersionFlag,
		FastSyncFlag,
		CacheFlag,
		LightKDFFlag,
		JSpathFlag,
		ListenPortFlag,
		MaxPeersFlag,
		MaxPendingPeersFlag,
		EarthbaseFlag,
		GasPriceFlag,
		MinerThreadsFlag,
		MiningEnabledFlag,
		MiningGPUFlag,
		AutoDAGFlag,
		TargetGasLimitFlag,
		NATFlag,
		NatspecEnabledFlag,
		NoDiscoverFlag,
		NodeKeyFileFlag,
		NodeKeyHexFlag,
		RPCEnabledFlag,
		RPCListenAddrFlag,
		RPCPortFlag,
		RPCApiFlag,
		WSEnabledFlag,
		WSListenAddrFlag,
		WSPortFlag,
		WSApiFlag,
		WSAllowedOriginsFlag,
		IPCDisabledFlag,
		IPCApiFlag,
		IPCPathFlag,
		ExecFlag,
		PreloadJSFlag,
		WhisperEnabledFlag,
		DevModeFlag,
		TestNetFlag,
		NetworkIdFlag,
		RPCCORSDomainFlag,
		VerbosityFlag,
		VModuleFlag,
		LogDirFlag,
		LogStatusFlag,
		MLogFlag,
		MLogDirFlag,
		MLogComponentsFlag,
		BacktraceAtFlag,
		MetricsFlag,
		FakePoWFlag,
		SolcPathFlag,
		GpoMinGasPriceFlag,
		GpoMaxGasPriceFlag,
		GpoFullBlockRatioFlag,
		GpobaseStepDownFlag,
		GpobaseStepUpFlag,
		GpobaseCorrectionFactorFlag,
		ExtraDataFlag,
		Unused1,
	}

	app.Before = func(ctx *cli.Context) error {

		// It's a patch.
		// Don't know why urfave/cli isn't catching the unknown command on its own.
		if ctx.Args().Present() {
			commandExists := false
			for _, cmd := range app.Commands {
				if cmd.HasName(ctx.Args().First()) {
					commandExists = true
				}
			}
			if !commandExists {
				if e := cli.ShowCommandHelp(ctx, ctx.Args().First()); e != nil {
					return e
				}
			}
		}

		runtime.GOMAXPROCS(runtime.NumCPU())

		glog.CopyStandardLogTo("INFO")

		if ctx.GlobalIsSet(aliasableName(LogDirFlag.Name, ctx)) {
			if p := ctx.GlobalString(aliasableName(LogDirFlag.Name, ctx)); p != "" {
				if e := os.MkdirAll(p, os.ModePerm); e != nil {
					return e
				}
				glog.SetLogDir(p)
				glog.SetAlsoToStderr(true)
			}
		} else {
			glog.SetToStderr(true)
		}

		if s := ctx.String("metrics"); s != "" {
			go metrics.CollectToFile(s)
		}

		// This should be the only place where reporting is enabled
		// because it is not intended to run while testing.
		// In addition to this check, bad block reports are sent only
		// for chains with the main network genesis block and network id 1.
		ed.EnableBadBlockReporting = true

		// (whilei): I use `log` instead of `glog` because git diff tells me:
		// > The output of this command is supposed to be machine-readable.
		gasLimit := ctx.GlobalString(aliasableName(TargetGasLimitFlag.Name, ctx))
		if _, ok := core.TargetGasLimit.SetString(gasLimit, 0); !ok {
			log.Fatalf("malformed %s flag value %q", aliasableName(TargetGasLimitFlag.Name, ctx), gasLimit)
		}

		// Set morden chain by default for dev mode.
		if ctx.GlobalBool(aliasableName(DevModeFlag.Name, ctx)) {
			if !ctx.GlobalIsSet(aliasableName(ChainIdentityFlag.Name, ctx)) {
				if e := ctx.Set(aliasableName(ChainIdentityFlag.Name, ctx), "morden"); e != nil {
					log.Fatalf("failed to set chain value: %v", e)
				}
			}
		}

		return nil
	}

	app.After = func(ctx *cli.Context) error {
		logger.Flush()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Invalid command: %q. Please find `geth` usage below. \n", command)
		cli.ShowAppHelp(c)
		os.Exit(3)
	}
	return app
}

func main() {
	app := makeCLIApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// ged is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
<<<<<<< HEAD:cmd/ged/main.go
func ged(ctx *cli.Context) error {
	node := MakeSystemNode(Version, ctx)
	startNode(ctx, node)
	node.Wait()

	return nil
}

// initGenesis will initialise the given JSON format genesis file and writes it as
// the zero'd block (i.e. genesis) or will fail hard if it can't succeed.
func initGenesis(ctx *cli.Context) {
	path := ctx.Args().First()
	if len(path) == 0 {
		log.Fatal("need path argument to genesis JSON file")
	}

	chainDB, err := eddb.NewLDBDatabase(filepath.Join(MustMakeDataDir(ctx), "chaindata"), 0, 0)
	if err != nil {
		log.Fatal("could not open database: ", err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read genesis file: ", err)
	}
	defer f.Close()

	dump := new(core.GenesisDump)
	if json.NewDecoder(f).Decode(dump); err != nil {
		log.Fatalf("%s: %s", path, err)
	}

	block, err := core.WriteGenesisBlock(chainDB, dump)
	if err != nil {
		log.Fatal("failed to write genesis block: ", err)
	}
	log.Printf("successfully wrote genesis block and/or chain rule set: %x", block.Hash())
}

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces and the
// miner.
func startNode(ctx *cli.Context, stack *node.Node) {
	// Start up the node itself
	StartNode(stack)

	// Unlock any account specifically requested
	var earthdollar *ed.Earthdollar
	if err := stack.Service(&earthdollar); err != nil {
		log.Fatal("earthdollar service not running: ", err)
	}
	accman := earthdollar.AccountManager()
	passwords := MakePasswordList(ctx)

	accounts := strings.Split(ctx.GlobalString(UnlockedAccountFlag.Name), ",")
	for i, account := range accounts {
		if trimmed := strings.TrimSpace(account); trimmed != "" {
			unlockAccount(ctx, accman, trimmed, i, passwords)
		}
	}
	// Start auxiliary services if enabled
	if ctx.GlobalBool(MiningEnabledFlag.Name) {
		if err := earthdollar.StartMining(ctx.GlobalInt(MinerThreadsFlag.Name), ctx.GlobalString(MiningGPUFlag.Name)); err != nil {
			log.Fatalf("Failed to start mining: ", err)
		}
	}
}

func makedag(ctx *cli.Context) error {
	args := ctx.Args()
	wrongArgs := func() {
		log.Fatal(`Usage: ged makedag <block number> <outputdir>`)
	}
	switch {
	case len(args) == 2:
		blockNum, err := strconv.ParseUint(args[0], 0, 64)
		dir := args[1]
		if err != nil {
			wrongArgs()
		} else {
			dir = filepath.Clean(dir)
			// seems to require a trailing slash
			if !strings.HasSuffix(dir, "/") {
				dir = dir + "/"
			}
			_, err = ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal("Can't find dir")
			}
			fmt.Println("making DAG, this could take awhile...")
			edhash.MakeDAG(blockNum, dir)
		}
	default:
		wrongArgs()
	}
	return nil
}

func gpuinfo(ctx *cli.Context) error {
	ed.PrintOpenCLDevices()
	return nil
}

func gpubench(ctx *cli.Context) error {
	args := ctx.Args()
	wrongArgs := func() {
		log.Fatal(`Usage: ged gpubench <gpu number>`)
	}
	switch {
	case len(args) == 1:
		n, err := strconv.ParseUint(args[0], 0, 64)
		if err != nil {
			wrongArgs()
		}
		ed.GPUBench(n)
	case len(args) == 0:
		ed.GPUBench(0)
	default:
		wrongArgs()
	}
	return nil
}

func version(c *cli.Context) error {
	fmt.Println("Ged")
	fmt.Println("Version:", Version)
	fmt.Println("Protocol Versions:", ed.ProtocolVersions)
	fmt.Println("Network Id:", c.GlobalInt(NetworkIdFlag.Name))
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
=======
func geth(ctx *cli.Context) error {
	n := MakeSystemNode(Version, ctx)
	ethe := startNode(ctx, n)

	if ctx.GlobalIsSet(LogStatusFlag.Name) {
		dispatchStatusLogs(ctx, ethe)
	}

	n.Wait()
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/main.go

	return nil
}
