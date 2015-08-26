package compute

import (
	"testing"

	"github.com/megamsys/opennebula-go/api"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestCreate(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vmObj := VirtualMachine{Name: "yeshapp", TemplateName: "fullfledged", Cpu: "1", VCpu: "1", Memory: "4500", Client: &client} //memory in terms of MB! duh!

	_, error := vmObj.Create()
	c.Assert(error, check.IsNil)
}

func (s *S) TestDelete(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vmObj := VirtualMachine{Name: "yeshapp", Client: &client} //memory in terms of MB! duh!

	_, error := vmObj.Delete()
	c.Assert(error, check.IsNil)
}
