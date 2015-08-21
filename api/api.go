package api

import (
	"fmt"
	//"bytes"
	"github.com/kolo/xmlrpc"
//	"github.com/megamsys/opennebula-go/compute"
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
 * Do an RPC Call
 *
 */
func Call(RPC *xmlrpc.Client, command string, args []interface{}) ([]interface{}, error) {

	result := []interface{}{}
	cerr := RPC.Call(command, args, &result)
	if cerr != nil {
		return nil, cerr
	}

	return result, nil
}
