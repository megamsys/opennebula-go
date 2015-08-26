package virtualmachine

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/megamsys/opennebula-go/api"
)

func TestGetVirtualMachineByName(t *testing.T) {
	Convey("GetVirtualMachineByName", t, func() {

		endpoint := "http://localhost:2633/RPC2"
		client, _ := api.NewRPCClient(endpoint, "oneadmin", "RaifZuewjoc4")
		vm := VirtualMachineReqs{VMName: "yeshapp", Client: &client}
		vm.GetVirtualMachineByName()
		//So(error, ShouldBeNil)

	})
}
