# go-btcrpc

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/KeisukeYamashita/go-btcrpc)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![CircleCI](https://circleci.com/gh/KeisukeYamashita/go-btcrpc.svg?style=svg)](https://circleci.com/gh/KeisukeYamashita/go-btcrpc)
[![Job Status](https://inspecode.rocro.com/badges/github.com/KeisukeYamashita/go-btcrpc/status?token=SGqr7pQjbMTQuMDLOPk_rvq_hGeF_hoLj_B7tbRKSXg)](https://inspecode.rocro.com/jobs/github.com/KeisukeYamashita/go-btcrpc/latest?completed=true)
[![Report](https://inspecode.rocro.com/badges/github.com/KeisukeYamashita/go-btcrpc/report?token=SGqr7pQjbMTQuMDLOPk_rvq_hGeF_hoLj_B7tbRKSXg&branch=master)](https://inspecode.rocro.com/reports/github.com/KeisukeYamashita/go-btcrpc/branch/master/summary)


go-btcrpc is a Go library use for interacting to the bitcoin node from your server with JSON-RPC which is a standard protocol for blockchain.

This library provides the easiest way to send JSON-RPC.

![image.png](https://qiita-image-store.s3.amazonaws.com/0/153320/c6806795-ffc1-6579-6dd3-38948fb0b851.png)

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

| method| discription |
|:----:|:----:|
| getNewAddress | creates a new address to your account. |
| getBalance | get the balance of your address. |
| getBlockHash | get the hash of the block. |
| getBlock | get the block by hash of the block. |
| getBlockCount | get the newest block count by hash of the block. |
| decodeTransaction | decode the raw transaction to humanreadable transaction by hash of the block. |
| getrawTransaction | get the raw transaction hash block count by hash of the block. |


## Use tests
Set up your environmental valiables in `.env` to conduct this test.

```
cp .env.sample .env
```

Then write in your endpoint in this file.


Finally run your test. It will pass if your bitcoin node is setted up correctly.

```
GO_ENV = test go test btcrpc
```

## Contribution
To contribute, just send me a pull request!
If it is valid, you will be added on the contribution doc in `/doc/contributor.md` .

## License
Copyright 2017 Keisuke Yamashita.
Licensed under the MIT License.
