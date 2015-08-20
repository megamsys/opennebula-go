package compute

import (
	"github.com/megamsys/opennebula-go/api"
   "github.com/megamsys/opennebula-go/flavor"
)

type VirtualMachine struct {
	OpenNebulaTemplate string
	Bootstrap          string
	SSHUser            string
	SSHPort            string
	Cpu                string
	VCpu               string
	Memory             string
	RunList            string
	Distro             string
	VMName             string
}



}

type VirtualMachineInternal struct {
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

	//1. pull template details - get_by_name - rpc

flavorObj := flavor.FlavorOpts{TemplateName: vm.OpenNebulaTemplate}
finalFlavor := flavorObj.GetFlavorByName(secretKey)

   //2. get with id to configure additions
	//

	client, err := api.RPCClient(creds.Endpoint)
	if err != nil {
		fmt.Println(err)
	}
	res := api.Call(client, "one.vm.allocate")
   if res != nil {
		fmt.Println(res)
	}
}

func (vm *VirtualMachine) DestroyVM() {}
