package api

import (
	"fmt"

	"github.com/kolo/xmlrpc"
)

/*
* Creates an RPCClient with endpoint and returns it
*
 */
func RPCClient(endpoint string) (*xmlrpc.Client, error) {
	RPCclient, err := xmlrpc.NewClient(endpoint, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(RPCclient)
	return RPCclient, nil
}

/*
* Execute an RPC Call
*
 */

func Call(RPC *xmlrpc.Client, command string, secretKey string) {
	//Need to include args
	args := "VM:INFO"
	RPC.Call(command, secretKey, args)

}
