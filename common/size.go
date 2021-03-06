// Copyright 2014 The go-earthdollar Authors
// This file is part of the go-earthdollar library.
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

package common

import (
	"fmt"
	"math/big"
)

type StorageSize float64

func (self StorageSize) String() string {
	if self > 1000000 {
		return fmt.Sprintf("%.2f mB", self/1000000)
	} else if self > 1000 {
		return fmt.Sprintf("%.2f kB", self/1000)
	} else {
		return fmt.Sprintf("%.2f B", self)
	}
}

func (self StorageSize) Int64() int64 {
	return int64(self)
}

// The different number of units
var (
	Douglas  = BigPow(10, 42)
	Einstein = BigPow(10, 21)
	Ed    = BigPow(10, 18)
	Kam   = BigPow(10, 15)
	Rajpal    = BigPow(10, 12)
	Shannon  = BigPow(10, 9)
	Babbage  = BigPow(10, 6)
	Ada      = BigPow(10, 3)
	Seed      = big.NewInt(1)
)

//
// Currency to string
// Returns a string representing a human readable format
func CurrencyToString(num *big.Int) string {
	var (
		fin   *big.Int = num
		denom string   = "Seed"
	)

	switch {
	case num.Cmp(Ed) >= 0:
		fin = new(big.Int).Div(num, Ed)
		denom = "Ed"
	case num.Cmp(Kam) >= 0:
		fin = new(big.Int).Div(num, Kam)
		denom = "Kam"
	case num.Cmp(Rajpal) >= 0:
		fin = new(big.Int).Div(num, Rajpal)
		denom = "Rajpal"
	case num.Cmp(Shannon) >= 0:
		fin = new(big.Int).Div(num, Shannon)
		denom = "Shannon"
	case num.Cmp(Babbage) >= 0:
		fin = new(big.Int).Div(num, Babbage)
		denom = "Babbage"
	case num.Cmp(Ada) >= 0:
		fin = new(big.Int).Div(num, Ada)
		denom = "Ada"
	}

	// TODO add comment clarifying expected behavior
	if len(fin.String()) > 5 {
		return fmt.Sprintf("%sE%d %s", fin.String()[0:5], len(fin.String())-5, denom)
	}

	return fmt.Sprintf("%v %s", fin, denom)
}
