package compute

import (
	"encoding/xml"
	"fmt"

	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/flavor"
	"github.com/megamsys/opennebula-go/xmlUtil"
)

type VirtualMachine struct {
	OpenNebulaTemplateName string
	OpenNebulaTemplateId   int
	Bootstrap              string
	SSHUser                string
	SSHPort                string
	Cpu                    string
	VCpu                   string
	Memory                 string
	RunList                string
	Distro                 string
	VMName                 string
}

type Credentials struct {
	Username string
	Password string
	Endpoint string
}

func (VM *VirtualMachine) CreateVM(creds *Credentials) {

	secretKey := creds.Username + ":" + creds.Password

	flavorObj := flavor.FlavorOpts{TemplateName: VM.OpenNebulaTemplateName}

	/*
	 * get a particular template to configure it
	 */
	template, ferr := flavorObj.GetTemplateByName(creds.Endpoint, secretKey)
	if ferr != nil {
		fmt.Println(ferr)
	}

	/*
	 * Assign Values
	 */

	template[0].Template.Cpu = VM.Cpu
	template[0].Template.Vcpu = VM.VCpu
	template[0].Template.Memory = VM.Memory

	finalXML := xmlUtil.VMTEMPLATE_POOL{}
	finalXML.VmTemplate = template

	finalData, _ := xml.Marshal(finalXML.VmTemplate[0].Template)
	data := string(finalData)

	/*
	 * Updating templates
	 */

	flavUpd := flavor.FlavorOpts{TemplateId: finalXML.VmTemplate[0].Id, TemplateData: data}
	flavUpd.UpdateTemplate(creds.Endpoint, secretKey)

	/*
	 * Allocate VM with ID
	 */

	args := []interface{}{secretKey, finalXML.VmTemplate[0].Id}
	client, err := api.RPCClient(creds.Endpoint)
	if err != nil {
		fmt.Println(err)
	}
	_, cerr := api.Call(client, "one.vm.allocate", args)
	if cerr != nil {
		fmt.Println(cerr)
	}
}
