# go-btcrpc

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-btcrpc is a Go library use for interacting to the bitcoin node from your server with JSON-RPC which is a standard protocol for blockchain.

This library provides the easiest way to send JSON-RPC, which is not a major standard for

## Installation
Use go get to install and update.

```
go get github.com/BlockChainUtils/go-btcrpc.git
```

<<<<<<< HEAD
=======
## Setup
You need to setup enviromental variables.

Firstly, copy the `.test.env.sample` as `.test.env`

```
cp test.env.sample .test.env
```

Setup in your `.test.env`.

```
BTCD_ENDPOINT: NODE_ENDPOINT
USERNAME: USERNAME_FOR_BASICAUTH
PASSWORD: PASSWORD_FOR_BASICAUTH
```

>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
## Usage and Example
This shows you that easiest request to the node which is getting the infos.

```
package main

import (
  "fmt"
  "github.com/BlockChainUtils/go-btcrpc"
  )

func main(){
<<<<<<< HEAD
  info, err = rpc.GetInfo(username, password)
=======
  info, err := rpc.GetInfo(username, password)
>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
  fmt.Print(info)
}
```

<<<<<<< HEAD
### curl command
=======
### Equal curl command
>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
This is equal to this curl command from your terminal.

```
curl -X "POST" "<YOUR_BITCOIN_NODE>" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -u '<YOUR_USER_NAME>:<YOUR_PASSWORD>' \
     -d $'{
<<<<<<< HEAD
  "jsonrpc": "1.0",
=======
>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
  "method": "getinfo",
  "id": "1",
  "params": []
}'
```

It'll return a JSON.

### Available Methods
<<<<<<< HEAD
```
- "getbalance"          get the balance of your node.
- "getbalance"          get the balance of your node.
- "getbalance"          get the balance of your node.
- "getbalance"          get the balance of your node.
=======

```
- "getbalance"          get the balance of your node.
>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
```

## Use tests
Set up your enviromental valiables in `.test.env` to conduct this test.

```
cp .test.env.sample .test.env
```

Then write in your endpoint in this file.


Finally run your test. It will pass if your bitcoin node is setted up correctly.

```
go test bitrpc
```

## For whose also intersted in Ethereum
We are also providing a similar library with Ethereum.
Look at this [repository](https://github.com/BlockChainUtils/go-ethrpc).

## License
<<<<<<< HEAD
MIT
=======
Copyright 2017 Keisuke Yamashita.

Licensed under the MIT License.
>>>>>>> c03f471c3aa8a46f92e2493dcfe141f551c9dddc
