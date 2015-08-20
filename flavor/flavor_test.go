package flavor

import (
	"fmt"
	"testing"
)

func TestGetTemplatePool(t *testing.T) {
	Convey("GetTemplatePool", t, func() {
		endpoint := "http://localhost:2633/RPC2"
		client, _ := RPCClient(endpoint)
		key := "oneadmin:RaifZuewjoc4"
		args := []interface{}{key, -2, 0, 0}
		error := Call(client, "one.templatepool.info", args)
		fmt.Println(error)
		So(error, ShouldBeNil)
	})
}
