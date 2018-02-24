# go-btcrpc

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-btcrpc is a Go library use for interacting to the bitcoin node from your server with JSON-RPC which is a standard protocol for blockchain.

This library provides the easiest way to send JSON-RPC.

## Installation
Use go get to install and update.

```
go get github.com/KeisukeYamashita/go-btcrpc.git
```

## Setup
You need to setup environmental variables.

Firstly, copy the `.env.sample` as `.env`

```
cp .env.sample .env
```

Setup in your `.env`.

```
BTCD_ENDPOINT: NODE_ENDPOINT
USERNAME: USERNAME_FOR_BASICAUTH
PASSWORD: PASSWORD_FOR_BASICAUTH
```

## Usage and Example
This shows you that easiest request to the node which is getting the infos.

```go
package main

import (
  "fmt"
  btcrpc "github.com/KeisukeYamashita/go-btcrpc"
  )

func main() {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	address := "my88QLpf2RYYDdNMmDwYvfx6TFc6NXaELa"
	balance := c.GetBalance(address)
	fmt.Print(balance) // 0.13514 BTC
}
```

### Equal curl command

```
curl -X "POST" "<YOUR_BITCOIN_NODE>" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -u '<YOUR_USER_NAME>:<YOUR_PASSWORD>' \
     -d $'{
  "method": "getbalance",
  "id": "1",
  "params": [
    "my88QLpf2RYYDdNMmDwYvfx6TFc6NXaELa"
  ]
}'
```

It'll return a JSON.

### Available Methods

```
- "getNewAddress"          creates a new address to your account.
- "getBalance"          get the balance of your address.
- "getBlockHash"          get the hash of the block.
- "getBlock"          get the block by hash of the block.
- "getBlockCount"          get the newest block count by hash of the block.
- "decodeTransaction"          decode the raw transaction to humanreadable transaction by hash of the block.
- "getrawTransaction"          get the raw transaction hash block count by hash of the block.
```

## Use tests
Set up your environmental valiables in `.env` to conduct this test.

```
cp .env.sample .env
```

Then write in your endpoint in this file.


Finally run your test. It will pass if your bitcoin node is setted up correctly.

```
go test btcrpc
```

## Contribution
To contribute, just send me a pull request!
If it is valid, you will be on the contribution doc.

## License
Copyright 2017 Keisuke Yamashita.
Licensed under the MIT License.
