package virtualmachine

import (
	"encoding/xml"

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
func (v *VirtualMachineReqs) GetVirtualMachineByName() ([]*UserVM, error) {

	args := []interface{}{v.Client.Key, -2, -1, -1, -1}
	VMPool, err := v.Client.Call(v.Client.RPCClient, "one.vmpool.info", args)
	if err != nil {
		return nil, err
	}

	xmlVM := UserVMs{}
	assert, _ := VMPool[1].(string)
	if err = xml.Unmarshal([]byte(assert), &xmlVM); err != nil {
		return nil, err
	}

	var matchedVM = make([]*UserVM, len(xmlVM.UserVM))

	for _, u := range xmlVM.UserVM {
		if u.Name == v.VMName {
			matchedVM[0] = u
		}
	}

	return matchedVM, nil

}
