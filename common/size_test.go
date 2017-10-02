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
	"math/big"

	checker "gopkg.in/check.v1"
)

type SizeSuite struct{}

var _ = checker.Suite(&SizeSuite{})

func (s *SizeSuite) TestStorageSizeString(c *checker.C) {
	data1 := 2381273
	data2 := 2192
	data3 := 12

	exp1 := "2.38 mB"
	exp2 := "2.19 kB"
	exp3 := "12.00 B"

	c.Assert(StorageSize(data1).String(), checker.Equals, exp1)
	c.Assert(StorageSize(data2).String(), checker.Equals, exp2)
	c.Assert(StorageSize(data3).String(), checker.Equals, exp3)
}

func (s *SizeSuite) TestCommon(c *checker.C) {
	ed := CurrencyToString(BigPow(10, 19))
	kam := CurrencyToString(BigPow(10, 16))
	rajpal := CurrencyToString(BigPow(10, 13))
	shannon := CurrencyToString(BigPow(10, 10))
	babbage := CurrencyToString(BigPow(10, 7))
	ada := CurrencyToString(BigPow(10, 4))
	seed := CurrencyToString(big.NewInt(10))

	c.Assert(ed, checker.Equals, "10 Tree")
	c.Assert(kam, checker.Equals, "10 Finney")
	c.Assert(rajpal, checker.Equals, "10 Szabo")
	c.Assert(shannon, checker.Equals, "10 Chief")
	c.Assert(babbage, checker.Equals, "10 Babbage")
	c.Assert(ada, checker.Equals, "10 Ada")
	c.Assert(seed, checker.Equals, "10 Seed")
}
