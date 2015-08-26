package api

import (
	"fmt"
	"log"

	"github.com/kolo/xmlrpc"
)

/*
 * RPC Client and secret key
 */
type Rpc struct {
	RPCClient xmlrpc.Client
	Key       string
}

/*
* Creates an RPCClient with endpoint and returns it
*
 */
func NewRPCClient(endpoint string, username string, password string) (Rpc, error) {
	RPCclient, err := xmlrpc.NewClient(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	RpcObj := Rpc{RPCClient: *RPCclient, Key: username + ":" + password}
	fmt.Println(RpcObj.Key)
	return RpcObj, nil
}

/*
 * Do an RPC Call
 *
 */
func (c *Rpc) Call(RPC xmlrpc.Client, command string, args []interface{}) ([]interface{}, error) {

	result := []interface{}{}
	cerr := RPC.Call(command, args, &result)
	if cerr != nil {
		return nil, cerr
	}

	return result, nil
}
