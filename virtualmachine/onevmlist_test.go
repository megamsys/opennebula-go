package virtualmachine

import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
	"fmt"
)



func (s *S) TestGetVMTemplate(c *check.C) {
	client, _ := api.NewClient(s.cm)
	vm := Vnc{VmId: "119", T: client}
	_, err := vm.GetVm()
	c.Assert(err, check.NotNil)
}
