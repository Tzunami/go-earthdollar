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

package utils

import "github.com/Tzunami/go-earthdollar/p2p/discover"

// HomesteadBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Homestead network.
var HomesteadBootNodes = []*discover.Node{
	// ETC Volunteer Go Bootnodes
	// EPROJECT Upgrade by supplying a default list, parsing oracle & random selection
	discover.MustParseNode("enode://ec8ae764f7cb0417bdfb009b9d0f18ab3818a3a4e8e7c67dd5f18971a93510a2e6f43cd0b69a27e439a9629457ea804104f37c85e41eed057d3faabbf7744cdf@13.74.157.139:30429"),
	discover.MustParseNode("enode://c2e1fceb3bf3be19dff71eec6cccf19f2dbf7567ee017d130240c670be8594bc9163353ca55dd8df7a4f161dd94b36d0615c17418b5a3cdcbb4e9d99dfa4de37@13.74.157.139:30430"),
    discover.MustParseNode("enode://fe29b82319b734ce1ec68b84657d57145fee237387e63273989d354486731e59f78858e452ef800a020559da22dcca759536e6aa5517c53930d29ce0b1029286@13.74.157.139:30431"),
	discover.MustParseNode("enode://1d7187e7bde45cf0bee489ce9852dd6d1a0d9aa67a33a6b8e6db8a4fbc6fcfa6f0f1a5419343671521b863b187d1c73bad3603bae66421d157ffef357669ddb8@13.74.157.139:30432"),
	discover.MustParseNode("enode://0e4cba800f7b1ee73673afa6a4acead4018f0149d2e3216be3f133318fd165b324cd71b81fbe1e80deac8dbf56e57a49db7be67f8b9bc81bd2b7ee496434fb5d@13.74.157.139:30433"),
	
	//boot.gastracker.io

	// Pending & Not Resolving
	//discover.MustParseNode("enode://b61123cc535d6bac44f9e6ff8637a30a10198f80b5582148dcd84ef8039a4b90e326bb7f6964588a46bcf1ccd8e8bba65db514fc72e3026ff13b20959f45f654@etc.naphex.rocks:20203"), // nap-
	//discover.MustParseNode("enode://d55f15f28317c21c359c8f62b93b7059aa2fcd586c0b0d431f97c4b8f27ee8f58fbe060b72eff95790b7ecd34c2a9b02458a783e61d8ec2aa37cdad6b0fc6d9a@node1.ethc.io:20203"),  //phr-
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETC Nodes
	discover.MustParseNode("enode://1d7187e7bde45cf0bee489ce9852dd6d1a0d9aa67a33a6b8e6db8a4fbc6fcfa6f0f1a5419343671521b863b187d1c73bad3603bae66421d157ffef357669ddb8@13.74.157.139:30432"), //boot.gastracker.io

	// ETH/DEV Go Bootnodes
	discover.MustParseNode("enode://0e4cba800f7b1ee73673afa6a4acead4018f0149d2e3216be3f133318fd165b324cd71b81fbe1e80deac8dbf56e57a49db7be67f8b9bc81bd2b7ee496434fb5d@13.74.157.139:30433"),

	// ETH/DEV Cpp Bootnodes
}
