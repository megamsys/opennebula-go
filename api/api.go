package api

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/kolo/xmlrpc"
	"github.com/megamsys/libgo/cmd"
)

const (
	ENDPOINT = "endpoint"
	USERID   = "userid"
	TEMPLATE = "template"
	PASSWORD = "password"
	IMAGE    = "image"
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
	log.Debugf(cmd.Colorfy("  > [one-go] connecting", "blue", "", "bold"))

	RPCclient, err := xmlrpc.NewClient(endpoint, nil)

	if err != nil {
		//TO-DO: trap and send connRefused error.
		return nil, err
	}
	log.Debugf(cmd.Colorfy("  > [one-go] connected", "blue", "", "bold")+" %s", endpoint)

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
	log.Debugf(cmd.Colorfy("  > [one-go] ", "blue", "", "bold")+"%s", command)
	//log.Debugf(cmd.Colorfy("\n> args   ", "cyan", "", "bold")+" %v\n", args)

	result := []interface{}{}
	err := RPC.Call(command, args, &result)

	if err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("\n> response ", "cyan", "", "bold")+" %v", result)
	log.Debugf(cmd.Colorfy("  > [one-go] ( ´ ▽ ` ) SUCCESS", "blue", "", "bold"))
	return result, nil
}
