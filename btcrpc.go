package btcrpc

import (
	"encoding/json"
	"errors"

	"github.com/KeisukeYamashita/jsonrpc"
)

// BasicAuth ...
type BasicAuth struct {
	Username string
	Password string
}

// RPCClient ...
type RPCClient struct {
	*jsonrpc.RPCClient
}

// RPCer ...
type RPCer interface {
	GetBlockHash(height int32) (string, error)
}

// NewRPCClient ...
func NewRPCClient(endpoint string, basicAuth *BasicAuth) *RPCClient {
	c := new(RPCClient)
	c.RPCClient = jsonrpc.NewRPCClient(endpoint)
	c.RPCClient.SetBasicAuth(basicAuth.Username, basicAuth.Password)
	return c
}

// GetBalance ...
func (c *RPCClient) GetBalance(address string) (float32, error) {
	resp, err := c.RPCClient.Call("getblockhash", []string{address})
	if err != nil {
		return "", err
	}

	if resp.Error != nil {
		return "", errors.New(resp.Error.Message)
	}

	var balance float32
	resp.GetObject(&balance)
	return balance, nil
}

// GetBlockHash ...
func (c *RPCClient) GetBlockHash(height int32) (string, error) {
	resp, err := c.RPCClient.Call("getblockhash", []int32{height})
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

// GetBlock ...
func (c *RPCClient) GetBlock(h string) (*Block, error) {
	resp, err := c.RPCClient.Call("getblock", []string{h})
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

// GetBlockCount ...
func (c *RPCClient) GetBlockCount() (int32, error) {
	resp, err := c.RPCClient.Call("getblockcount", []interface{}{})
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

// GetRawTransactions ...
func (c *RPCClient) GetRawTransactions(txids []string) ([]string, error) {
	rawTxs := make([]string, len(txids))
	for i, txid := range txids {
		resp, err := c.RPCClient.Call("getrawtransaction", []string{txid})
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

// DecodeRawTransactions ...
func (c *RPCClient) DecodeRawTransactions(rawTxs []string) ([]*Transaction, error) {
	txs := make([]*Transaction, len(rawTxs))
	for i, rawTx := range rawTxs {
		resp, err := c.RPCClient.Call("decoderawtransaction", []string{rawTx})
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
