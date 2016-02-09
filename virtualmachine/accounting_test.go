package virtualmachine

import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestGet(c *check.C) {
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")

	vm := Accounting{Client: client, starttime: time.Now().Unix(), endtime: time.Now().Unix()}
	_, error := vm.Get()
	c.Assert(error, check.IsNil)
}
