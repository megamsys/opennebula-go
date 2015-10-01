package api

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/kolo/xmlrpc"
	"github.com/megamsys/libgo/cmd"
)

const (
	ENDPOINT       = "endpoint"
	USERID         = "userid"
	TEMPLATE       = "template"
	PASSWORD       = "password"
	IMAGE          = "image"
)

var ErrConnRefused = errors.New("connection refused")


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
	log.Debugf(cmd.Colorfy("\n> [one-go]", "white", "", "bold") + cmd.Colorfy(" client", "green", "", ""))

	RPCclient, err := xmlrpc.NewClient(endpoint, nil)

	if err != nil {
		//TO-DO: trap and send connRefused error.
		return nil, err
	}
	log.Debugf(cmd.Colorfy("\n> connected", "purple", "", "bold")+" %s\n", endpoint)

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
	log.Debugf(cmd.Colorfy("  > request", "blue", "", "bold")+" %s", command)
	//log.Debugf(cmd.Colorfy("\n> args   ", "cyan", "", "bold")+" %v\n", args)

	result := []interface{}{}
	err := RPC.Call(command, args, &result)
	if err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("\n> response ", "cyan", "", "bold")+" %v", result)
	log.Debugf(cmd.Colorfy("  > request SUCCESS", "blue", "", "bold")+" %s", command)
	return result, nil
}
