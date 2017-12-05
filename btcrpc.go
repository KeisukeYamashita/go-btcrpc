package btcrpc

import (
	"context"
	"github.com/GuiltyMorishita/jsonrpc"
	"google.golang.org/appengine/urlfetch"
)

// BtcRPCer ...
type BtcRPCer interface {
  GetBalance(address, block string) (balance string, err error)
	GetTransactionCount(address, block string) (count uint64, err error)
	SendRawTransaction(txData string) (txHash string, err error)
	UseAppEngineContext(ctx context.Context)
}

// BtcRPC ...
type BtcRPC struct {
  rpcClient *jsonrpc.RPCClient
}

// NewBtcRPC ...
func NewBtcRPC(endpoint string) *BtcRPC {
	return &BtcRPC{
		rpcClient: jsonrpc.NewRPCClient(endpoint),
	}
}

func (rpc *BtcRPC) GetBalance(address string) (balance string, err error) {
	response, err := rpc.rpcClient.Call("getbalance", address)
	if err != nil {
		return
	}

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
