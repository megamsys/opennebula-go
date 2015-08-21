package compute

import (
	"encoding/xml"
	"reflect"
	"fmt"
"github.com/megamsys/opennebula-go/api"
	"github.com/megamsys/opennebula-go/flavor"
"github.com/megamsys/opennebula-go/xmlUtil"
)

type VirtualMachine struct {
	OpenNebulaTemplateName   string
	OpenNebulaTemplateId   int
	Bootstrap            string
	SSHUser              string
	SSHPort              string
	Cpu                  string
	VCpu                 string
	Memory               string
	RunList              string
	Distro               string
	VMName               string
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

	template[0].Template.Cpu = string("![CDATA[" + VM.Cpu + "]]")
	template[0].Template.Vcpu = string("![CDATA[" + VM.VCpu + "]]")
	template[0].Template.Memory = string("![CDATA[" + VM.Memory + "]]")

  finalXML := xmlUtil.VMTEMPLATE_POOL{}
  finalXML.VmTemplate = template


  finalData, _ := xml.Marshal(finalXML)
  data1 := string(finalData)
  fmt.Println(data1)
  fmt.Println(reflect.TypeOf(finalData))

/*
 * Updating templates
 */
//data := `<VMTEMPLATE><ID>6</ID><UID>0</UID><GID>3</GID><UNAME>osradmin</UNAME><GNAME>NewUpdate</GNAME><NAME>supertest</NAME><PERMISSIONS><OWNER_U>1</OWNER_U><OWNER_M>1</OWNER_M><OWNER_A>0</OWNER_A><GROUP_U>0</GROUP_U><GROUP_M>0</GROUP_M><GROUP_A>0</GROUP_A><OTHER_U>0</OTHER_U><OTHER_M>0</OTHER_M><OTHER_A>0</OTHER_A></PERMISSIONS><REGTIME>1440059694</REGTIME><TEMPLATE><CONTEXT><NETWORK><![CDATA[YES]]></NETWORK><SSH_PUBLIC_KEY><![CDATA[$USER[SSH_PUBLIC_KEY]]]></SSH_PUBLIC_KEY></CONTEXT><CPU><![CDATA[1]]></CPU><CPU_COST><![CDATA[10]]></CPU_COST><DESCRIPTION><![CDATA[testtemplate...]]></DESCRIPTION><HYPERVISOR><![CDATA[kvm]]></HYPERVISOR><LOGO><![CDATA[images/logos/ubuntu.png]]></LOGO><MEMORY><![CDATA[512]]></MEMORY><MEMORY_COST><![CDATA[10]]></MEMORY_COST><SUNSTONE_CAPACITY_SELECT><![CDATA[YES]]></SUNSTONE_CAPACITY_SELECT><SUNSTONE_NETWORK_SELECT><![CDATA[YES]]></SUNSTONE_NETWORK_SELECT></TEMPLATE></VMTEMPLATE>`


flavUpd := flavor.FlavorOpts{TemplateId: finalXML.VmTemplate[0].Id, TemplateData: data1}
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
