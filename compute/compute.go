package compute

import (
	"encoding/xml"
	"errors"

	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/template"
	"github.com/megamsys/opennebula-go/virtualmachine"
)

var ErrNoVM = errors.New("no vm found, Did you launch them ?")

const (
	TEMPLATE_INSTANTIATE = "one.template.instantiate"
	ONE_VM_ACTION        = "one.vm.action"
	DELETE               = "delete"
	RESUME               = "resume"
	REBOOT               = "reboot"
	POWEROFF             = "poweroff"

	ASSEMBLY_ID   = "assembly_id"
	ASSEMBLIES_ID = "assemblies_id"
	ACCOUNTS_ID   = "accounts_id"
	PLATFORM_ID   = "platform_id"
)

type VirtualMachine struct {
	Name         string
	TemplateName string
	TemplateId   int
	Image        string
	ContextMap   map[string]string
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
func (v *VirtualMachine) Create() ([]interface{}, error) {
	templateObj := template.TemplateReqs{TemplateName: v.TemplateName, Client: v.Client}

	/*
	 * get a particular template to configure it
	 */
	XMLtemplate, err := templateObj.GetTemplateByName()
	if err != nil {
		return nil, err
	}

	/*
	 * Assign Values
	 */

	XMLtemplate[0].Template.Cpu = v.Cpu
	XMLtemplate[0].Template.VCpu = v.VCpu
	XMLtemplate[0].Template.Memory = v.Memory
	XMLtemplate[0].Template.Disk.Image = v.Image
	XMLtemplate[0].Template.Context.Accounts_id = v.ContextMap[ACCOUNTS_ID]
	XMLtemplate[0].Template.Context.Platform_id = v.ContextMap[PLATFORM_ID]
	XMLtemplate[0].Template.Context.Assembly_id = v.ContextMap[ASSEMBLY_ID]
	XMLtemplate[0].Template.Context.Assemblies_id = v.ContextMap[ASSEMBLIES_ID]

	finalXML := template.UserTemplates{}
	finalXML.UserTemplate = XMLtemplate

	finalData, _ := xml.Marshal(finalXML.UserTemplate[0].Template)
	data := string(finalData)

	/*
	 * Instantiate a template
	 */
	args := []interface{}{v.Client.Key, finalXML.UserTemplate[0].Id, v.Name, false, data}
	res, err := v.Client.Call(v.Client.RPCClient, TEMPLATE_INSTANTIATE, args)
	//close connection
	defer v.Client.RPCClient.Close()

	if err != nil {
		return nil, err
	}
	return res, nil
}

/**
 *
 * Deletes a new virtualMachine
 *
 **/
func (v *VirtualMachine) Delete() ([]interface{}, error) {
	uvm, err := listByName(v.Name, v.Client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{v.Client.Key, DELETE, uvm.Id}
	res, err := v.Client.Call(v.Client.RPCClient, ONE_VM_ACTION, args)
	//close connection
	defer v.Client.RPCClient.Close()
	if err != nil {
		return nil, err
	}

	return res, nil

}

/**
*
* Resume a new virtualMachine
*
**/
func (v *VirtualMachine) Resume() ([]interface{}, error) {
	uvm, err := listByName(v.Name, v.Client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{v.Client.Key, RESUME, uvm.Id}
	res, err := v.Client.Call(v.Client.RPCClient, ONE_VM_ACTION, args)
	//close connection
	defer v.Client.RPCClient.Close()
	if err != nil {
		return nil, err
	}

	return res, nil

}

/**
*
* REBoot a new virtualMachine
*
**/
func (v *VirtualMachine) Reboot() ([]interface{}, error) {
	uvm, err := listByName(v.Name, v.Client)
	if err != nil {
		return nil, err
	}

	args := []interface{}{v.Client.Key, REBOOT, uvm.Id}
	res, err := v.Client.Call(v.Client.RPCClient, ONE_VM_ACTION, args)
	//close connection
	defer v.Client.RPCClient.Close()
	if err != nil {
		return nil, err
	}

	return res, nil

}

/**
*
* POWEROFF a new virtualMachine
*
**/
func (v *VirtualMachine) Poweroff() ([]interface{}, error) {
	uvm, err := listByName(v.Name, v.Client)
	if err != nil {
		return nil, err
	}
	args := []interface{}{v.Client.Key, POWEROFF, uvm.Id}
	res, err := v.Client.Call(v.Client.RPCClient, ONE_VM_ACTION, args)
	defer v.Client.RPCClient.Close()
	if err != nil {
		return nil, err
	}

	return res, nil

}

func listByName(name string, client *api.Rpc) (*virtualmachine.UserVM, error) {
	vms := virtualmachine.VirtualMachineReqs{VMName: name, Client: client}

	svm, err := vms.GetVirtualMachineByName()
	if err != nil {
		return nil, err
	}

	if len(svm) <= 0 {
		return nil, ErrNoVM
	}

	if svm[0] == nil {
		return nil, ErrNoVM
	}
	return svm[0], nil
}
