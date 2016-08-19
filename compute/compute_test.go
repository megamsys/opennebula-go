package compute

import (
	//"fmt"
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
	cm[api.ENDPOINT] = "http://192.168.0.117:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "dyovAupAuck9"
	s.cm = cm
}

/*
func (s *S) TestCreate(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	fmt.Println(cl)
	v := VirtualMachine {
		Name: "testmegam4",
		TemplateName: "megam",
		Cpu: "1",
		Memory: "1024",
		Image: "megam",
		ClusterId: "101" ,
		T: cl,
		ContextMap: map[string]string{"assembly_id": "ASM-007", "assemblies_id": "AMS-007", ACCOUNTS_ID: "info@megam.io"},
		Vnets:map[string]string{"0":"pub2_ipv4"},
		} //memory in terms of MB! duh!

	c.Assert(v, check.NotNil)
  res, err := v.Create()
	fmt.Println(res)
	c.Assert(err, check.NotNil)
}
/*

func (s *S) TestReboot(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "testrj", T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.Reboot()
	c.Assert(err, check.NotNil)
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

func (s *S) TestPoweroffKVM(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	vmObj := VirtualMachine{Name: "kvm106", T: cl}
	c.Assert(vmObj, check.NotNil)
	_, err := vmObj.Resume()
	c.Assert(err, check.IsNil)
}

func (s *S) TestDelete(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "testmegam", T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.Delete()
	c.Assert(err, check.IsNil)
}

func (s *S) TestDiskSnap(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := VirtualMachine{Name: "rj",T: cl}
	c.Assert(v, check.NotNil)
	_, err := v.DiskSnap()
	c.Assert(err, check.IsNil)
}

*/
