package compute

import (
	"github.com/megamsys/opennebula-go/api"

)

type VirtualMachine struct {
	id           string
	template_str string
	name         string
	uuid         string
	state        string
	status       string
	ip           string
	mac          string
	vcpu         string
	cpu          string
	memory       string
	user         string
	gid          string
	group        string
	onevm_object string
	flavor       string
}

type Creds struct {
	Username string
	Password string
	Endpoint string
}

func (vm *VirtualMachine) CreateVM(creds *Creds) {

	secretKey := creds.Username + ":" + creds.Password
	client := api.RPCClient(creds.Endpoint)
	api.Call(client, "one.vm.allocate")

}

func (vm *VirtualMachine) DestroyVM() {}
