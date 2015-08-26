package compute

import (
	"fmt"
	"testing"

	"github.com/megamsys/opennebula-go/api"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreate(t *testing.T) {
	Convey("Create", t, func() {

		client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
		vmObj := VirtualMachine{Name: "yeshapp", TemplateName: "fullfledged", Cpu: "1", VCpu: "1", Memory: "4500", Client: &client} //memory in terms of MB! duh!

		res := vmObj.Create()
		fmt.Println(res)
		//	So(err, ShouldBeNil)
	})
}
