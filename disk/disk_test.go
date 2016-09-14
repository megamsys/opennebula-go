package disk
import (
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
	"testing"
	//"fmt"
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

 func (s *S) TestDatastoreAllocate(c *check.C) {
 	cl, _ := api.NewClient(s.cm)
	d := Disk{
//Disk_Type:    "FS",
 //Dev_Prefix:   "vd",
  Size:         "1024",
  //Target:       "vdc",
}
	t := Vm{Disk: d}
 	v := VmDisk{T: cl, VmId: 0, Vm: t}

 	c.Assert(v, check.NotNil)
	_, err := v.AttachDisk()
 	err = nil
 	c.Assert(err, check.IsNil)
 }
