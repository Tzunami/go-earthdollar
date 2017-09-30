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

    // Canada
    discover.MustParseNode("enode://25d809c3adfe4c5d6a9678f90dd337df6fd101a7b0cc01c6596a911ecb7917cf475ee1fd82db97f2ccdcb32caf9cfc24070b13828a3149690dccfd10b7ca75f1@35.182.142.37:20203"), 
    discover.MustParseNode("enode://d3a074b27c88c6e765d39867f6b341f7c577fe264c46071948c85c96b328ed76148630087da155df4ee75439bb9b3b23b8f28c55750cbbe066f553796a6eddd8@52.60.68.21:20203"), 

    // India
    discover.MustParseNode("enode://92f2235e1b3a403c8901a0520929b68f7ba91400b8ecc36952a3b7a2a82484eafaf16366a5f75e72a0ed14a05cc8b3ddbf794b49a75c7e731eaff0928145c30b@35.154.120.185:20203"), 
    discover.MustParseNode("enode://291251fa0a7c93e24f36edf04bf1d6c2b450b3fd0b1ea0d4e108fee446070e86fbdaa09c5f2700eb26072d149bb3b4ee51ed6f5d31d3550a43350857e9da7744@13.126.74.243:20203"), 

    // Germany
    discover.MustParseNode("enode://03478f6c9efb22ec75752169c8701967f4ba029b410cd1ed95ebd3916080f02f045034a566af767db95252b333231f2e20cfad6231515297e39ff40d2288a398@52.58.22.43:20203"), 
    discover.MustParseNode("enode://905c384d045d1ce6014c1bf0f8a99e8450cc42ece7ed8949e52b6b4e533093dab50710c3615aca8042416062b413c2cf317fde422acaada44e51075bf78eca93@35.156.76.0:20203"), 

}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	//discover.MustParseNode("enode://fb28713820e718066a2f5df6250ae9d07cff22f672dbf26be6c75d088f821a9ad230138ba492c533a80407d054b1436ef18e951bb65e6901553516c8dffe8ff0@104.155.176.151:30304"), //boot.gastracker.io
	//discover.MustParseNode("enode://afdc6076b9bf3e7d3d01442d6841071e84c76c73a7016cb4f35c0437df219db38565766234448f1592a07ba5295a867f0ce87b359bf50311ed0b830a2361392d@104.154.136.117:30403"), //boot1.etcdevteam.com
	//discover.MustParseNode("enode://21101a9597b79e933e17bc94ef3506fe99a137808907aa8fefa67eea4b789792ad11fb391f38b00087f8800a2d3dff011572b62a31232133dd1591ac2d1502c8@104.198.71.200:30403"),  //boot2.etcdevteam.com
	//discover.MustParseNode("enode://fd008499e9c4662f384b3cff23438879d31ced24e2d19504c6389bc6da6c882f9c2f8dbed972f7058d7650337f54e4ba17bb49c7d11882dd1731d26a6e62e3cb@35.187.57.94:30304"),    //boot3.etcdevteam.com
}
