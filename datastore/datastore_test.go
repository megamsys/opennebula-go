package datastore

import (
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
	cm[api.PASSWORD] = "asdf"
	s.cm = cm
}

func (s *S) TestDatastoreAllocate(c *check.C) {
	cl, _ := api.NewClient(s.cm)
  t := Datastore{
    Name: "cephds",
    Ds_mad: "ceph",
    Tm_mad: "ceph",
    Disk_type: "rbd",
    Bridge_list: "192.168.1.103",
    Ceph_host: "megam",
    Pool_name: "one",
    Ceph_user: "libvirt",
    Ceph_secret: "3d74a4a1-e6fc-4485-a6d6-3ddadfad",
  }
	v := DatastoreTemplate{T: cl, Template: t}

	c.Assert(v, check.NotNil)

}
