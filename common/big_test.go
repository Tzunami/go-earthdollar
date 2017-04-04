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
	"testing"
)

func TestMisc(t *testing.T) {
	a := Big("10")
	b := Big("57896044618658097711785492504343953926634992332820282019728792003956564819968")
	c := []byte{1, 2, 3, 4}
	z := BitTest(a, 1)

	if z != true {
		t.Error("Expected true got", z)
	}

	U256(a)
	S256(a)

	U256(b)
	S256(b)

	BigD(c)
}

func TestBigMax(t *testing.T) {
	a := Big("10")
	b := Big("5")

	max1 := BigMax(a, b)
	if max1 != a {
		t.Errorf("Expected %d got %d", a, max1)
	}

	max2 := BigMax(b, a)
	if max2 != a {
		t.Errorf("Expected %d got %d", a, max2)
	}
}

func TestBigMin(t *testing.T) {
	a := Big("10")
	b := Big("5")

	min1 := BigMin(a, b)
	if min1 != b {
		t.Errorf("Expected %d got %d", b, min1)
	}

	min2 := BigMin(b, a)
	if min2 != b {
		t.Errorf("Expected %d got %d", b, min2)
	}
}
