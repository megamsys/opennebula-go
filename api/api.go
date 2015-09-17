package api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/megamsys/libgo/cmd"
	"github.com/kolo/xmlrpc"
)

/*
 * RPC Client and secret key
 */
type Rpc struct {
	RPCClient xmlrpc.Client
	Key       string
}

/**
 *
 * Creates an RPCClient with endpoint and returns it
 *
 **/
func NewRPCClient(endpoint string, username string, password string) (*Rpc, error) {
	RPCclient, err := xmlrpc.NewClient(endpoint, nil)

	if err != nil {
		return nil, fmt.Errorf("unable to connect with one %q: %s", endpoint, err)
	}
	log.Debugf(cmd.Colorfy("\n> [one-go]", "white", "", "bold") + cmd.Colorfy("client", "green", "", "") + "\n")
	log.Debugf(cmd.Colorfy("\n> connected", "purple", "", "bold") + " %s\n", endpoint)

	return &Rpc{
		RPCClient: *RPCclient,
		Key:       username + ":" + password}, nil
}

/**
 *
 * Do an RPC Call
 *
 **/
func (c *Rpc) Call(RPC xmlrpc.Client, command string, args []interface{}) ([]interface{}, error) {
	log.Debugf("[one-go] -- request --- \ncommand: %s\nargs: %s\n[one-go] --- request --- \n", command, args)
	log.Debugf(cmd.Colorfy("\n> request", "blue", "", "bold") + "%s", command)
	log.Debugf(cmd.Colorfy("\n> args   ", "cyan", "", "bold") + " %v\n", args)

	result := []interface{}{}
	err := RPC.Call(command, args, &result)
	if err != nil {
		return nil, fmt.Errorf("one rpc execution %q failure: %s", command, err)
	}
	log.Debugf(cmd.Colorfy("\n> response ", "cyan", "", "bold") + " %v\n", result)

	return result, nil
}
