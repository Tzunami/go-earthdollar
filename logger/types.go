// Copyright 2015 The go-earthdollar Authors
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

package logger

import (
	"math/big"
	"time"
)

type utctime8601 struct{}

func (utctime8601) MarshalJSON() ([]byte, error) {
	timestr := time.Now().UTC().Format(time.RFC3339Nano)
	// Bounds check
	if len(timestr) > 26 {
		timestr = timestr[:26]
	}
	return []byte(`"` + timestr + `Z"`), nil
}

type JsonLog interface {
	EventName() string
}

type LogEvent struct {
	// Guid string      `json:"guid"`
	Ts utctime8601 `json:"ts"`
	// Level string      `json:"level"`
}

type LogStarting struct {
	ClientString    string `json:"client_impl"`
	ProtocolVersion int    `json:"ed_version"`
	LogEvent
}

func (l *LogStarting) EventName() string {
	return "starting"
}

type P2PConnected struct {
	RemoteId            string `json:"remote_id"`
	RemoteAddress       string `json:"remote_addr"`
	RemoteVersionString string `json:"remote_version_string"`
	NumConnections      int    `json:"num_connections"`
	LogEvent
}

func (l *P2PConnected) EventName() string {
	return "p2p.connected"
}

type P2PDisconnected struct {
	NumConnections int    `json:"num_connections"`
	RemoteId       string `json:"remote_id"`
	LogEvent
}

func (l *P2PDisconnected) EventName() string {
	return "p2p.disconnected"
}

type EdMinerNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *EdMinerNewBlock) EventName() string {
	return "ed.miner.new_block"
}

type EdChainReceivedNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	RemoteId      string   `json:"remote_id"`
	LogEvent
}

func (l *EdChainReceivedNewBlock) EventName() string {
	return "ed.chain.received.new_block"
}

type EdChainNewHead struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *EdChainNewHead) EventName() string {
	return "ed.chain.new_head"
}

type EdTxReceived struct {
	TxHash   string `json:"tx_hash"`
	RemoteId string `json:"remote_id"`
	LogEvent
}

func (l *EdTxReceived) EventName() string {
	return "ed.tx.received"
}

//
//
// The types below are legacy and need to be converted to new format or deleted
//
//

// type P2PConnecting struct {
// 	RemoteId       string `json:"remote_id"`
// 	RemoteEndpoint string `json:"remote_endpoint"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PConnecting) EventName() string {
// 	return "p2p.connecting"
// }

// type P2PHandshaked struct {
// 	RemoteCapabilities []string `json:"remote_capabilities"`
// 	RemoteId           string   `json:"remote_id"`
// 	NumConnections     int      `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PHandshaked) EventName() string {
// 	return "p2p.handshaked"
// }

// type P2PDisconnecting struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnecting) EventName() string {
// 	return "p2p.disconnecting"
// }

// type P2PDisconnectingBadHandshake struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingBadHandshake) EventName() string {
// 	return "p2p.disconnecting.bad_handshake"
// }

// type P2PDisconnectingBadProtocol struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingBadProtocol) EventName() string {
// 	return "p2p.disconnecting.bad_protocol"
// }

// type P2PDisconnectingReputation struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingReputation) EventName() string {
// 	return "p2p.disconnecting.reputation"
// }

// type P2PDisconnectingDHT struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingDHT) EventName() string {
// 	return "p2p.disconnecting.dht"
// }

// type P2PEdDisconnectingBadBlock struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PEdDisconnectingBadBlock) EventName() string {
// 	return "p2p.ed.disconnecting.bad_block"
// }

// type P2PEdDisconnectingBadTx struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PEdDisconnectingBadTx) EventName() string {
// 	return "p2p.ed.disconnecting.bad_tx"
// }

// type EdNewBlockBroadcasted struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockBroadcasted) EventName() string {
// 	return "ed.newblock.broadcasted"
// }

// type EdNewBlockIsKnown struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockIsKnown) EventName() string {
// 	return "ed.newblock.is_known"
// }

// type EdNewBlockIsNew struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockIsNew) EventName() string {
// 	return "ed.newblock.is_new"
// }

// type EdNewBlockMissingParent struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockMissingParent) EventName() string {
// 	return "ed.newblock.missing_parent"
// }

// type EdNewBlockIsInvalid struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockIsInvalid) EventName() string {
// 	return "ed.newblock.is_invalid"
// }

// type EdNewBlockChainIsOlder struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockChainIsOlder) EventName() string {
// 	return "ed.newblock.chain.is_older"
// }

// type EdNewBlockChainIsCanonical struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockChainIsCanonical) EventName() string {
// 	return "ed.newblock.chain.is_cannonical"
// }

// type EdNewBlockChainNotCanonical struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EdNewBlockChainNotCanonical) EventName() string {
// 	return "ed.newblock.chain.not_cannonical"
// }

// type EdTxCreated struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxHexRLP  string `json:"tx_hexrlp"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EdTxCreated) EventName() string {
// 	return "ed.tx.created"
// }

// type EdTxBroadcasted struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EdTxBroadcasted) EventName() string {
// 	return "ed.tx.broadcasted"
// }

// type EdTxValidated struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EdTxValidated) EventName() string {
// 	return "ed.tx.validated"
// }

// type EdTxIsInvalid struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	Reason    string `json:"reason"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EdTxIsInvalid) EventName() string {
// 	return "ed.tx.is_invalid"
// }
