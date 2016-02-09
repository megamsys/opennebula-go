package virtualmachine

import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
	"testing"
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


func (s *S) TestGetByName(c *check.C) {
	client, _ := api.NewClient(s.cm)
	vm := Query{VMName: "yeshapp", T: client}
	_, err := vm.GetByName()
	c.Assert(err, check.IsNil)
}
