package compute

import (
	"encoding/xml"
	"log"

	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/template"
	"github.com/megamsys/opennebula-go/virtualmachine"
)

const (
	TEMPLATE_INSTANTIATE = "one.template.instantiate"
	ONE_VM_ACTION        = "one.vm.action"
	DELETE               = "delete"
)

type VirtualMachine struct {
	Name         string
	TemplateName string
	TemplateId   int
	Assembly_id  string
	Cpu          string
	VCpu         string
	Memory       string
	Client       *api.Rpc
}

/**
 *
 * Creates a new VirtualMachine
 *
 **/
func (VM *VirtualMachine) Create() ([]interface{}, error) {

	templateObj := template.TemplateReqs{TemplateName: VM.TemplateName, Client: VM.Client}

	/*
	 * get a particular template to configure it
	 */
	XMLtemplate, ferr := templateObj.GetTemplateByName()
	if ferr != nil {
		log.Fatal(ferr)
	}

	/*
	 * Assign Values
	 */

	XMLtemplate[0].Template.Cpu = VM.Cpu
	XMLtemplate[0].Template.VCpu = VM.VCpu
	XMLtemplate[0].Template.Memory = VM.Memory

	XMLtemplate[0].Template.Context.Assembly_id = VM.Assembly_id

	finalXML := template.UserTemplates{}
	finalXML.UserTemplate = XMLtemplate

	finalData, _ := xml.Marshal(finalXML.UserTemplate[0].Template)
	data := string(finalData)

	/*
	 * Instantiate a template
	 */
	args := []interface{}{VM.Client.Key, finalXML.UserTemplate[0].Id, VM.Name, false, data}
	res, cerr := VM.Client.Call(VM.Client.RPCClient, TEMPLATE_INSTANTIATE, args)
	if cerr != nil {
		log.Fatal(cerr)
	}
	return res, nil
}

/**
 *
 * Deletes a new virtualMachine
 *
 **/
func (VM *VirtualMachine) Delete() ([]interface{}, error) {

	vmObj := virtualmachine.VirtualMachineReqs{VMName: VM.Name, Client: VM.Client}

	SingleVM, ferr := vmObj.GetVirtualMachineByName()
	if ferr != nil {
		log.Fatal(ferr)
	}

	args := []interface{}{VM.Client.Key, DELETE, SingleVM[0].Id}
	res, cerr := VM.Client.Call(VM.Client.RPCClient, ONE_VM_ACTION, args)
	if cerr != nil {
		log.Fatal(cerr)
		return nil, cerr
	}

	return res, nil

}
