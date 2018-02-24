package btcrpc

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("GO_ENV") != "test" {
		log.Fatal("Error loading .env file")
	}
}

func TestRPCClient(t *testing.T) {
	Convey("Success", t, func() {
		basicAuth := &BasicAuth{
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
		}
		c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
		_, err := c.GetBlockCount()
		So(err, ShouldBeNil)
	})
	Convey("InvalidBasicAuth", t, func() {
		basicAuth := new(BasicAuth)
		c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
		_, err := c.GetBlockCount()
		So(err, ShouldNotBeNil)
	})
}

func TestNewAddress(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		_, err := c.GetNewAddress("KeisukeYamashita")
		So(err, ShouldBeNil)
	})
}

func TestGetBalance(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.GetBalance("address")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, 0.00000000) // Block Height
	})

	Convey("No addresss", t, func() {
		result, err := c.GetBalance("I do not exists")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, 0.00000000) // 0.00000000
	})
}

func TestGetBlockHash(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.GetBlockHash(500954)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "00000000000a098712c8c196b42dfcb36ecf77c4620ba4719c868012ebaeec8a") // Block Hash
	})

	Convey("Invalid block index", t, func() {
		_, err := c.GetBlockHash(-1)
		So(err.Error(), ShouldEqual, "Block height out of range")
	})
}

// TODO: change for testnet*testing.T) {
// 	basicAuth := &BasicAuth{
// 		Username: os.Getenv("USERNAME"),
// 		Password: os.Getenv("PASSWORD"),
// 	}
// 	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)

// 	Convey("Invalid block hash", t, func() {
// 		_, err := c.GetBlock("Invalid transactionId")
// 		So(err.Error(), ShouldEqual, "Block not found")
// 	})
// }

func TestGetBlockCount(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		_, err := c.GetBlockCount()
		So(err, ShouldBeNil)
	})
}

// TODO: change for testnet
// func TestGetRawTransactions(t *testing.T) {
// 	basicAuth := &BasicAuth{
// 		Username: os.Getenv("USERNAME"),
// 		Password: os.Getenv("PASSWORD"),
// 	}
// 	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
// 	Convey("Success", t, func() {
// 		result, err := c.GetRawTransactions([]string{"8878522366f6a4d9a7b4a88dca6d36b036f09d6093c1d39b8cf7bedd4dbf75d6", "abb3880f8f7d66effa540e8ca9ce8878e8b4978c9843ee8c3001c07d09ae5ce3"}) // Genesis Block
// 		So(err, ShouldBeNil)
// 		So(result[0], ShouldEqual, "010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff5503daa40741d6902f5030707d41d6902f4fc582f22f4254432e544f502ffabe6d6d3b1daff61de82061d499ed27ab8e058e615edea9175c01c8134c5b0f3321e2218000000000000000c100676b0000a1aeef6d0100ffffffff022e5cb258000000001976a914ba507bae8f1643d2556000ca26b9301b9069dc6b88ac0000000000000000266a24aa21a9ed8a7f6c7f59f1859daf7475f1d28be4113abb8b5365d4868b69b877a4f4665f280120000000000000000000000000000000000000000000000000000000000000000000000000")
// 		So(result[1], ShouldEqual, "0100000001962328b4c86763f92ecf78180f015c236e53f34cb6860780d539f8a4072b8501060000006b483045022100c8c78ceb7d1affe434fe4952f32d066d1cd8e1dc418f10e752afaa2cf924629e0220615ca51233d3510868e3416ab4e8b0f7b83e68839716198586325abf2775dfe2012103f803ea646b6bdeb9ad2ebb486724e3a82a1b211666a9830714ba94e6214dfeabffffffff02d40a0200000000001976a914c7c04f2c8dfa53f542ce8abf49bd47ba06804a4788ac4d483700000000001976a914f7d6366850ec376c277e45b7c6984c126d25904e88ac00000000")
// 	})

// 	Convey("Invalid transactionId", t, func() {
// 		_, err := c.GetRawTransactions([]string{"Invalid transactionId"})
// 		So(err.Error(), ShouldContainSubstring, "parameter 1 must be hexadecimal string")
// 	})
// }

func TestDecodeRawTransaction(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("BTCD_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.DecodeRawTransactions([]string{"010000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff5503daa40741d6902f5030707d41d6902f4fc582f22f4254432e544f502ffabe6d6d3b1daff61de82061d499ed27ab8e058e615edea9175c01c8134c5b0f3321e2218000000000000000c100676b0000a1aeef6d0100ffffffff022e5cb258000000001976a914ba507bae8f1643d2556000ca26b9301b9069dc6b88ac0000000000000000266a24aa21a9ed8a7f6c7f59f1859daf7475f1d28be4113abb8b5365d4868b69b877a4f4665f280120000000000000000000000000000000000000000000000000000000000000000000000000", "0100000001962328b4c86763f92ecf78180f015c236e53f34cb6860780d539f8a4072b8501060000006b483045022100c8c78ceb7d1affe434fe4952f32d066d1cd8e1dc418f10e752afaa2cf924629e0220615ca51233d3510868e3416ab4e8b0f7b83e68839716198586325abf2775dfe2012103f803ea646b6bdeb9ad2ebb486724e3a82a1b211666a9830714ba94e6214dfeabffffffff02d40a0200000000001976a914c7c04f2c8dfa53f542ce8abf49bd47ba06804a4788ac4d483700000000001976a914f7d6366850ec376c277e45b7c6984c126d25904e88ac00000000"})
		So(err, ShouldBeNil)
		So(result[0].Txid, ShouldEqual, "8878522366f6a4d9a7b4a88dca6d36b036f09d6093c1d39b8cf7bedd4dbf75d6")
		So(result[1].Txid, ShouldEqual, "abb3880f8f7d66effa540e8ca9ce8878e8b4978c9843ee8c3001c07d09ae5ce3")
	})

	Convey("Invalid rawtx", t, func() {
		_, err := c.DecodeRawTransactions([]string{"Invalid rawTransaction"})
		So(err.Error(), ShouldEqual, "TX decode failed")
	})

}
