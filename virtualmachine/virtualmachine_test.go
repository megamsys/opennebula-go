package virtualmachine

import (
	"fmt"
	"testing"

	"github.com/megamsys/opennebula-go/api"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestGetVirtualMachineByName(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vm := VirtualMachineReqs{VMName: "yeshapp", Client: &client}
	res, error := vm.GetVirtualMachineByName()
	fmt.Println(res[0].Id)
	c.Assert(error, check.IsNil)
}
