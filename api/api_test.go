package api

import (
	"fmt"
	"testing"

	//	"github.com/megamsys/opennebula-go/api"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateNewRPCClient(t *testing.T) {
	Convey("CreateNewRPCClient", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		_, err := RPCClient(endpoint)
		So(err, ShouldBeNil)
	})
}

func TestRPCCall(t *testing.T) {
	Convey("RPCCall", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		client, _ := RPCClient(endpoint)
		key := "oneadmin:RaifZuewjoc4"
		args := []interface{}{key, -2, 3, 3}
		_, error := Call(client, "one.templatepool.info", args)
		fmt.Println(error)
		So(error, ShouldBeNil)
	})
}
