[![Build Status](https://travis-ci.org/Tzunami/go-earthdollar.svg?branch=master)](https://travis-ci.org/Tzunami/go-earthdollar)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/github/Tzunami/go-earthdollar?svg=true)](https://ci.appveyor.com/project/splix/go-earthdollar)
[![API Reference](https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/Tzunami/go-earthdollar)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Tzunami/go-earthdollar?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Download](https://api.bintray.com/packages/Tzunami/GoEarthdollarClassic/go-earthdollar/images/download.svg)](https://bintray.com/Tzunami/GoEarthdollarClassic/go-earthdollar/_latestVersion)

## Earthdollar Go (Earthdollar Blockchain)

<<<<<<< HEAD
Official golang implementation of the Earthdollar protocol supporting the
original chain. A version which can **honestly** offer both a censorship
resistant and unstoppable application platform for developers.

This is a project migrated from the now hard forked Earthdollar (EDF) github project, we
will need to slowly migrate pieces of the infrastructure required to
maintain the project. We will apply all upstream patches unrelated to the DAO HF while organizing
development.
=======
Official Go language implementation of the Earthdollar protocol supporting the
_original_ chain. Earthdollar (ETC) offers a censorship-resistant and powerful application platform for developers in parallel to Earthdollar (ETHF), while differentially rejecting the DAO bailout.

## Install
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

### :rocket: From a release binary
The simplest way to get started running a node is to visit our [Releases page](https://github.com/Tzunami/go-earthdollar/releases) and download a zipped executable binary (matching your operating system, of course), then moving the unzipped file `ged` to somewhere in your `$PATH`. Now you should be able to open a terminal and run `$ ged help` to make sure it's working. For additional installation instructions please check out the [Installation Wiki](https://github.com/Tzunami/go-earthdollar/wiki/Home#Developers).

CLI one-liner for Darwin:
```bash
$ curl -L -o ~/Downloads/ged-3.5.zip https://github.com/Tzunami/go-earthdollar/releases/download/v3.5.0/ged-osx-v3.5.0.zip; unzip ~/Downloads/ged-3.5.zip -d $HOME/bin/

<<<<<<< HEAD
Building ged requires both a Go and a C compiler.

To install the full suite of utilities run `go install github.com/Tzunami/go-earthdollar/cmd/...`.

## Executables

The go-earthdollar project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`ged`** | Our main Earthdollar CLI client. It is the entry point into the Earthdollar network (main-, test- or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as an gateway into the Earthdollar network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. Please see our [Command Line Options](https://github.com/Tzunami/go-earthdollar/wiki/Command-Line-Options) wiki page for details. |
| `abigen` | Source code generator to convert Earthdollar contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Earthdollar contract ABIs](https://github.com/Tzunami/wiki/wiki/Earthdollar-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/Tzunami/go-earthdollar/wiki/Native-DApps:-Go-bindings-to-Earthdollar-contracts) wiki page for details. |
| `bootnode` | Stripped down version of our Earthdollar client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |
| `disasm` | Bytecode disassembler to convert EVM (Earthdollar Virtual Machine) bytecode into more user friendly assembly-like opcodes (e.g. `echo "6001" | disasm`). For details on the individual opcodes, please see pages 22-30 of the [Earthdollar Yellow Paper](http://gavwood.com/paper.pdf). |
| `evm` | Developer utility version of the EVM (Earthdollar Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow insolated, fine graned debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug`). |
| `gedrpctest` | Developer utility tool to support our [earthdollar/rpc-test](https://github.com/Tzunami/rpc-tests) test suite which validates baseline conformity to the [Earthdollar JSON RPC](https://github.com/Tzunami/wiki/wiki/JSON-RPC) specs. Please see the [test suite's readme](https://github.com/Tzunami/rpc-tests/blob/master/README.md) for details. |
| `rlpdump` | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com/Tzunami/wiki/wiki/RLP)) dumps (data encoding used by the Earthdollar protocol both network as well as consensus wise) to user friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`). |

## Running ged

Going through all the possible command line flags is out of scope here (please consult our
[CLI Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/Command-Line-Options)), but we've
enumerated a few common parameter combos to get you up to speed quickly on how you can run your
own Ged instance.
=======
$ ged help
```

### :hammer: Building the source

If your heart is set on the bleeding edge, install from source. However, please be advised that you may encounter some strange things, and we can't prioritize support beyond the release versions. Recommended for developers only.

#### Dependencies
Building ged requires both Go >=1.8 and a C compiler.

#### Get source and dependencies
`$ go get -v github.com/Tzunami/go-earthdollar/...`

#### Installing command executables

To install...
- the full suite of utilities: `$ go install github.com/Tzunami/go-earthdollar/cmd/...`
- just __ged__: `$ go install github.com/Tzunami/go-earthdollar/cmd/ged`

Executables built from source will, by default, be installed in `$GOPATH/bin/`.

## Executables

This repository includes several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`ged`** | The main Earthdollar CLI client. It is the entry point into the Earthdollar network (main-, test-, or private net), capable of running as a full node (default) archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Earthdollar network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. Please see our [Command Line Options](https://github.com/Tzunami/go-earthdollar/wiki/Command-Line-Options) wiki page for details. |
| `abigen` | Source code generator to convert Earthdollar contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Earthdollar contract ABIs](https://github.com.earthdollarproject/wiki/wiki/Earthdollar-Contract-ABI) with expanded functionality if the contract bytecode is also available. However it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://github.com/Tzunami/go-earthdollar/wiki/Native-DApps-in-Go) wiki page for details. |
| `bootnode` | Stripped down version of our Earthdollar client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |
| `disasm` | Bytecode disassembler to convert EVM (Earthdollar Virtual Machine) bytecode into more user friendly assembly-like opcodes (e.g. `echo "6001" | disasm`). For details on the individual opcodes, please see pages 22-30 of the [Earthdollar Yellow Paper](http://gavwood.com/paper.pdf). |
| `evm` | Developer utility version of the EVM (Earthdollar Virtual Machine) that is capable of running bytecode snippets within a configurable environment and execution mode. Its purpose is to allow insolated, fine graned debugging of EVM opcodes (e.g. `evm --code 60ff60ff --debug`). |
| `gedrpctest` | Developer utility tool to support our .earthdollar/rpc-test](https://github.com.earthdollarproject/rpc-tests) test suite which validates baseline conformity to the [Earthdollar JSON RPC](https://github.com.earthdollarproject/wiki/wiki/JSON-RPC) specs. Please see the [test suite's readme](https://github.com.earthdollarproject/rpc-tests/blob/master/README.md) for details. |
| `rlpdump` | Developer utility tool to convert binary RLP ([Recursive Length Prefix](https://github.com.earthdollarproject/wiki/wiki/RLP)) dumps (data encoding used by the Earthdollar protocol both network as well as consensus wise) to user friendlier hierarchical representation (e.g. `rlpdump --hex CE0183FFFFFFC4C304050583616263`). |

## :green_book: Ged: the basics

### Data directory
By default, ged will store all node and blockchain data in a __parent directory__ depending on your OS:
- Linux: `$HOME/.earthdollar/`
- Mac: `$HOME/Library/Earthdollar/`
- Windows: `$HOME/AppData/Roaming/Earthdollar/`

__You can specify this directory__ with `--data-dir=$HOME/id/rather/put/it/here`.

Within this parent directory, ged will use a __/subdirectory__ to hold data for each network you run. The defaults are:
 - `/mainnet` for the Mainnet
 - `/morden` for the Morden Testnet

__You can specify this subdirectory__ with `--chain=mycustomnet`.

> __Migrating__: If you have existing data created prior to the [3.4 Release](https://github.com/Tzunami/go-earthdollar/releases), ged will attempt to migrate your existing standard ETC data to this structure. To learn more about managing this migration please read our [3.4 release notes on our Releases page](https://github.com/Tzunami/go-earthdollar/wiki/Release-3.4.0-Notes).
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

### Full node on the main Earthdollar network

<<<<<<< HEAD
By far the most common scenario is people wanting to simply interact with the Earthdollar network:
create accounts; transfer funds; deploy and interact with contracts. For this particular use-case
the user doesn't care about years-old historical data, so we can fast-sync quickly to the current
state of the network. To do so:
=======
```
$ ged
```

It's that easy! This will establish an ETC blockchain node and download ("sync") the full blocks for the entirety of the ETC blockchain. __However__, before you go ahead with plain ol' `ged`, we would encourage reviewing the following section...

#### :speedboat: `--fast`

The most common scenario is users wanting to simply interact with the Earthdollar network: create accounts; transfer funds; deploy and interact with contracts, and mine. For this particular use-case the user doesn't care about years-old historical data, so we can _fast-sync_ to the current state of the network. To do so:

```
$ ged --fast
```

Using ged in fast sync mode causes it to download only block _state_ data -- leaving out bulky transaction records -- which avoids a lot of CPU and memory intensive processing. 

Fast sync will be automatically __disabled__ (and full sync enabled) when:
- your chain database contains *any* full blocks
- your node has synced up to the current head of the network blockchain

In case of using `--mine` together with `--fast`, ged will operate as described; syncing in fast mode up to the head, and then begin mining once it has synced its first full block at the head of the chain.

*Note:* To further increase ged's performace, you can use a `--cache=512` flag to bump the memory allowance of the database (e.g. 512MB) which can significantly improve sync times, especially for HDD users. This flag is optional and you can set it as high or as low as you'd like, though we'd recommend the 512MB - 2GB range.

### Create or manage account(s)

Ged is able to create, import, update, unlock, and otherwise manage your private (encrypted) key files. Key files are in JSON format and, by default, stored in the respective chain folder's `/keystore` directory; you can specify a custom location with the `--keystore` flag.

```
$ ged account new
```

This command will create a new account and prompt you to enter a passphrase to protect your account.

Other `account` subcommands include:
```
SUBCOMMANDS:

        list    print account addresses
        new     create a new account
        update  update an existing account
        import  import a private key into a new account

```

Learn more at the [Accounts Wiki Page](https://github.com/Tzunami/go-earthdollar/wiki/Managing-Accounts). If you're interested in using ged to manage a lot (~100,000+) of accounts, please visit the [Indexing Accounts Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/Indexing-Accounts).
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3


### Interact with the Javascript console
```
<<<<<<< HEAD
$ ged --fast --cache=512 console
=======
$ ged console
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3
```

This command will start up Ged's built-in interactive [JavaScript console](https://github.com/Tzunami/go-earthdollar/wiki/JavaScript-Console), through which you can invoke all official [`web3` methods](https://github.com.earthdollarproject/wiki/wiki/JavaScript-API) as well as Ged's own [management APIs](https://github.com/Tzunami/go-earthdollar/wiki/Management-APIs). This too is optional and if you leave it out you can always attach to an already running Ged instance with `ged attach`.

Learn more at the [Javascript Console Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/JavaScript-Console).


### And so much more!

<<<<<<< HEAD
 * Start ged in fast sync mode (`--fast`), causing it to download more data in exchange for avoiding
   processing the entire history of the Earthdollar network, which is very CPU intensive.
 * Bump the memory allowance of the database to 512MB (`--cache=512`), which can help significantly in
   sync times especially for HDD users. This flag is optional and you can set it as high or as low as
   you'd like, though we'd recommend the 512MB - 2GB range.
 * Start up Ged's built-in interactive [JavaScript console](https://github.com/Tzunami/go-earthdollar/wiki/JavaScript-Console),
   (via the trailing `console` subcommand) through which you can invoke all official [`web3` methods](https://github.com/Tzunami/wiki/wiki/JavaScript-API)
   as well as Ged's own [management APIs](https://github.com/Tzunami/go-earthdollar/wiki/Management-APIs).
   This too is optional and if you leave it out you can always attach to an already running Ged instance
   with `ged --attach`.

### Full node on the Earthdollar test network

Transitioning towards developers, if you'd like to play around with creating Earthdollar contracts, you
almost certainly would like to do that without any real money involved until you get the hang of the
entire system. In other words, instead of attaching to the main network, you want to join the **test**
network with your node, which is fully equivalent to the main network, but with play-ED only.

```
$ ged --testnet --fast --cache=512 console
=======
For a comprehensive list of command line options, please consult our [CLI Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/Command-Line-Options).

## :orange_book: Ged: developing and advanced useage

### Morden Testnet
If you'd like to play around with creating Earthdollar contracts, you
almost certainly would like to do that without any real money involved until you get the hang of the entire system. In other words, instead of attaching to the main network, you want to join the **test** network with your node, which is fully equivalent to the main network, but with play-Ether only.

```
$ ged --chain=morden --fast --cache=512 console
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3
```

The `--fast`, `--cache` flags and `console` subcommand have the exact same meaning as above and they are equally useful on the testnet too. Please see above for their explanations if you've skipped to here.

<<<<<<< HEAD
Specifying the `--testnet` flag however will reconfigure your Ged instance a bit:

 * Instead of using the default data directory (`~/.earthdollar` on Linux for example), Ged will nest
   itself one level deeper into a `testnet` subfolder (`~/.earthdollar/testnet` on Linux).
 * Instead of connecting the main Earthdollar network, the client will connect to the test network,
   which uses different P2P bootnodes, different network IDs and genesis states.

*Note: Although there are some internal protective measures to prevent transactions from crossing
over between the main network and test network (different starting nonces), you should make sure to
always use separate accounts for play-money and real-money. Unless you manually move accounts, Ged
will by default correctly separate the two networks and will not make any accounts available between
them.*
=======
Specifying the `--chain=morden` flag will reconfigure your Ged instance a bit:

 -  As mentioned above, Ged will host its testnet data in a `morden` subfolder (`~/.earthdollar/morden`).
 - Instead of connecting the main Earthdollar network, the client will connect to the test network, which uses different P2P bootnodes, different network IDs and genesis states.

You may also optionally use `--testnet` or `--chain=testnet` to enable this configuration. 

> *Note: Although there are some internal protective measures to prevent transactions from crossing over between the main network and test network (different starting nonces), you should make sure to always use separate accounts for play-money and real-money. Unless you manually move accounts, Ged
will by default correctly separate the two networks and will not make any accounts available between them.*
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

### Programatically interfacing Ged nodes

<<<<<<< HEAD
As a developer, sooner rather than later you'll want to start interacting with Ged and the Earthdollar
network via your own programs and not manually through the console. To aid this, Ged has built in
support for a JSON-RPC based APIs ([standard APIs](https://github.com/Tzunami/wiki/wiki/JSON-RPC) and
[Ged specific APIs](https://github.com/Tzunami/go-earthdollar/wiki/Management-APIs)). These can be
exposed via HTTP, WebSockets and IPC (unix sockets on unix based platroms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Ged, whereas the HTTP
and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons.
These can be turned on/off and configured as you'd expect.
=======
As a developer, sooner rather than later you'll want to start interacting with Ged and the Earthdollar network via your own programs and not manually through the console. To aid this, Ged has built in support for a JSON-RPC based APIs ([standard APIs](https://github.com.earthdollarproject/wiki/wiki/JSON-RPC) and
[Ged specific APIs](https://github.com/Tzunami/go-earthdollar/wiki/Management-APIs)). These can be exposed via HTTP, WebSockets and IPC (unix sockets on unix based platroms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Ged, whereas the HTTP and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons. These can be turned on/off and configured as you'd expect.
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
<<<<<<< HEAD
  * `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpcport` HTTP-RPC server listening port (default: 8811)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: "ed,net,web3")
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: "localhost")
  * `--wsport` WS-RPC server listening port (default: 8812)
  * `--wsapi` API's offered over the WS-RPC interface (default: "ed,net,web3")
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,ed,miner,net,personal,shh,txpool,web3")
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a Ged node configured with the above flags and you'll need to speak [JSON-RPC](http://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before
doing so! Hackers on the internet are actively trying to subvert Earthdollar nodes with exposed APIs!
Further, all browser tabs can access locally running webservers, so malicious webpages could try to
subvert locally available APIs!**

### Operating a private network

Maintaining your own private network is more involved as a lot of configurations taken for granted in
the official networks need to be manually set up.

#### Defining the private genesis state

First, you'll need to create the genesis state of your networks, which all nodes need to be aware of
and agree upon. This consists of a small JSON file (e.g. call it `genesis.json`):

```json
{
  "alloc"      : {},
  "coinbase"   : "0x0000000000000000000000000000000000000000",
  "difficulty" : "0x20000",
  "extraData"  : "",
  "gasLimit"   : "0x2fefd8",
  "nonce"      : "0x0000000000000042",
  "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp"  : "0x00"
}
```
=======
  * `--rpc-addr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpc-port` HTTP-RPC server listening port (default: 8545)
  * `--rpc-api` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
  * `--rpc-cors-domain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws-addr` WS-RPC server listening interface (default: "localhost")
  * `--ws-port` WS-RPC server listening port (default: 8812)
  * `--ws-api` API's offered over the WS-RPC interface (default: "eth,net,web3")
  * `--ws-origins` Origins from which to accept websockets requests
  * `--ipc-disable` Disable the IPC-RPC server
  * `--ipc-api` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
  * `--ipc-path` Filename for IPC socket/pipe within the datadir (explicit paths escape it)
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect via HTTP, WS or IPC to a Ged node configured with the above flags and you'll need to speak [JSON-RPC](http://www.jsonrpc.org/specification) on all transports. You can reuse the same connection for multiple requests!

> Note: Please understand the security implications of opening up an HTTP/WS based transport before doing so! Hackers on the internet are actively trying to subvert Earthdollar nodes with exposed APIs! Further, all browser tabs can access locally running webservers, so malicious webpages could try to subvert locally available APIs!*

<<<<<<< HEAD
With the genesis state defined in the above JSON file, you'll need to initialize **every** Ged node
with it prior to starting it up to ensure all blockchain parameters are correctly set:
=======
### Operating a private/custom network
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

As of [Ged 3.4](https://github.com/Tzunami/go-earthdollar/releases) you are now able to configure a private chain by specifying an __external chain configuration__ JSON file, which includes necessary genesis block data as well as feature configurations for protocol forks, bootnodes, and chainID.

Please find full [example  external configuration files representing the Mainnet and Morden Testnet specs in the /config subdirectory of this repo](). You can use either of these files as a starting point for your own customizations.

It is important for a private network that all nodes use compatible chains. In the case of custom chain configuration, the chain configuration file (`chain.json`) should be equivalent for each node.

#### Define external chain configuration

Specifying an external chain configuration file will allow fine-grained control over a custom blockchain/network configuration, including the genesis state and extending through bootnodes and fork-based protocol upgrades.

```shell
$ ged --chain=morden dump-chain-config <datadir>/customnet/chain.json
$ sed s/mainnet/customnet/ <datadir>/customnet/chain.json
$ vi <datadir>/customnet/chain.json # make your custom edits
$ ged --chain=customnet [--flags] [command]
```
<<<<<<< HEAD
$ ged init path/to/genesis.json
```
=======
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

The external chain configuration file specifies valid settings for the following top-level fields:

| JSON Key | Notes |
| --- | --- |
| `chainID` |  Chain identity. Determines local __/subdir__ for chain data, with required `chain.json` located in it. It is required, but must not be identical for each node. Please note that this is _not_ the chainID validation introduced in _EIP-155_; that is configured as a protocal upgrade within `forks.features`. |
| `name` | _Optional_. Human readable name, ie _Earthdollar Mainnet_, _Morden Testnet._ |
| `state.startingNonce` | _Optional_. Initialize state db with a custom nonce. |
| `network` | Determines Network ID to identify valid peers. |
| `consensus` | _Optional_. Proof of work algorithm to use, either "edhash" or "ethast-test" (for development) |
| `genesis` | Determines __genesis state__. If running the node for the first time, it will write the genesis block. If configuring an existing chain database with a different genesis block, it will overwrite it. |
| `chainConfig` | Determines configuration for fork-based __protocol upgrades__, ie _EIP-150_, _EIP-155_, _EIP-160_, _ECIP-1010_, etc ;-). Subkeys are `forks` and `badHashes`. |
| `bootstrap` | _Optional_. Determines __bootstrap nodes__ in [enode format](https://github.com.earthdollarproject/wiki/wiki/enode-url-format). |


*Fields `name`, `state.startingNonce`, and `consensus` are optional. Ged will panic if any required field is missing, invalid, or in conflict with another flag. This renders `--chain` __incompatible__ with `--testnet`. It remains __compatible__ with `--data-dir`.*

To learn more about external chain configuration, please visit the [External Command Line Options Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/Command-Line-Options).

##### Create the rendezvous point

Once all participating nodes have been initialized to the desired genesis state, you'll need to start a __bootstrap node__ that others can use to find each other in your network and/or over the internet. The clean way is to configure and run a dedicated bootnode:

```
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```

<<<<<<< HEAD
With the bootnode online, it will display an [`enode` URL](https://github.com/Tzunami/wiki/wiki/enode-url-format)
that other nodes can use to connect to it and exchange peer information. Make sure to replace the
displayed IP address information (most probably `[::]`) with your externally accessible IP to get the
actual `enode` URL.
=======
With the bootnode online, it will display an `enode` URL that other nodes can use to connect to it and exchange peer information. Make sure to replace the
displayed IP address information (most probably `[::]`) with your externally accessible IP to get the actual `enode` URL.
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

*Note: You could also use a full fledged Ged node as a bootnode, but it's the less recommended way.*

To learn more about enodes and enode format, visit the [Enode Wiki page](https://github.com.earthdollarproject/wiki/wiki/enode-url-format).

<<<<<<< HEAD
With the bootnode operational and externally reachable (you can try `telnet <ip> <port>` to ensure
it's indeed reachable), start every subsequent Ged node pointed to the bootnode for peer discovery
via the `--bootnodes` flag. It will probably also be desirable to keep the data directory of your
private network separated, so do also specify a custom `--datadir` flag.

```
$ ged --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above>
=======
##### Starting up your member nodes

With the bootnode operational and externally reachable (you can try `telnet <ip> <port>` to ensure it's indeed reachable), start every subsequent Ged node pointed to the bootnode for peer discovery via the `--bootnodes` flag. It will probably be desirable to keep private network data separate from defaults; to do so, specify a custom `--datadir` and/or `--chain` flag.

```
$ ged --datadir=path/to/custom/data/folder \
       --chain=kittynet \
       --bootnodes=<bootnode-enode-url-from-above>
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3
```

*Note: Since your network will be completely cut off from the main and test networks, you'll also need to configure a miner to process transactions and create new blocks for you.*

#### Running a private miner

<<<<<<< HEAD
Mining on the public Earthdollar network is a complex task as it's only feasible using GPUs, requiring
an OpenCL or CUDA enabled `edminer` instance. For information on such a setup, please consult the
[EDMining subreddit](https://www.reddit.com/r/EDMining/) and the [Genoil miner](https://github.com/Genoil/cpp.earthdollar)
repository.

In a private network setting however, a single CPU miner instance is more than enough for practical
purposes as it can produce a stable stream of blocks at the correct intervals without needing heavy
resources (consider running on a single thread, no need for multiple ones either). To start a Ged
instance for mining, run it with all your usual flags, extended by:
=======
Mining on the public Earthdollar network is a complex task as it's only feasible using GPUs, requiring an OpenCL or CUDA enabled `ethminer` instance. For information on such a setup, please consult the [EtherMining subreddit](https://www.reddit.com/r/EtherMining/) and the [Genoil miner](https://github.com/Genoil/cpp.earthdollar) repository.

In a private network setting however, a single CPU miner instance is more than enough for practical purposes as it can produce a stable stream of blocks at the correct intervals without needing heavy resources (consider running on a single thread, no need for multiple ones either). To start a Ged instance for mining, run it with all your usual flags, extended by:
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

```
$ ged <usual-flags> --mine --minerthreads=1 --earthbase=0x0000000000000000000000000000000000000000
```

<<<<<<< HEAD
Which will start mining blocks and transactions on a single CPU thread, crediting all proceedings to
the account specified by `--earthbase`. You can further tune the mining by changing the default gas
limit blocks converge to (`--targetgaslimit`) and the price transactions are accepted at (`--gasprice`).
=======
Which will start mining blocks and transactions on a single CPU thread, crediting all proceedings to the account specified by `--earthbase`. You can further tune the mining by changing the default gas limit blocks converge to (`--targetgaslimit`) and the price transactions are accepted at (`--gasprice`).

For more information about managing accounts, please see the [Managing Accounts Wiki page](https://github.com/Tzunami/go-earthdollar/wiki/Managing-Accounts).

>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

## Contribution

Thank you for considering to help out with the source code!

<<<<<<< HEAD
If you'd like to contribute to go-earthdollar, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. If you wish to submit more
complex changes though, please check up with the core devs first on [our gitter channel](https://gitter.im/Tzunami/go-earthdollar)
to ensure those changes are in line with the general philosophy of the project and/or get some
early feedback which can make both your efforts much lighter as well as our review and merge
procedures quick and simple.
=======
The core values of democratic engagement, transparency, and integrity run deep with us. We welcome contributions from everyone, and are grateful for even the smallest of fixes.  :clap:
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3

This project is migrated from the now hard-forked [Earthdollar (ETHF) Github project](https://github.com.earthdollar), and we will need to incrementally migrate pieces of the infrastructure required to maintain the project. 

<<<<<<< HEAD
 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `develop` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "ed, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/Tzunami/go-earthdollar/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies and testing procedures.

## License

The go-earthdollar library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-earthdollar binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
=======
If you'd like to contribute to go.earthdollar, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base. If you wish to submit more complex changes, please check up with the core devs first on [our Slack channel (#development)](http:/.earthdollarclassic.herokuapp.com/) to ensure those changes are in line with the general philosophy of the project and/or get some early feedback which can make both your efforts much lighter as well as our review and merge procedures quick and simple.

Please see the [Wiki](https://github.com/Tzunami/go-earthdollar/wiki) for more details on configuring your environment, managing project dependencies, and testing procedures.

## License

The go.earthdollar library (i.e. all code outside of the `cmd` directory) is licensed under the [GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the `COPYING.LESSER` file.

The go.earthdollar binaries (i.e. all code inside of the `cmd` directory) is licensed under the [GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `COPYING` file.
>>>>>>> 462a0c24946f17de60f3ba1226255a938bc47de3
