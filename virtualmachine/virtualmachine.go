package virtualmachine

import (
	"encoding/xml"
	"log"

	"github.com/megamsys/opennebula-go/api"
)

type VirtualMachineReqs struct {
	VMName string
	Client *api.Rpc
}

type UserVMs struct {
	UserVM []*UserVM `xml:"VM"`
}

type UserVM struct {
	Id   int    `xml:"ID"`
	Uid  int    `xml:"UID"`
	Name string `xml:"NAME"`
}

/**
 *
 * Given a name, this function will return the VM
 *
 **/
func (VM *VirtualMachineReqs) GetVirtualMachineByName() ([]*UserVM, error) {

	args := []interface{}{VM.Client.Key, -2, -1, -1, -1}
	VMPool, cerr := VM.Client.Call(VM.Client.RPCClient, "one.vmpool.info", args)
	if cerr != nil {
		log.Fatal(cerr)
	}

	xmlVM := UserVMs{}
	assert, _ := VMPool[1].(string)
	_ = xml.Unmarshal([]byte(assert), &xmlVM)

	var matchedVM = make([]*UserVM, len(xmlVM.UserVM))

	for _, v := range xmlVM.UserVM {
		if v.Name == VM.VMName {
			matchedVM[0] = v
		}
	}

	return matchedVM, nil

}
