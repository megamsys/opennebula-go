package api

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestCreateNewRPCClient(c *check.C) {
	_, error := NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	c.Assert(error, check.IsNil)
}

func (s *S) TestRPCCall(c *check.C) {
	client, clientErr := NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	args := []interface{}{client.Key, -2, 3, 3}
	_, callErr := client.Call(client.RPCClient, "one.templatepool.info", args)
	c.Assert(clientErr, check.IsNil)
	c.Assert(callErr, check.IsNil)

}
