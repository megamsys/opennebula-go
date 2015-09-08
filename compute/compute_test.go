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
	client, _ := api.NewRPCClient("http://192.168.1.100:2633/RPC2", "oneadmin", "yib4OquafUp1")
	vmObj := VirtualMachine{Name: "yeshapp_new_files", TemplateName: "yesh_trusty", Cpu: "1", Memory: "1024", Assembly_id: "ASM007", Client: &client} //memory in terms of MB! duh!

	_, error := vmObj.Create()
	c.Assert(error, check.IsNil)
}

/*

func (s *S) TestDelete(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	vmObj := VirtualMachine{Name: "yeshapp", Client: &client}

	_, error := vmObj.Delete()
	c.Assert(error, check.IsNil)
}
*/
