package compute

import (
	"bytes"
	"encoding/gob"
	"encoding/xml"
	"fmt"

	"github.com/megamsys/opennebula-go/flavor"
)

type VirtualMachine struct {
	OpenNebulaTemplate   string
	OpenNebulaTemplateId int
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

type XMLStructure struct {
	VmTemplatePool xml.Name    `xml:"VMTEMPLATE_POOL"`
	VmTemplate     *VMTemplate `xml:"VMTEMPLATE"`
}

type VMTemplate struct {
	Id          string       `xml:"ID"`
	Uid         string       `xml:"UID"`
	Gid         string       `xml:"GID"`
	Uname       string       `xml:"UNAME"`
	Gname       string       `xml:"GNAME"`
	Name        string       `xml:"NAME"`
	Permissions *Permissions `xml:"PERMISSIONS"`
	Template    *Template    `xml:"TEMPLATE"`
	RegTime     string       `xml:"REGTIME"`
}

type Template struct {
	Vcpu   string `xml:"VCPU"`
	Cpu    string `xml:"CPU"`
	Memory string `xml:"MEMORY"`
}

type Permissions struct {
	Owner_U int `xml:"OWNER_U"`
	Owner_M int `xml:"OWNER_M"`
	Owner_A int `xml:"OWNER_A"`
	Group_U int `xml:"GROUP_U"`
	Group_M int `xml:"GROUP_M"`
	Group_A int `xml:"GROUP_A"`
	Other_U int `xml:"OTHER_U"`
	Other_M int `xml:"OTHER_M"`
	Other_A int `xml:"OTHER_A"`
}

type Credentials struct {
	Username string
	Password string
	Endpoint string
}

func (vm *VirtualMachine) CreateVM(creds *Credentials) {

	secretKey := creds.Username + ":" + creds.Password

	flavorObj := flavor.FlavorOpts{TemplateId: vm.OpenNebulaTemplateId}
	fmt.Println("calling----------get template")
	finalFlavorXML, ferr := flavorObj.GetTemplate(creds.Endpoint, secretKey)
	if ferr != nil {
		fmt.Println(ferr)
	}
	fmt.Println(finalFlavorXML)
	/*-------
	fmt.Println("------------------------------")
	fmt.Println(finalFlavorXML)
	data := `<VMTEMPLATE_POOL><VMTEMPLATE><ID>0</ID><UID>0</UID><GID>0</GID><UNAME>oneadmin</UNAME><GNAME>oneadmin</GNAME><NAME>ubuntu</NAME><PERMISSIONS><OWNER_U>1</OWNER_U><OWNER_M>1</OWNER_M><OWNER_A>0</OWNER_A><GROUP_U>0</GROUP_U><GROUP_M>0</GROUP_M><GROUP_A>0</GROUP_A><OTHER_U>0</OTHER_U><OTHER_M>0</OTHER_M><OTHER_A>0</OTHER_A></PERMISSIONS><REGTIME>1440059694</REGTIME><TEMPLATE><CONTEXT><NETWORK><![CDATA[YES]]></NETWORK><SSH_PUBLIC_KEY><![CDATA[$USER[SSH_PUBLIC_KEY]]]></SSH_PUBLIC_KEY></CONTEXT><CPU><![CDATA[1]]></CPU><CPU_COST><![CDATA[10]]></CPU_COST><DESCRIPTION><![CDATA[testtemplate...]]></DESCRIPTION><HYPERVISOR><![CDATA[kvm]]></HYPERVISOR><LOGO><![CDATA[images/logos/ubuntu.png]]></LOGO><MEMORY><![CDATA[512]]></MEMORY><MEMORY_COST><![CDATA[10]]></MEMORY_COST><SUNSTONE_CAPACITY_SELECT><![CDATA[YES]]></SUNSTONE_CAPACITY_SELECT><SUNSTONE_NETWORK_SELECT><![CDATA[YES]]></SUNSTONE_NETWORK_SELECT></TEMPLATE></VMTEMPLATE></VMTEMPLATE_POOL>`
	//	data := `<VMTEMPLATE_POOL><VMTEMPLATE><ID>0</ID><UID>0</UID></VMTEMPLATE></VMTEMPLATE_POOL>`
	//data := finalFlavorXML[1]
	//data, _ := xml.Marshal(finalFlavorXML[1])
	//data := GetBytes(finalFlavorXML[1])
	//datum := string(data)
	fmt.Println(reflect.TypeOf(finalFlavorXML))

	t := XMLStructure{}
	x := xml.Unmarshal([]byte(data), &t)
	fmt.Println(x)
	fmt.Println("---------------------------------------")
	fmt.Println(t.VmTemplate.Id)

	args := []interface{}{secretKey, 1}
	client, err := api.RPCClient(creds.Endpoint)
	if err != nil {
		fmt.Println(err)
	}
	_, cerr := api.Call(client, "one.vm.allocate", args)
	if cerr != nil {
		fmt.Println(cerr)
	} */
}

func GetBytes(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func (vm *VirtualMachine) DestroyVM() {}
