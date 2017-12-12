package btcrpc

import (
	"context"
	"github.com/KeisukeYamashita/jsonrpc"
	"google.golang.org/appengine/urlfetch"
)

// BtcRPC sends JSON-RPC1.0 request over http to the provided rpc endpoint.
// BtcRPC is created using the factory function NewBtcRPC().
type BtcRPC struct {
	rpcClient *jsonrpc.RPCClient
}

// NewBtcRPC returns a new RPCClient instance with endpoint and basic authentication.
// One you conduct this function you can send HTTP request with basic authentication anytimes.
func NewBtcRPC(endpoint, username, password string) *BtcRPC {
	basicAuth = &jsonrpc.BasicAuth{
		username: username,
		password: password,
	}
	return &BtcRPC{
		rpcClient: jsonrpc.NewRPCClient(endpoint, basicAuth),
	}
}

func (rpc *BtcRPC) GetBalance(account string) (balance float32, err error) {
	response, err := rpc.rpcClient.Call("getbalance", account)
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
