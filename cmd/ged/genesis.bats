#!/usr/bin/env bats

: ${GED_CMD:=$GOPATH/bin/ged}

setup() {
	DATA_DIR=`mktemp -d`
	# Default constants.
	GENESIS_MAINNET=0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3
	GENESIS_TESTNET=0x0cd786a2425d16f152c658316c423e6ce1181e15c3295826d7c9904cba9ce303
}

teardown() {
	rm -fr $DATA_DIR
	unset GENESIS_MAINNET
	unset GENESIS_TESTNET
}

<<<<<<< HEAD:cmd/ged/genesis.bats
@test "genesis" {
	echo '{
	"alloc"      : {},
	"coinbase"   : "0x0000000000000000000000000000000000000000",
	"difficulty" : "0x020000",
	"extraData"  : "",
	"gasLimit"   : "0x2fefd8",
	"nonce"      : "0x0000000000000042",
	"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp"  : "0x00"
}' > $DATA_DIR/genesis.json

	run $GED_CMD --datadir $DATA_DIR init $DATA_DIR/genesis.json
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *"successfully wrote genesis block and/or chain rule set"* ]]

	run $GED_CMD --datadir $DATA_DIR --maxpeers 0 --nodiscover --nat none --ipcdisable --exec 'ed.getBlock(0).nonce' console
	echo "$output"
	[[ "$output" == *'"0x0000000000000042"'* ]]
}

@test "genesis empty chain config" {
	echo '{
	"alloc"      : {},
	"coinbase"   : "0x0000000000000000000000000000000000000000",
	"difficulty" : "0x020000",
	"extraData"  : "",
	"gasLimit"   : "0x2fefd8",
	"nonce"      : "0x0000000000000042",
	"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp"  : "0x00",
	"config"     : {}
}' > $DATA_DIR/genesis.json

	run $GED_CMD --datadir $DATA_DIR init $DATA_DIR/genesis.json
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *"successfully wrote genesis block and/or chain rule set"* ]]

	run $GED_CMD --datadir $DATA_DIR --maxpeers 0 --nodiscover --nat none --ipcdisable --exec 'ed.getBlock(0).nonce' console
	echo "$output"
	[[ "$output" == *'"0x0000000000000042"'* ]]
}

@test "genesis chain config" {
	echo '{
	"alloc"      : {},
	"coinbase"   : "0x0000000000000000000000000000000000000000",
	"difficulty" : "0x020000",
	"extraData"  : "",
	"gasLimit"   : "0x2fefd8",
	"nonce"      : "0x0000000000000042",
	"mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
	"timestamp"  : "0x00",
	"config"     : {}
}' > $DATA_DIR/genesis.json

	run $GED_CMD --datadir $DATA_DIR init $DATA_DIR/genesis.json
=======
# Mainnet.
@test "defaults: genesis block hash mainnet constant @ _" {
	run $GETH_CMD --data-dir $DATA_DIR --exec 'eth.getBlock(0).hash' console
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *'"0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3"'* ]]
}

# Testnet.
@test "defaults: genesis block hash constant @ --chain=morden" {
	run $GETH_CMD --chain=morden --data-dir $DATA_DIR --exec 'eth.getBlock(0).hash' console
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *'"0x0cd786a2425d16f152c658316c423e6ce1181e15c3295826d7c9904cba9ce303"'* ]]
}

@test "defaults: genesis block hash constant @ --chain=testnet" {
	run $GETH_CMD --chain=testnet --data-dir $DATA_DIR --exec 'eth.getBlock(0).hash' console
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/genesis.bats
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *'"0x0cd786a2425d16f152c658316c423e6ce1181e15c3295826d7c9904cba9ce303"'* ]]
}

<<<<<<< HEAD:cmd/ged/genesis.bats
	run $GED_CMD --datadir $DATA_DIR --maxpeers 0 --nodiscover --nat none --ipcdisable --exec 'ed.getBlock(0).nonce' console
=======
@test "defaults: genesis block hash constant @ --testnet" {
	run $GETH_CMD --testnet --data-dir $DATA_DIR --exec 'eth.getBlock(0).hash' console
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3:cmd/geth/genesis.bats
	echo "$output"

	[ "$status" -eq 0 ]
	[[ "$output" == *'"0x0cd786a2425d16f152c658316c423e6ce1181e15c3295826d7c9904cba9ce303"'* ]]
}
