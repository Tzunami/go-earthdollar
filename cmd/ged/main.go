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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/urfave/cli.v1"

	"github.com/Tzunami/go-earthdollar/console"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/ed"
	"github.com/Tzunami/go-earthdollar/logger"
	"github.com/Tzunami/go-earthdollar/logger/glog"
	"github.com/Tzunami/go-earthdollar/metrics"
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
		apiCommand,
		{
			Action:  makedag,
			Name:    "make-dag",
			Aliases: []string{"makedag"},
			Usage:   "Generate edhash dag (for testing)",
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
			Usage:  "Print earthdollar version numbers",
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
so 'ged mdoc -md discover' will generate markdown documentation only
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

		// Check for --exec set without console OR attach
		if ctx.IsSet(ExecFlag.Name) {
			// If no command is used, OR command is not one of the valid commands attach/console
			if cmdName := ctx.Args().First(); cmdName == "" || (cmdName != "console" && cmdName != "attach") {
				log.Printf("Error: --%v flag requires use of 'attach' OR 'console' command, command was: '%v'", ExecFlag.Name, cmdName)
				cli.ShowCommandHelp(ctx, consoleCommand.Name)
				cli.ShowCommandHelp(ctx, attachCommand.Name)
				os.Exit(1)
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
		fmt.Fprintf(c.App.Writer, "Invalid command: %q. Please find `ged` usage below. \n", command)
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
func ged(ctx *cli.Context) error {
	n := MakeSystemNode(Version, ctx)
	ede := startNode(ctx, n)

	if ctx.GlobalIsSet(LogStatusFlag.Name) {
		dispatchStatusLogs(ctx, ede)
	}

	n.Wait()

	return nil
}
