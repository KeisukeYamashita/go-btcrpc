/*
Package btcrpc implements RPC methods to interact with the bitcoin node bitcoind.
See informations in the pages link below about Bitcoin development.
https://bitcoin.org/en/development
*/
package btcrpc

import (
	"encoding/json"
	"errors"

	jsonrpc "github.com/KeisukeYamashita/go-jsonrpc"
)

/*
BasicAuth is for Bitcoin Node supports basic auth.
Some nodes do not need this. In that case, leave these blank.
*/
type BasicAuth struct {
	Username string
	Password string
}

/*
RPCClient ...
*/
type RPCClient struct {
	*jsonrpc.RPCClient
}

/*
RPCer ...
*/
type RPCer interface {
	GetBlockHash(height int32) (string, error)
}

/*
NewRPCClient creates JSONRPC clients for your bitcoin node.
*/
func NewRPCClient(endpoint string, basicAuth *BasicAuth) *RPCClient {
	c := new(RPCClient)
	c.RPCClient = jsonrpc.NewRPCClient(endpoint)
	c.RPCClient.SetBasicAuth(basicAuth.Username, basicAuth.Password)
	return c
}

/*
GetNewAddress gets new address associated with the account name given.
If the account name is blank(nil), if will also returns a addresss with no associated account.
*/
func (c *RPCClient) GetNewAddress(account string) (string, error) {
	resp, err := c.RPCClient.Call("getnewaddress", account)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var address string
	resp.GetObject(&address)
	return address, nil
}

/*
GetBalance gets the balance of the address.
It is only possible to get the balance which is made by this node, otherwise it will return 0.00.
*/
func (c *RPCClient) GetBalance(address string) (float32, error) {
	resp, err := c.RPCClient.Call("getbalance", address)
	if err != nil {
		return -1, err
	}

	if resp.Error != nil {
		return -1, errors.New(resp.Error.Message)
	}

	var balance float32
	resp.GetObject(&balance)
	return balance, nil
}

/*
GetBlockHash gets the block hash(id) associated with the block height.
*/
func (c *RPCClient) GetBlockHash(height int32) (string, error) {
	resp, err := c.RPCClient.Call("getblockhash", height)
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var hash string
	resp.GetObject(&hash)
	return hash, nil
}

/*
GetBlock gets the block information associated with the block hash(id).
It contains a lot of infos about transactions.
*/
func (c *RPCClient) GetBlock(h string) (*Block, error) {
	resp, err := c.RPCClient.Call("getblock", h)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.New(resp.Error.Message)
	}

	jsonData, err := json.Marshal(resp.Result)
	if err != nil {
		return nil, err
	}

	var block Block
	json.Unmarshal(jsonData, &block)
	return &block, nil
}

/*
GetBlockCount gets the latest block height.
Note that The methods name is "Count" not "Height".
*/
func (c *RPCClient) GetBlockCount() (int32, error) {
	resp, err := c.RPCClient.Call("getblockcount")
	if err != nil {
		return -1, err
	}

	if resp.Error != nil {
		return -1, errors.New(resp.Error.Message)
	}

	var count int32
	resp.GetObject(&count)
	return count, nil
}

/*
GetRawTransactions gets the raw transactions associated with the transaction hashes(ids).
*/
func (c *RPCClient) GetRawTransactions(txids []string) ([]string, error) {
	rawTxs := make([]string, len(txids))
	for i, txid := range txids {
		resp, err := c.RPCClient.Call("getrawtransaction", txid)
		if err != nil {
			return nil, err
		}

		if resp.Error != nil {
			return nil, errors.New(resp.Error.Message)
		}

		var rawTx string
		resp.GetObject(&rawTx)
		rawTxs[i] = rawTx
	}
	return rawTxs, nil
}

/*
DecodeRawTransactions decodes the raw transactions to human readable transactions.
*/
func (c *RPCClient) DecodeRawTransactions(rawTxs []string) ([]*Transaction, error) {
	txs := make([]*Transaction, len(rawTxs))
	for i, rawTx := range rawTxs {
		resp, err := c.RPCClient.Call("decoderawtransaction", rawTx)
		if err != nil {
			return nil, err
		}

		if resp.Error != nil {
			return nil, errors.New(resp.Error.Message)
		}

		jsonData, err := json.Marshal(resp.Result)
		if err != nil {
			return nil, err
		}

		var tx Transaction
		json.Unmarshal(jsonData, &tx)
		txs[i] = &tx
	}
	return txs, nil
}
