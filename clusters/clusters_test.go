package clusters

import (
  "encoding/xml"
	"testing"
	"github.com/megamsys/opennebula-go/api"
  "fmt"
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
	cm[api.ENDPOINT] = "http://88.198.139.81:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "yeghorbAjif4"
	s.cm = cm
}


func (s *S) TestCluster(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := Clusters{T: cl} //memory in terms of MB! duh!

	c.Assert(v, check.NotNil)
	res, err := v.ClusterPoolinfo()
  xmlVM := &Clusters{}
  fmt.Println(res)
  assert := res[1].(string)
  if err = xml.Unmarshal([]byte(assert), xmlVM); err != nil {
     fmt.Println(err)
  }
  fmt.Printf("%#v",xmlVM.Cluster[0])

	c.Assert(err, check.IsNil)
}
