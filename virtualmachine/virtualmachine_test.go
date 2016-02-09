package virtualmachine

import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
)
func (s *S) TestGetVirtualMachineByName(c *check.C) {
	client, _ := api.NewRPCClient("http://192.168.1.105:2633/RPC2", "oneadmin", "yourWuOtHij3")
	vm := VirtualMachineReqs{VMName: "vijayorion", Client: client}
	res, error := vm.GetVirtualMachineByName()

	c.Assert(error, check.IsNil)
	c.Assert(res, check.NotNil)
}
