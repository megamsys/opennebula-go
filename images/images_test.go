package images
/*
import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
	"testing"
	"fmt"
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
 	cm[api.ENDPOINT] = "http://192.168.0.118:2633/RPC2"
 	cm[api.USERID] = "oneadmin"
 	cm[api.PASSWORD] = "WuedmopFupt6"
 	s.cm = cm
 }

 func (s *S) TestImageShow(c *check.C) {
 	cl, _ := api.NewClient(s.cm)

 	v := &Image{T: cl, Id: 94}

 	c.Assert(v, check.NotNil)
	res, err := v.ImageShow()
  fmt.Println("Image State: ",res.State_string())
 	c.Assert(err, check.IsNil)
 }

 func (s *S) TestImageList(c *check.C) {
 	cl, _ := api.NewClient(s.cm)

 	v := &Image{T: cl}

 	c.Assert(v, check.NotNil)
	res, err := v.ImageList()
  fmt.Println(res)
 	c.Assert(err, check.IsNil)
 }
 */
