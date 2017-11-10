// Copyright 2017 The go-earthdollar Authors
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
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"gopkg.in/urfave/cli.v1"

	"github.com/Tzunami/go-earthdollar/node"
	"github.com/Tzunami/go-earthdollar/rpc"
)

var (
	apiCommand = cli.Command{
		Action: execAPI,
		Name:   "api",
		Usage:  "Run any API command",
		Description: `
The api command allows to run any method defined in any API module.

Each parameter should be passed as JSON representation:
- no quotations for numbers or booleans,
- strings must be correclty quoted, like '"some value"' (quotes must be
  included in string passed to application),
- complex objects could be passed as JSON string.

Examples:
$ ged api ed getBlockByNumber 123 true
$ ged ed getBlockByNumber '"latest"' true
$ ged --chain morden api ed sendTransaction '{"from": "0x396599f365093186742c17aab158bf515e978bc7", "gas": "0x5208", "gasPrice": "0x02540be400", "to": "0xa02cee0fc1d3fb4dde86b79fe93e4140671fd949"}'

Output will be in JSON format.
`,
	}
)

func execAPI(ctx *cli.Context) error {
	client, err := getClient(ctx)
	if err != nil {
		return err
	}

	if err := validateArguments(ctx, client); err != nil {
		return err
	}

	result, err := callRPC(ctx, client)
	if err != nil {
		return err
	}
	return prettyPrint(result)
}

func getClient(ctx *cli.Context) (rpc.Client, error) {
	chainDir := MustMakeChainDataDir(ctx)
	var uri = "ipc:" + node.DefaultIPCEndpoint(chainDir)
	return rpc.NewClient(uri)
}

func validateArguments(ctx *cli.Context, client rpc.Client) error {
	if len(ctx.Args()) < 2 {
		return fmt.Errorf("api command requires at least 2 arguments (module and method), %d provided",
			len(ctx.Args()))
	}
	modules, err := client.SupportedModules()
	if err != nil {
		return err
	}

	module := ctx.Args()[0]
	if _, ok := modules[module]; !ok {
		return fmt.Errorf("unknown API module: %s", module)
	}

	return nil
}

func callRPC(ctx *cli.Context, client rpc.Client) (interface{}, error) {
	var (
		module = ctx.Args()[0]
		method = ctx.Args()[1]
		args   = ctx.Args()[2:]
	)
	req := rpc.JSONRequest{
		Id:      json.RawMessage(strconv.Itoa(rand.Int())),
		Method:  module + "_" + method,
		Version: "2.0",
		Payload: json.RawMessage("[" + strings.Join(args, ",") + "]"),
	}

	if err := client.Send(req); err != nil {
		return nil, err
	}

	var res rpc.JSONResponse
	if err := client.Recv(&res); err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, fmt.Errorf("error in %s_%s: %s (code: %d)",
			module, method, res.Error.Message, res.Error.Code)
	}
	if res.Result != nil {
		return res.Result, nil
	}

	return nil, errors.New("no API response")
}

func prettyPrint(result interface{}) error {
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}
