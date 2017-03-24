package virtualmachine

import (
	"fmt"
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
)

// func (s *S) TestGetByName(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
// 	vm := Query{VMName: "kvm109", T: client}
// 	_, err := vm.GetByName()
// 		fmt.Println(err)
// 	c.Assert(err, check.NotNil)
// }

func (s *S) TestGet(c *check.C) {
	client, _ := api.NewClient(s.cm)
	vm := Vnc{T: client, VmId: "877"}
	b, err := vm.GetVm()
	c.Assert(err, check.IsNil)
}

//
// func (s *S) TestGetVMTemplate(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
// 	vm := Vnc{VmId: "177", T: client}
// 	_, err := vm.GetVm()
// 	c.Assert(err, check.IsNil)
// }
//
// func (s *S) TestAttachNic(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
//   vm := Vnc{VmId: "425", T: client}
//   err := vm.AttachNic("ipv4-pri")
//   res, err := vm.GetVm()
//   fmt.Println("Error :", err, "\n", res.VmTemplate)
// 	c.Assert(nil, check.NotNil)
// }

// func (s *S) TestDetachNic(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
//   vm := Vnc{VmId: "425", T: client}
// 	nic := 3
//   err := vm.DetachNic(nic)
//   fmt.Println("Error :", err, "\n")
// 	c.Assert(nil, check.NotNil)
// }
