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

// Contains a simple library definition to allow creating a Ged instance from
// straight C code.

package main

// #ifdef __cplusplus
// extern "C" {
// #endif
//
// extern int run(const char*);
//
// #ifdef __cplusplus
// }
// #endif
import "C"
import (
	"fmt"
	"os"
	"strings"
)

//export doRun
func doRun(args *C.char) C.int {
	// This is equivalent to ged.main, just modified to handle the function arg passing
	if err := app.Run(strings.Split("ged "+C.GoString(args), " ")); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return -1
	}
	return 0
}
