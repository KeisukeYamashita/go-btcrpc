package btcrpc

// Block ...
type Block struct {
	Hash          string   `json:"hash"`
	Version       int32    `json:"version"`
	Txs           []string `json:"txs"`
	Time          int32    `json:"time"`
	Size          int32    `json:"size"`
	Nonce         int32    `json:"nonce"`
	Weight        int32    `json:"weight"`
	VersionHex    string   `json:"versionHex"`
	Difficulty    int32    `json:"difficulty"`
	Mediantime    int32    `json:"mediantime"`
	Chainwork     string   `json:"chainwork"`
	Strippedsize  int32    `json:"strippedsize"`
	Merkleroot    string   `json:"merkleroot"`
	Bits          string   `json:"bits"`
	NextBlockhash string   `json:"nextBlockhash"`
	Confirmations int32    `json:"confirmations"`
	Height        int32    `json:"height"`
}

// Transaction ...
type Transaction struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int32  `json:"id"`
	Size     int32  `json:"size"`
	Vsize    int32  `json:"vsize"`
	Locktime int32  `json:"locktime"`
	Vins     Vin
	Vouts    []Vout
}

// Vin ...
type Vin map[int]interface{}

// Vout ...
type Vout struct {
	Value        float32 `json:"value"`
	N            int32   `json:"n"`
	ScriptPubKey ScriptPubKey
}

// ScriptPubKey ...
type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int32    `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

// VinTransaction is a struct of Vin(inputs) with normal transaction.
type VinTransaction struct {
	Txid        string `json:"txid"`
	Vout        int32  `json:"vout"`
	ScriptSig   ScriptSig
	TxinWitness []string `json:"txinWitness"`
	Sequence    int64    `json:"Sequence"`
}

// VinCoinbaseTransaction is a struct of Vins(inputs) with mining.
type VinCoinbaseTransaction struct {
	Coinbase string `json:"coinbase"`
	Sequence int64  `json:"sequence"`
}

// ScriptSig ...
type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}
