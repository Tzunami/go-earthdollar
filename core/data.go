// Copyright 2017 The go-earthdollar Authors
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

package core

import (
	//"math/big"

	//"github.com/Tzunami/go-earthdollar/common"
	//"github.com/Tzunami/go-earthdollar/core/vm"
)

// This file contains configuration literals.

// DefaultConfig is the Earthdollar standard setup.
var DefaultConfig = &ChainConfig{
	Forks: []*Fork{/* //earthdollar
		{
			Name:         "Homestead",
			Block:        big.NewInt(1150000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(20),
				ExtcodeCopy:     big.NewInt(20),
				Balance:         big.NewInt(20),
				SLoad:           big.NewInt(50),
				Calls:           big.NewInt(40),
				Suicide:         big.NewInt(0),
				ExpByte:         big.NewInt(10),
				CreateBySuicide: nil,
			},
		}, {
			Name:         "ETF",
			Block:        big.NewInt(1920000),
			NetworkSplit: true,
			Support:      false,
			RequiredHash: common.HexToHash("94365e3a8c0b35089c1d1195081fe7489b528a84b22199c916180db8b28ade7f"),
		}, {
			Name:         "GasReprice",
			Block:        big.NewInt(2500000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(700),
				ExtcodeCopy:     big.NewInt(700),
				Balance:         big.NewInt(400),
				SLoad:           big.NewInt(200),
				Calls:           big.NewInt(700),
				Suicide:         big.NewInt(5000),
				ExpByte:         big.NewInt(10),
				CreateBySuicide: big.NewInt(25000),
			},
		}, {
			Name:         "Diehard",
			Block:        big.NewInt(3000000),
			Length:       big.NewInt(2000000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(700),
				ExtcodeCopy:     big.NewInt(700),
				Balance:         big.NewInt(400),
				SLoad:           big.NewInt(200),
				Calls:           big.NewInt(700),
				Suicide:         big.NewInt(5000),
				ExpByte:         big.NewInt(50),
				CreateBySuicide: big.NewInt(25000),
			},
		},*/
	},
	BadHashes: []*BadHash{
		{
			// consensus issue that occurred on the Frontier network at block 116,522, mined on 2015-08-20 at 14:59:16+02:00
			// https://blog.earthdollar.org/2015/08/20/security-alert-consensus-issue
                        // earthdollar
			//Block: big.NewInt(116522),
			//Hash:  common.HexToHash("05bef30ef572270f654746da22639a7a0c97dd97a7050b9e252391996aaeb689"),
		},
	},
	//ChainId: big.NewInt(61),
}

// TestConfig is the semi-official setup for testing purposes.
var TestConfig = &ChainConfig{
	Forks: []*Fork{
		/*{ //earthdollar
			Name:         "Homestead",
			Block:        big.NewInt(494000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(20),
				ExtcodeCopy:     big.NewInt(20),
				Balance:         big.NewInt(20),
				SLoad:           big.NewInt(50),
				Calls:           big.NewInt(40),
				Suicide:         big.NewInt(0),
				ExpByte:         big.NewInt(10),
				CreateBySuicide: nil,
			},
		},
		{
			Name:         "GasReprice",
			Block:        big.NewInt(1783000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(700),
				ExtcodeCopy:     big.NewInt(700),
				Balance:         big.NewInt(400),
				SLoad:           big.NewInt(200),
				Calls:           big.NewInt(700),
				Suicide:         big.NewInt(5000),
				ExpByte:         big.NewInt(10),
				CreateBySuicide: big.NewInt(25000),
			},
		},
		{
			Name:         "ETF",
			Block:        big.NewInt(1885000),
			NetworkSplit: true,
			Support:      false,
			RequiredHash: common.HexToHash("2206f94b53bd0a4d2b828b6b1a63e576de7abc1c106aafbfc91d9a60f13cb740"),
		},
		{
			Name:         "Diehard",
			Block:        big.NewInt(1915000),
			Length:       big.NewInt(1500000),
			NetworkSplit: false,
			Support:      true,
			GasTable: &vm.GasTable{
				ExtcodeSize:     big.NewInt(700),
				ExtcodeCopy:     big.NewInt(700),
				Balance:         big.NewInt(400),
				SLoad:           big.NewInt(200),
				Calls:           big.NewInt(700),
				Suicide:         big.NewInt(5000),
				ExpByte:         big.NewInt(50),
				CreateBySuicide: big.NewInt(25000),
			},
		},*/
	},
	BadHashes: []*BadHash{
		{
			// consensus issue at Testnet #383792
			// http://earthdollar.stackexchange.com/questions/10183/upgraded-to-ged-1-5-0-bad-block-383792
                        // earthdollar
			//Block: big.NewInt(383792),
			//Hash:  common.HexToHash("9690db54968a760704d99b8118bf79d565711669cefad24b51b5b1013d827808"),
		},
		{
			// chain followed by non-diehard testnet
                        // earthdollar
			//Block: big.NewInt(1915277),
			//Hash:  common.HexToHash("3bef9997340acebc85b84948d849ceeff74384ddf512a20676d424e972a3c3c4"),
		},
	}, 
	//ChainId: big.NewInt(62),
}

// TestNetGenesis representing the Morden test net genesis block.
var TestNetGenesis = &GenesisDump{
	Nonce:      "0x00006d6f7264656e",
	Difficulty: "0x020000",
	Mixhash:    "0x00000000000000000000000000000000000000647572616c65787365646c6578",
	GasLimit:   "0x2FEFD8",
	Alloc: map[hex]*GenesisDumpAlloc{
		"0000000000000000000000000000000000000001": {Balance: "1"},
		"0000000000000000000000000000000000000002": {Balance: "1"},
		"0000000000000000000000000000000000000003": {Balance: "1"},
		"0000000000000000000000000000000000000004": {Balance: "1"},
		"102e61f5d8f9bc71d0ad4a084df4e65e05ce0e1c": {Balance: "1606938044258990275541962092341162602522202993782792835301376"},
	},
}

// OlympicGenesis representing the Olympic genesis block.
var OlympicGenesis = &GenesisDump{}

// DefaultGenesis representing the default Earthdollar genesis block.
var DefaultGenesis = &GenesisDump{
	Difficulty: "0x020000",
	GasLimit:   "0x47E7C4",
	Nonce:      "0x000000000000002a",
	Alloc: map[hex]*GenesisDumpAlloc{
                "e856f883f4862cb7f55a35a5b554451798902d16":  {Balance: "100000000000000000000000000"},  
                "4e32fb7cb1d33861aa2677d7ff32da16027e7e08":  {Balance: "100000000000000000000000000"},
                "2ba175ee5b11ac09eabbef73234452b5857a0f01":  {Balance: "100000000000000000000000000"},
                "681c1dcdfaaf43b37bb5db81d219e801c5d6426f":  {Balance: "100000000000000000000000000"}, 
                "5b1c61d10fe21e45182c71987abda0eab33ea9e7":  {Balance: "100000000000000000000000000"}, 
                "84bb68e581f8513945d7c2269e134f61abdceb77":  {Balance: "100000000000000000000000000"}, 
                "1ed132a81aaea349d619c71a580d1426fc8cf6dc":  {Balance: "100000000000000000000000000"}, 
                "aa7a66a45e61f2e31980150dc2e79898cf2b9b6b":  {Balance: "100000000000000000000000000"}, 
                "150a588f68344a61800b3c3761a37e57231bf454":  {Balance: "100000000000000000000000000"}, 
                "b7fa96bb09aaa87c642c7fb753d2ef0b410ffd29":  {Balance: "100000000000000000000000000"}, 
                "062305dbbeff97f2cd7d16a2e76780c64b0794e9":  {Balance: "100000000000000000000000000"}, 
                "d3842991acd4823fa0f22f7915aba179ca1c84ff":  {Balance: "100000000000000000000000000"}, 
                "80ef182cfd269467c8d8732aae65c046da5ccee7":  {Balance: "100000000000000000000000000"}, 
                "e91efd17378a653d3d36b336bfdeefd858bf0eb4":  {Balance: "100000000000000000000000000"}, 
                "61e342a5430c9fd2d9427a5794ff85bfea20af77":  {Balance: "100000000000000000000000000"}, 
                "bae738480167bd65284a6f85d8bc661f22b2364e":  {Balance: "100000000000000000000000000"}, 
                "ba9fc55c1a79b4ec3a2c78c6e82996c74d6dc6ba":  {Balance: "100000000000000000000000000"}, 
                "5768a44376352a25155452337ddeb647b7988ac0":  {Balance: "100000000000000000000000000"}, 
                "61f2a927f5f7d91786f8779cd0ea4d769201f1ce":  {Balance: "100000000000000000000000000"}, 
                "eef42335bc391518bf07a03518918c7ab0de9e9c":  {Balance: "100000000000000000000000000"}, 
	},
}
