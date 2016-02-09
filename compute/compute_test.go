package compute

import (
	"testing"

	"github.com/megamsys/opennebula-go/api"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct {
	cm map[string]string
}

var _ = check.Suite(&S{})

func (s *S) SetUpSuite(c *check.C) {
	cm := make(map[string]string)
	cm[api.ENDPOINT] = "http://103.56.92.4:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "yourWuOtHij3"
	s.cm = cm
}

func (s *S) TestCreate(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "test", TemplateName: "megam", Cpu: "1", Memory: "1024", Image: "megam", T: cl, ContextMap: map[string]string{"assembly_id": "ASM-007", "assemblies_id": "AMS-007"}} //memory in terms of MB! duh!
	c.Assert(v, check.NotNil)
	_, err := v.Create()
	c.Assert(err, check.IsNil)
}

func (s *S) TestReboot(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "test", T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.Reboot()
	c.Assert(err, check.IsNil)
}

func (s *S) TestResume(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "test", T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.Resume()
	c.Assert(err, check.IsNil)
}

func (s *S) TestPoweroff(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	vmObj := VirtualMachine{Name: "test", T: cl}
	c.Assert(vmObj, check.NotNil)
	_, err := vmObj.Poweroff()
	c.Assert(err, check.IsNil)
}

func (s *S) TestDelete(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "test", T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.Delete()
	c.Assert(err, check.IsNil)
}
