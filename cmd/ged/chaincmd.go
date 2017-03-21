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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

<<<<<<< HEAD:cmd/ged/chaincmd.go
	"github.com/Tzunami/go-earthdollar/cmd/utils"
	"github.com/Tzunami/go-earthdollar/common"
	"github.com/Tzunami/go-earthdollar/console"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/core/state"
	"github.com/Tzunami/go-earthdollar/core/types"
	"github.com/Tzunami/go-earthdollar/logger/glog"
=======
	"github.com/Tzunami/go-earthdollar/common"
	"github.com/Tzunami/go-earthdollar/console"
	"github.com/Tzunami/go-earthdollar/core"
	"github.com/Tzunami/go-earthdollar/core/state"
	"github.com/Tzunami/go-earthdollar/core/types"
	"github.com/Tzunami/go-earthdollar/logger/glog"
>>>>>>> 09218adc3dc58c6d349121f8b1c0cf0b62331087:cmd/ged/chaincmd.go
	"gopkg.in/urfave/cli.v1"
)

var (
	importCommand = cli.Command{
		Action: importChain,
		Name:   "import",
		Usage:  `import a blockchain file`,
	}
	exportCommand = cli.Command{
		Action: exportChain,
		Name:   "export",
		Usage:  `export blockchain into file`,
		Description: `
Requires a first argument of the file to write to.
Optional second and third arguments control the first and
last block to write. In this mode, the file will be appended
if already existing.
		`,
	}
	upgradedbCommand = cli.Command{
		Action: upgradeDB,
		Name:   "upgradedb",
		Usage:  "upgrade chainblock database",
	}
	removedbCommand = cli.Command{
		Action: removeDB,
		Name:   "removedb",
		Usage:  "Remove blockchain and state databases",
	}
	dumpCommand = cli.Command{
		Action: dump,
		Name:   "dump",
		Usage:  `dump a specific block from storage`,
		Description: `
The arguments are interpreted as block numbers or hashes.
Use "earthdollar dump 0" to dump the genesis block.
`,
	}
)

func importChain(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		log.Fatal("This command requires an argument.")
	}
	chain, chainDb := MakeChain(ctx)
	start := time.Now()
	err := ImportChain(chain, ctx.Args().First())
	chainDb.Close()
	if err != nil {
		log.Fatal("Import error: ", err)
	}
	fmt.Printf("Import done in %v", time.Since(start))
	return nil
}

func exportChain(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		log.Fatal("This command requires an argument.")
	}
	chain, _ := MakeChain(ctx)
	start := time.Now()

	fp := ctx.Args().First()
	if len(ctx.Args()) < 3 {
		if err := ExportChain(chain, fp); err != nil {
			log.Fatal(err)
		}
	} else {
		// This can be improved to allow for numbers larger than 9223372036854775807
		first, err := strconv.ParseUint(ctx.Args().Get(1), 10, 64)
		if err != nil {
			log.Fatal("export paramater: ", err)
		}
		last, err := strconv.ParseUint(ctx.Args().Get(2), 10, 64)
		if err != nil {
			log.Fatal("export paramater: ", err)
		}
		if err = ExportAppendChain(chain, fp, first, last); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Export done in %v", time.Since(start))
	return nil
}

func removeDB(ctx *cli.Context) error {
	confirm, err := console.Stdin.PromptConfirm("Remove local database?")
	if err != nil {
		log.Fatal(err)
	}

	if confirm {
		fmt.Println("Removing chaindata...")
		start := time.Now()

		os.RemoveAll(filepath.Join(ctx.GlobalString(DataDirFlag.Name), "chaindata"))

		fmt.Printf("Removed in %v\n", time.Since(start))
	} else {
		fmt.Println("Operation aborted")
	}
	return nil
}

func upgradeDB(ctx *cli.Context) error {
	glog.Infoln("Upgrading blockchain database")

	chain, chainDb := MakeChain(ctx)
	bcVersion := core.GetBlockChainVersion(chainDb)
	if bcVersion == 0 {
		bcVersion = core.BlockChainVersion
	}

	// Export the current chain.
	filename := fmt.Sprintf("blockchain_%d_%s.chain", bcVersion, time.Now().Format("20060102_150405"))
	exportFile := filepath.Join(ctx.GlobalString(DataDirFlag.Name), filename)
	if err := ExportChain(chain, exportFile); err != nil {
		log.Fatal("Unable to export chain for reimport ", err)
	}
	chainDb.Close()
	os.RemoveAll(filepath.Join(ctx.GlobalString(DataDirFlag.Name), "chaindata"))

	// Import the chain file.
	chain, chainDb = MakeChain(ctx)
	core.WriteBlockChainVersion(chainDb, core.BlockChainVersion)
	err := ImportChain(chain, exportFile)
	chainDb.Close()
	if err != nil {
		log.Fatalf("Import error %v (a backup is made in %s, use the import command to import it)", err, exportFile)
	} else {
		os.Remove(exportFile)
		glog.Infoln("Import finished")
	}
	return nil
}

func dump(ctx *cli.Context) error {
	chain, chainDb := MakeChain(ctx)
	for _, arg := range ctx.Args() {
		var block *types.Block
		if hashish(arg) {
			block = chain.GetBlock(common.HexToHash(arg))
		} else {
			num, _ := strconv.Atoi(arg)
			block = chain.GetBlockByNumber(uint64(num))
		}
		if block == nil {
			fmt.Println("{}")
			log.Fatal("block not found")
		} else {
			state, err := state.New(block.Root(), chainDb)
			if err != nil {
				log.Fatal("could not create new state: ", err)
			}
			fmt.Printf("%s\n", state.Dump())
		}
	}
	chainDb.Close()
	return nil
}

// hashish returns true for strings that look like hashes.
func hashish(x string) bool {
	_, err := strconv.Atoi(x)
	return err != nil
}
