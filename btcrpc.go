package btcrpc

import (
	"context"
	"github.com/k0kubun/pp"
	"github.com/GuiltyMorishita/jsonrpc"
	"google.golang.org/appengine/urlfetch"
)

// EthRPCer ...
type BtcRPCer interface {
	GetBalance(address, block string) (balance string, err error)
	GetTransactionCount(address string) (count uint64, err error)
	SendRawTransaction(txData string) (txHash string, err error)
	UseAppEngineContext(ctx context.Context)
}

// BtcRPC ...
type BtcRPC struct {
	rpcClient *jsonrpc.RPCClient
}

// NewEthRPC ...
func NewBtcRPC(endpoint string) *BtcRPC {
	return &BtcRPC{
		rpcClient: jsonrpc.NewRPCClient(endpoint),
	}
}

func (rpc *BtcRPC) GetBalance(username, password, account string) (balance string, err error) {
	rpc.rpcClient.SetBasicAuth(username, password)
	response, err := rpc.rpcClient.Call("getbalance", account)
	pp.Println(response)
	if err != nil {
		return
	}
	pp.Println(response==nil)
	pp.Println(response.Error==nil)

	if response.Error != nil {
		err = response.Error
		return
	}

	response.GetObject(&balance)
	return
}

func (rpc *BtcRPC) UseAppEngineContext(ctx context.Context) {
	rpc.rpcClient.SetHTTPClient(urlfetch.Client(ctx))
}
