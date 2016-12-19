package compute

import (

	"encoding/xml"
	"errors"
	"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/template"
	"github.com/megamsys/opennebula-go/virtualmachine"
)

var (
	ErrNoVM = errors.New("no vm found, Did you launch them ?")
)
const (
	TEMPLATE_INSTANTIATE = "one.template.instantiate"
	ONE_VM_ACTION        = "one.vm.action"
	ONE_DISK_SNAPSHOT    = "one.vm.disksaveas"
	ONE_IMAGE_REMOVE     = "one.image.delete"
	DELETE               = "terminate"
	REBOOT               = "reboot"
	POWEROFF             = "poweroff"
	RESUME               = "resume"
	ASSEMBLY_ID          = "assembly_id"
	ASSEMBLIES_ID        = "assemblies_id"
	ACCOUNTS_ID          = "accounts_id"
	ORG_ID               = "org_id"
	API_KEY              = "api_key"
	PLATFORM_ID          = "platform_id"
	SSH_PUBLIC_KEY       = "SSH_PUBLIC_KEY"
)

type VirtualMachine struct {
	Name         string
	TemplateName string
	TemplateId   int
	Image        string
	ContextMap   map[string]string
	Cpu          string
	CpuCost      string
	VCpu         string
	Memory       string
	MemoryCost   string
	HDD          string
	HDDCost      string
	Region       string
	ClusterId    string
	VMId         int
	Vnets        map[string]string
	T            *api.Rpc
}

type Image struct {
	Name         string
	VMId         int
	ImageId      int
	Region       string
	T            *api.Rpc
}


// Creates a new VirtualMachine
func (v *VirtualMachine) Create() (interface{}, error) {
	templateObj := template.TemplateReqs{TemplateName: v.TemplateName, T: v.T}
	defer v.T.Client.Close()

	XMLtemplate, err := templateObj.Get()
	if err != nil {
		return nil,err
	}

	XMLtemplate[0].Template.Cpu = v.Cpu
	XMLtemplate[0].Template.VCpu = v.VCpu
	XMLtemplate[0].Template.Memory = v.Memory
	XMLtemplate[0].Template.Disk.Image = v.Image
	XMLtemplate[0].Template.Disk.Size = v.HDD
	// XMLtemplate[0].Template.Cpu_cost = v.CpuCost
	// XMLtemplate[0].Template.Memory_cost = v.MemoryCost
	// XMLtemplate[0].Template.Disk_cost = v.HDDCost
	XMLtemplate[0].Template.Context.Accounts_id = v.ContextMap[ACCOUNTS_ID]
	XMLtemplate[0].Template.Context.Platform_id = v.ContextMap[PLATFORM_ID]
	XMLtemplate[0].Template.Context.Assembly_id = v.ContextMap[ASSEMBLY_ID]
	XMLtemplate[0].Template.Context.Assemblies_id = v.ContextMap[ASSEMBLIES_ID]
	XMLtemplate[0].Template.Context.Org_id = v.ContextMap[ORG_ID]
	XMLtemplate[0].Template.Context.SSH_Public_key = v.ContextMap[SSH_PUBLIC_KEY]
	XMLtemplate[0].Template.Sched_requirments = "CLUSTER_ID=\"" + v.ClusterId + "\""

	if len(v.Vnets) > 0 {
		XMLtemplate[0].Template.Nic = XMLtemplate[0].Template.Nic[:0]
		for _, v := range v.Vnets {
			net := &template.NIC{Network: v, Network_uname: "oneadmin"}
			XMLtemplate[0].Template.Nic = append(XMLtemplate[0].Template.Nic, net)
		}
	}

	finalXML := template.UserTemplates{}
	finalXML.UserTemplate = XMLtemplate
	finalData, _ := xml.Marshal(finalXML.UserTemplate[0].Template)
	data := string(finalData)
	args := []interface{}{v.T.Key, finalXML.UserTemplate[0].Id, v.Name, false, data}
	res, err := v.T.Call(TEMPLATE_INSTANTIATE, args)
	if err != nil {
		return nil,err
	}

	return res, nil
}


/**
*
* REBoot a new virtualMachine
*
**/
func (v *VirtualMachine) Reboot() (interface{}, error) {

	args := []interface{}{v.T.Key, REBOOT, v.VMId}
	res, err := v.T.Call(ONE_VM_ACTION, args)
	if err != nil {
		return nil,err
	}
	//close connection
	defer v.T.Client.Close()

	return res, nil

}

/**
*
* POWEROFF a new virtualMachine
*
**/
func (v *VirtualMachine) Poweroff() (interface{}, error) {

	args := []interface{}{v.T.Key, POWEROFF, v.VMId}
	res, err := v.T.Call(ONE_VM_ACTION, args)
	if err != nil {
		return nil,err
	}
	defer v.T.Client.Close()

	return res, nil

}

/**
*
* Resume a new virtualMachine
*
**/
func (v *VirtualMachine) Resume() (interface{}, error) {

	args := []interface{}{v.T.Key, RESUME, v.VMId}
	res, err := v.T.Call(ONE_VM_ACTION, args)
	//close connection
	defer v.T.Client.Close()
	if err != nil {
		return nil,err
	}

	return res, nil

}

/**
 *
 * Deletes a new virtualMachine
 *
 **/
func (v *VirtualMachine) Delete() (interface{}, error) {

	args := []interface{}{v.T.Key, DELETE, v.VMId}
	res, err := v.T.Call(ONE_VM_ACTION, args)
	//close connection
	defer v.T.Client.Close()
	if err != nil {
		return nil,err
	}

	return res, nil

}

/**
 *
 * VM save as a new Image (DISK_SNAPSHOT)
 *
 **/


func (v *Image) DiskSnap() (interface{}, error) {
	args := []interface{}{v.T.Key, v.VMId, 0, v.Name, "", -1}
	res, err := v.T.Call(ONE_DISK_SNAPSHOT, args)
	//close connection
	defer v.T.Client.Close()
	if err != nil {
		return nil,err
	}
	return res, nil
}

func (v *Image) RemoveImage() (interface{}, error) {
	args := []interface{}{v.T.Key, v.ImageId}
	res, err := v.T.Call(ONE_IMAGE_REMOVE, args)
	//close connection
	defer v.T.Client.Close()
	if err != nil {
		return nil,err
	}

	return res,nil
}

func listByName(name string, client *api.Rpc) (*virtualmachine.UserVM, error) {
	vms := virtualmachine.Query{VMName: name, T: client}

	svm, err := vms.GetByName()
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
