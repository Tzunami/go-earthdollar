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

import "github.com/Tzunami/go-earthdollar/p2p/discover"

// HomesteadBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Homestead network.
var HomesteadBootNodes = []*discover.Node{
	// EPROJECT Upgrade by supplying a default list, parsing oracle & random selection

        discover.MustParseNode("enode://839ab987e9a2d093761f796be3a493f93da89fe89c99851a045b026a6524ffbcacf22eeb046d38aea82a2e98adfaebab15d9667be17fddad864f02b0cf15a71e@35.182.142.37:20203?discport=20203"), //Canada
	discover.MustParseNode("enode://bcdf724a9b59d00fb5392b49d62a04a1b46a3e3fb66ff376dd4f3d2cfd0c255afdff3f83b934a57016df7d121fdb83628431aefbb437ac14580290ce1121b960@52.60.68.21:20203?discport=20203"),//Canada
        discover.MustParseNode("enode://5486cf80330f26659e63b708f87f318b7663659980ca927f33cba4e5cacd6825b2e9de57b1df14a14964d67092620803b7ac115b74c5b9a64998d9b5162fb9af@35.154.120.185:20203?discport=20203"),// India
        discover.MustParseNode("enode://4a1a7248a3ba1d0fd14f033e0a27514514899ef8e1b6315bf951f123e428a43a0a69dbbb40d5888f865a81401ebb2a967c7cc57d8345f03410b5b4b5d2095a2c@13.126.74.243:20203?discport=20203"),// India
        discover.MustParseNode("enode://b689025552ef44b8c77c5aa3270371b85027a8d985aae6e743a186fd7cec106139b8ae933baefb83d70100e09daecd208602b89b02636a175b9c86afc25d6e49@52.58.22.43:20203?discport=20203"),//Germany
        discover.MustParseNode("enode://fd205125bae923b64021f276ab0a7f5f790c4a40c3837cc4d94b25863be83985cca0f1b45783e84194309136304efc51bc51abb5e7bf8e277fc78091b91de7c8@35.156.76.0:20203?discport=20203"),//Germany
        
	//discover.MustParseNode("enode://353b84ba90365a0b56249429a4eb7b31101a6ea8efbaeb395b9eaf8f82bcfa5673ab6dbdb72612393ccff85da54e2f6694ff0de22699c3b7ec7a0586ca215c04@35.182.15.33 :20203"), //mike 1
	//discover.MustParseNode("enode://bcdf724a9b59d00fb5392b49d62a04a1b46a3e3fb66ff376dd4f3d2cfd0c255afdff3f83b934a57016df7d121fdb83628431aefbb437ac14580290ce1121b960@35.182.2.114:20203"), //mike 2	
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	//discover.MustParseNode("enode://fb28713820e718066a2f5df6250ae9d07cff22f672dbf26be6c75d088f821a9ad230138ba492c533a80407d054b1436ef18e951bb65e6901553516c8dffe8ff0@104.155.176.151:30304"), //boot.gastracker.io
	//discover.MustParseNode("enode://afdc6076b9bf3e7d3d01442d6841071e84c76c73a7016cb4f35c0437df219db38565766234448f1592a07ba5295a867f0ce87b359bf50311ed0b830a2361392d@104.154.136.117:30403"), //boot1.etcdevteam.com
	//discover.MustParseNode("enode://21101a9597b79e933e17bc94ef3506fe99a137808907aa8fefa67eea4b789792ad11fb391f38b00087f8800a2d3dff011572b62a31232133dd1591ac2d1502c8@104.198.71.200:30403"),  //boot2.etcdevteam.com
	//discover.MustParseNode("enode://fd008499e9c4662f384b3cff23438879d31ced24e2d19504c6389bc6da6c882f9c2f8dbed972f7058d7650337f54e4ba17bb49c7d11882dd1731d26a6e62e3cb@35.187.57.94:30304"),    //boot3.etcdevteam.com
}
