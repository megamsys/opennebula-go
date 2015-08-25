package xmlUtil

import (
	"encoding/xml"
)

/*
This has to go into a file called templates.go
Rename this to UserTemplates (if this contains all the templates of an user)
*/
type VMTEMPLATE_POOL struct {
	//VmTemplatePool xml.Name    `xml:"VMTEMPLATE_POOL"`
	VmTemplate []*VMTemplate `xml:"VMTEMPLATE"`
}

//Rename this file to UserTemplate  (if this contains a template information of an user.
type VMTemplate struct {
	Id          int          `xml:"ID"`
	Uid         int          `xml:"UID"`
	Gid         int          `xml:"GID"`
	Uname       string       `xml:"UNAME"`
	Gname       string       `xml:"GNAME"`
	Name        string       `xml:"NAME"`
	Permissions *Permissions `xml:"PERMISSIONS"`
	Template    *Template    `xml:"TEMPLATE"`
	RegTime     int          `xml:"REGTIME"`
}

type Template struct {
	Context                  *Context `xml:"CONTEXT"`
	Cpu                      string   `xml:"CPU"`
	Cpu_cost                 string   `xml:"CPU_COST"`
	Description              string   `xml:"DESCRIPTION"`
	Hypervisor               string   `xml:"HYPERVISOR"`
	Logo                     string   `xml:"LOGO"`
	Memory                   string   `xml:"MEMORY"`
	Memory_cost              string   `xml:"MEMORY_COST"`
	Sunstone_capacity_select string   `xml:"SUNSTONE_CAPACITY_SELECT"`
	Sunstone_Network_select  string   `xml:"SUNSTONE_NETWORK_SELECT"`
	Vcpu                     string   `xml:"VCPU"`
	Disk                     *Disk    `xml:"DISK"`
}

type Context struct {
	Network        string `xml:"NETWORK"`
	SSH_Public_key string `xml:"SSH_PUBLIC_KEY"`
}

type Disk struct {
	Size int    `xml:"SIZE"`
	Type string `xml:"TYPE"`
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

/*Remove this and move to a common interface called XMLMapper
- Marshall()
- UnMarshall()
Every API shall implement this interface. This interface shall reside in a file in the core/ directory.
*/
func UnmarshallXml(xmlData interface{}) VMTEMPLATE_POOL {

	xmlStrt := VMTEMPLATE_POOL{}
	assert := xmlData.(string)
	_ = xml.Unmarshal([]byte(assert), &xmlStrt)
	return xmlStrt
}
