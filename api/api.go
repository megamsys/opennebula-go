package api

import (
	"fmt"
	//"bytes"
	"github.com/kolo/xmlrpc"
	"github.com/megamsys/opennebula-go/compute"
)

/*
* Creates an RPCClient with endpoint and returns it
*
 */
func RPCClient(endpoint string) (*xmlrpc.Client, error) {
	fmt.Println("-createing client --------->>>>")
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

func Call(RPC *xmlrpc.Client, command string, args []interface{}) (string, error) {
	//result := make([]byte, 10)

&result := compute.XMLStructure{}

	//var result string
	err := RPC.Call(command, args, &result)
	if err != nil {
		return nil, err
	}
	//fmt.Println(&result)
	//fmt.Println(res)-
	fmt.Println("==============---------=================================--------------=========================")
fmt.Println(result.VmTemplates.Id)
	return result, nil
}
