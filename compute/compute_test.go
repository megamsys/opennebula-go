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
	//oneadmin:yib4OquafUp1
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vmObj := VirtualMachine{Name: "test", TemplateName: "sass", Cpu: "1", Memory: "1024", Image: "img-ubuntu", Client: client, ContextMap: map[string]string{"assembly_id": "ASM-007", "assemblies_id": "AMS-007"}} //memory in terms of MB! duh!

	_, error := vmObj.Create()
	c.Assert(error, check.IsNil)
}

/*

func (s *S) TestDelete(c *check.C) {
	client, _ := api.NewRPCClient("http://loca	lhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vmObj := VirtualMachine{Name: "yeshapp", Client: &client}

	_, error := vmObj.Delete()
	c.Assert(error, check.IsNil)
}
*/
