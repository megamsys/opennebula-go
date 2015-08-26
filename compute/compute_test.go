package compute

import (
	"testing"

	"github.com/megamsys/opennebula-go/api"
	. "github.com/smartystreets/goconvey/convey"
)

/*
func TestCreate(t *testing.T) {
	Convey("Create", t, func() {

		client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
		vmObj := VirtualMachine{Name: "yeshapp", TemplateName: "fullfledged", Cpu: "1", VCpu: "1", Memory: "4500", Client: &client} //memory in terms of MB! duh!

		res := vmObj.Create()
		fmt.Println(res)
		//	So(err, ShouldBeNil)
	})
}
*/
func TestDelete(t *testing.T) {
	Convey("Delete", t, func() {

		client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
		//	client, _ := api.NewRPCClient("http://192.168.1.100:2633/RPC2", "oneadmin", "yib4OquafUp1")
		vmObj := VirtualMachine{Name: "yesh", Client: &client} //memory in terms of MB! duh!

		_ = vmObj.Delete()
		//fmt.Println(res)
		//	So(err, ShouldBeNil)
	})
}
