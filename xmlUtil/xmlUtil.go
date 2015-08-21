
package xmlUtil

import (
	"encoding/xml"

)
type VMTEMPLATE_POOL struct {
	//VmTemplatePool xml.Name    `xml:"VMTEMPLATE_POOL"`
	VmTemplate     []*VMTemplate `xml:"VMTEMPLATE"`
}

type VMTemplate struct {
	Id          int       `xml:"ID"`
	Uid         int       `xml:"UID"`
	Gid         int       `xml:"GID"`
	Uname       string       `xml:"UNAME"`
	Gname       string       `xml:"GNAME"`
	Name        string       `xml:"NAME"`
	Permissions *Permissions `xml:"PERMISSIONS"`
	Template    *Template    `xml:"TEMPLATE"`
	RegTime     int       `xml:"REGTIME"`
}

type Template struct {
	Vcpu   string `xml:"VCPU"`
	Cpu    string `xml:"CPU"`
	Memory string `xml:"MEMORY"`
	Disk   *Disk `xml:"DISK"`
}

type Disk struct {
	Size int  `xml:"SIZE"`
	Type string   `xml:"TYPE"`
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

func UnmarshallXml(xmlData interface{}) VMTEMPLATE_POOL {

	xmlStrt := VMTEMPLATE_POOL{}
	assert := xmlData.(string)
	_ = xml.Unmarshal([]byte(assert), &xmlStrt)
	return xmlStrt
}
