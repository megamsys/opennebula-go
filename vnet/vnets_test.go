package vnet

import (

	"testing"
	"github.com/megamsys/opennebula-go/api"
  // "fmt"
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
	cm[api.PASSWORD] = "dyovAupAuck9"
	s.cm = cm
}

// func (s *S) TestGetVNets(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
// 	vm := VNETemplate{T: client}
// 	_, err := vm.VnetInfos(2)
//   err = nil
// 	c.Assert(err, check.NotNil)
// }
// func (s *S) TestListVNets(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
// 	vm := VNETemplate{T: client}
// 	_, err := vm.VnetsInfos(-1)
//   err = nil
// 	c.Assert(err, check.NotNil)
// }
// func (s *S) TestVnetCreate(c *check.C) {
// 	cl, _ := api.NewClient(s.cm)
//   temp := Vnet{}
//   ar := &Address{
//       Type: "IP4",
//       Size: "1",
//       StartIP: "192.168.1.128",
//     }
//   temp.Addrs = append(temp.Addrs,ar)
//   t := Vnet{
//     Name: "vnet1",
//     Type: "fixed",
//     Description: "vnet for iPV4 ",
//     Bridge: "one",
//     Network_addr: "10.0.0.0",
//     Network_mask: "255.255.255.0",
//     Dns: "10.0.0.1",
//     Gateway: "10.0.0.1",
//     Vn_mad: "dummy",
//     Addrs: temp.Addrs,
//   }
// 	v := VNETemplate{T: cl, Template: t}
//
// 	c.Assert(v, check.NotNil)
//
// 	_, err := v.CreateVnet(-1)
// 	err = nil
// 	c.Assert(err, check.NotNil)
// }

// func (s *S) TestVnetAddIp(c *check.C) {
// 	cl, _ := api.NewClient(s.cm)
//   temp := Vnet{}
//   ar := &Address{
//       Type: "IP4",
//       Size: "1",
//       StartIP: "192.168.1.104",
//     }
//   var i int = 0
//   temp.Addrs = append(temp.Addrs,ar)
//   t := Vnet{
//     Id:  i,
//     Addrs: temp.Addrs,
//   }
//   v := VNETemplate{T: cl, Template: t}
//
//   c.Assert(v, check.NotNil)
//   res, err := v.VnetAddIps()
//   fmt.Println(res)
//   c.Assert(err, check.IsNil)
// }
