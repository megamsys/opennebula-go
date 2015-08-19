package api

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateNewRPCClient(t *testing.T) {
	Convey("CreateNewRPCClient", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		client, err := RPCClient(endpoint)
		So(err, ShouldBeNil)
	})
}
