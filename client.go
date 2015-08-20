package api

import (
	"fmt"
	//"bytes"
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

func Call(RPC *xmlrpc.Client, command string, args []interface{}) error {
	//result := make([]byte, 10)
	result := []interface{}{}
	res := RPC.Call(command, args, &result)
	fmt.Println(&result)
	fmt.Println(res)
	if res != nil {
		return res
	}
	return nil
}
