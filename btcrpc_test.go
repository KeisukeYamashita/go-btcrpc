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
		f(NewBtcRPC(os.Getenv("BTCD_ENDPOINT"), os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	}
}


func TestGetBalance(t *testing.T) {
	Convey("WithBtcRPC", t, WithBtcRPC(func(rpc *BtcRPC) {
		Convey("Success", func() {
			balance, err := rpc.GetBalance("hogehoge")
			So(err, ShouldBeNil)
			So(balance, ShouldEqual, 0.00000)
		})

		Convey("Invalid Basic Auth", func() {
			_, err := rpc.GetBalance("hogehoge")
			So(err, ShouldContainString, "Invalid Basic Auth")
		})
	}))
}
