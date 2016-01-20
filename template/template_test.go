
package template

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

func (s *S) TestGetTemplateByName(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	flav := TemplateReqs{TemplateName: "newone", Client: client}
	res, error := flav.GetTemplateByName()
	c.Assert(error, check.IsNil)
	c.Assert(res, check.NotNil)
}
