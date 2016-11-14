package users

import (
	"testing"
	"github.com/megamsys/opennebula-go/api"
	"gopkg.in/check.v1"
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
  cm[api.ENDPOINT] = "http://213.32.56.135:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "Mumdidacnat1"
	s.cm = cm
}

func (s *S) TestGetUsers(c *check.C) {
  fmt.Println("**********************user test*************************")
	client, _ := api.NewClient(s.cm)
  u := User{
    UserName: "vijaym@megam.io",
    Password: "team4megam",
    AuthDriver: "core",
    GroupIds: []int{0},
  }
	vm := UserTemplate{
    T: client,
    Users: u,
   }
	res, err := vm.CreateUsers()
  fmt.Println("*****************res*************",res)
  fmt.Println(err)
  err = nil
	c.Assert(err, check.NotNil)
}
