package btcrpc

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load(".test.env")
	exitCode := m.Run()
	defer os.Exit(exitCode)
}

func WithBtcRPC(f func(rpc *BtcRPC)) func() {
	return func() {
		f(NewBtcRPC(os.Getenv("BTCD_ENDPOINT")))
	}
}


func TestGetBalance(t *testing.T) {
	Convey("WithBtcRPC", t, WithBtcRPC(func(rpc *BtcRPC) {
		Convey("Success", func() {
			balance, err := rpc.GetBalance("user","JgMchk9GTtGPqVYdTG2bpHMJ", "1JxEFrgYHF51HY2qsE6erYEAdMxTG4iXu4")
			So(err, ShouldBeNil)
			So(balance, ShouldEqual, "0x56bc75e2d63100000") // 100 ETH
		})

		Convey("Empty Address", func() {
			_, err := rpc.GetBalance("user","JgMchk9GTtGPqVYdTG2bpHMJ", "")
			So(err.Error(), ShouldContainSubstring, "hex string has length 0")
		})

		Convey("Invalid Address", func() {
			_, err := rpc.GetBalance("user","JgMchk9GTtGPqVYdTG2bpHMJ", "InvalidAddress")
			So(err.Error(), ShouldContainSubstring, "cannot unmarshal hex string without 0x prefix ")
		})
	}))
}
