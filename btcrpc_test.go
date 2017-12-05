
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
			balance, err := rpc.GetBalance("0x8FfCf7674ED27c7949Ceda9a0bD6799fe74aCf47")
			So(err, ShouldBeNil)
			So(balance, ShouldBeGreaterThan, 0)
		})

		Convey("Empty Address", func() {
			_, err := rpc.GetBalance("")
			So(err.Error(), ShouldContainSubstring, "hex string has length 0")
		})

		Convey("Invalid Address", func() {
			_, err := rpc.GetBalance("InvalidAddress")
			So(err.Error(), ShouldContainSubstring, "cannot unmarshal hex string without 0x prefix ")
		})
	}))
}
