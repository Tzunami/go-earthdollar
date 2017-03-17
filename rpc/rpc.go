<<<<<<< HEAD:internal/debug/loudpanic.go
// Copyright 2016 The go-earthdollar Authors
// This file is part of the go-earthdollar library.
=======
// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
>>>>>>> 09218adc3dc58c6d349121f8b1c0cf0b62331087:rpc/rpc.go
//
// The go-earthdollar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
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

package rpc

import (
	"fmt"
	"strings"
)

func NewClient(uri string) (Client, error) {
	if strings.HasPrefix(uri, "ipc:") {
		return newIPCClient(uri[4:])
	}
	if strings.HasPrefix(uri, "rpc:") {
		return &httpClient{endpoint: uri[4:]}, nil
	}
	if strings.HasPrefix(uri, "http://") {
		return &httpClient{endpoint: uri}, nil
	}
	if strings.HasPrefix(uri, "ws:") {
		return &wsClient{endpoint: uri}, nil
	}
	return nil, fmt.Errorf("unsupported RPC schema %q", uri)
}
