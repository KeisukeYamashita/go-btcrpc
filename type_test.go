package btcrpc

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTransactionStruct(t *testing.T) {
	VinMap := make(map[int]interface{})

	VinMap[1] = VinTransaction{
		Txid: "09409e9151aeed67ed592c181840855c1f45d6ce3c1b8b0b3616f4dea2c22ccb",
		Vout: 0,
		ScriptSig: ScriptSig{
			Asm: "",
			Hex: "",
		},
		TxinWitness: []string{"", "3045022100a6cda8a54b4fee586fe8cf2bde6ecdef919823410cae86ee7907a538893f47c602201304b82d1a7b6b162be2197b433bd645241f17aa8c70a08040474598aac5ac9e01", "30450221008829bb92c399c1182f67ebffb1834375d7d206f6fbf1536a39b4cfaa29374f9002200f4dcf8cf07fd257be9d2338de9e0041b4bf1c16a652182296632a44a912880d01"},
		Sequence:    4294967295,
	}
	VinMap[2] = VinCoinbaseTransaction{
		Coinbase: "043jifdga345",
		Sequence: 4294967295,
	}

	Vout := []Vout{
		Vout{
			Value: 50.00000000,
			N:     0,
			ScriptPubKey: ScriptPubKey{
				Asm:       "0496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858ee OP_CHECKSIG",
				Hex:       "410496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858eeac",
				ReqSigs:   1,
				Type:      "pubkey",
				Addresses: []string{"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"},
			},
		},
		Vout{
			Value: 1000.00000000,
			N:     1,
			ScriptPubKey: ScriptPubKey{
				Asm:       "0496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858ee OP_CHECKSIG",
				Hex:       "410496b538e853519c726a2c91e61ec11600ae1390813a627c66fb8be7947be63c52da7589379515d4e0a604f8141781e62294721166bf621e73a82cbf2342c858eeac",
				ReqSigs:   1,
				Type:      "pubkey",
				Addresses: []string{"1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX", "1KJDckyDF54pH65quicuyCK7UH4Af7Wtpp"},
			},
		},
	}
	rawTx := &Transaction{
		Txid:     "0e3e2357e806b6cdb1f70b54c3a3a17b6714ee1f0e68bebb44a74b1efd512098",
		Hash:     "0e3e2357e806b6cdb1f70b54c3a3a17b6714ee1f0e68bebb44a74b1efd512098",
		Version:  1,
		Size:     134,
		Vsize:    134,
		Locktime: 0,
		Vins:     VinMap,
		Vouts:    Vout,
	}
	Convey("SuccessWithVin", t, func() {
		var VinCoinbaseTranscactionNum int
		var VinTransactionNum int
		for i, Vin := range rawTx.Vins {
			if _, ok := Vin.(VinTransaction); ok {
				VinTransactionNum = i
			}
			if _, ok := Vin.(VinCoinbaseTransaction); ok {
				VinCoinbaseTranscactionNum = i
			}
		}
		So(VinTransactionNum, ShouldEqual, 1)
		So(VinCoinbaseTranscactionNum, ShouldEqual, 2)
	})
}
