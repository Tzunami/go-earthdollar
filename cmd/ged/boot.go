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
	discover.MustParseNode("enode://bcdf724a9b59d00fb5392b49d62a04a1b46a3e3fb66ff376dd4f3d2cfd0c255afdff3f83b934a57016df7d121fdb83628431aefbb437ac14580290ce1121b960@52.60.210.252:20203"),//Canada
	discover.MustParseNode("enode://4a1a7248a3ba1d0fd14f033e0a27514514899ef8e1b6315bf951f123e428a43a0a69dbbb40d5888f865a81401ebb2a967c7cc57d8345f03410b5b4b5d2095a2c@35.154.34.226:20203"),// India
        discover.MustParseNode("enode://fd205125bae923b64021f276ab0a7f5f790c4a40c3837cc4d94b25863be83985cca0f1b45783e84194309136304efc51bc51abb5e7bf8e277fc78091b91de7c8@52.57.244.196:20203"),//Germany
	//discover.MustParseNode("enode://5fbfb426fbb46f8b8c1bd3dd140f5b511da558cd37d60844b525909ab82e13a25ee722293c829e52cb65c2305b1637fa9a2ea4d6634a224d5f400bfe244ac0de@162.243.55.45:30303"),   //pys-
	//discover.MustParseNode("enode://42d8f29d1db5f4b2947cd5c3d76c6d0d3697e6b9b3430c3d41e46b4bb77655433aeedc25d4b4ea9d8214b6a43008ba67199374a9b53633301bca0cd20c6928ab@104.155.176.151:30303"), //boot.gastracker.io
	//discover.MustParseNode("enode://814920f1ec9510aa9ea1c8f79d8b6e6a462045f09caa2ae4055b0f34f7416fca6facd3dd45f1cf1673c0209e0503f02776b8ff94020e98b6679a0dc561b4eba0@104.154.136.117:30303"), //boot1.etcdevteam.com
	//discover.MustParseNode("enode://72e445f4e89c0f476d404bc40478b0df83a5b500d2d2e850e08eb1af0cd464ab86db6160d0fde64bd77d5f0d33507ae19035671b3c74fec126d6e28787669740@104.198.71.200:30303"),  //boot2.etcdevteam.com
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	//discover.MustParseNode("enode://fb28713820e718066a2f5df6250ae9d07cff22f672dbf26be6c75d088f821a9ad230138ba492c533a80407d054b1436ef18e951bb65e6901553516c8dffe8ff0@104.155.176.151:30304"), //boot.gastracker.io
	//discover.MustParseNode("enode://afdc6076b9bf3e7d3d01442d6841071e84c76c73a7016cb4f35c0437df219db38565766234448f1592a07ba5295a867f0ce87b359bf50311ed0b830a2361392d@104.154.136.117:30403"), //boot1.etcdevteam.com
	//discover.MustParseNode("enode://21101a9597b79e933e17bc94ef3506fe99a137808907aa8fefa67eea4b789792ad11fb391f38b00087f8800a2d3dff011572b62a31232133dd1591ac2d1502c8@104.198.71.200:30403"),  //boot2.etcdevteam.com
	//discover.MustParseNode("enode://fd008499e9c4662f384b3cff23438879d31ced24e2d19504c6389bc6da6c882f9c2f8dbed972f7058d7650337f54e4ba17bb49c7d11882dd1731d26a6e62e3cb@35.187.57.94:30304"),    //boot3.etcdevteam.com
}
