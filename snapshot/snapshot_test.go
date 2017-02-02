package snapshot

import (

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
  cm[api.ENDPOINT] = "http://192.168.0.116:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "5suvJafOtper"
	s.cm = cm
}

// func (s *S) TestCreateSnapshot(c *check.C) {
// 	cl, _ := api.NewClient(s.cm)
// 	v := Snapshot{
// 	 VMId: 333,
//   DiskId:         0,
// 		DiskDiscription: "backy_test",
//   T:     cl,
// }

// 	c.Assert(v, check.NotNil)
// 	res, err := v.CreateSnapshot()
// 	fmt.Println(res)
// 	err = nil
// 	c.Assert(err, check.NotNil)
// }


// func (s *S) TestDeleteSnapshot(c *check.C) {
// 	cl, _ := api.NewClient(s.cm)
// 	v := Snapshot{
// 	 VMId: 333,
//   DiskId:         0,
// 		SnapId: 0,
//   T:     cl,
// }
// 	c.Assert(v, check.NotNil)
// 	res, err := v.DeleteSnapshot()
// 	fmt.Println(res)
// 	err = nil
// 	c.Assert(err, check.NotNil)
// }




func (s *S) TestRevertSnapshot(c *check.C) {
	cl, _ := api.NewClient(s.cm)
	v := Snapshot{
	 VMId: 333,
  DiskId:         0,
		SnapId: 0,
  T:     cl,
}
	c.Assert(v, check.NotNil)
	res, err := v.RevertSnapshot()
	fmt.Println(res)
	err = nil
	c.Assert(err, check.NotNil)
}
