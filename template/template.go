package template

import (
	"encoding/xml"
  "fmt"
	"github.com/megamsys/opennebula-go/api"
)

const (
	TEMPLATEPOOL_INFO = "one.templatepool.info"
	TEMPLATE_UPDATE   = "one.template.update"
)

type UserTemplates struct {
	UserTemplate []*UserTemplate `xml:"VMTEMPLATE"`
}

type UserTemplate struct {
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
	Context                  *Context  `xml:"CONTEXT"`
	Cpu                      string    `xml:"CPU"`
	Cpu_cost                 string    `xml:"CPU_COST"`
	Description              string    `xml:"DESCRIPTION"`
	Hypervisor               string    `xml:"HYPERVISOR"`
	Logo                     string    `xml:"LOGO"`
	Memory                   string    `xml:"MEMORY"`
	Memory_cost              string    `xml:"MEMORY_COST"`
	Sunstone_capacity_select string    `xml:"SUNSTONE_CAPACITY_SELECT"`
	Sunstone_Network_select  string    `xml:"SUNSTONE_NETWORK_SELECT"`
	VCpu                     string    `xml:"VCPU"`
	Graphics                 *Graphics `xml:"GRAPHICS"`
	Disk                     *Disk     `xml:"DISK"`
	From_app                 string    `xml:"FROM_APP"`
	From_app_name            string    `xml:"FROM_APP_NAME"`
	Nic                      *NIC      `xml:"NIC"`
	Os                       *OS       `xml:"OS"`
}

type Graphics struct {
	Listen       string `xml:"LISTEN"`
	Port         string `xml:"PORT"`
	Type         string `xml:"TYPE"`
	RandomPassWD string `xml:"RANDOM_PASSWD"`
}

type OS struct {
	Arch string `xml:"ARCH"`
}

type NIC struct {
	Network       string `xml:"NETWORK"`
	Network_uname string `xml:"NETWORK_UNAME"`
}

type Context struct {
	Network        string `xml:"NETWORK"`
	Files          string `xml:"FILES"`
	SSH_Public_key string `xml:"SSH_PUBLIC_KEY"`
	Set_Hostname   string `xml:"SET_HOSTNAME"`
	Node_Name      string `xml:"NODE_NAME"`
	Assembly_id    string `xml:"ASSEMBLY_ID"`
}

type Disk struct {
	Driver      string `xml:"DRIVER"`
	Image       string `xml:"IMAGE"`
	Image_Uname string `xml:"IMAGE_UNAME"`
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

type TemplateReqs struct {
	TemplateName string
	TemplateId   int
	TemplateData string
	Client       *api.Rpc
}

/**
 *
 * Given a templateId it would return the XML of that particular template
 *
 **/
func (t *TemplateReqs) GetTemplate() ([]interface{}, error) {
	args := []interface{}{t.Client.Key, -2, t.TemplateId, t.TemplateId}
	res, err := t.Client.Call(t.Client.RPCClient, TEMPLATEPOOL_INFO, args)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/**
 *
 * Gets a particular template with a template name given
 *
 **/
func (t *TemplateReqs) GetTemplateByName() ([]*UserTemplate, error) {
	args := []interface{}{t.Client.Key, -2, -1, -1}
	templatePool, err := t.Client.Call(t.Client.RPCClient, TEMPLATEPOOL_INFO, args)

	if err != nil {
		return nil, fmt.Errorf("unable to get templates by name:%s", err)
	}

	xmlStrt := UserTemplates{}

	assert := templatePool[1].(string)

	if err = xml.Unmarshal([]byte(assert), &xmlStrt); err != nil {
		return nil, fmt.Errorf("unable to unmarshall the xml templates:%s", err)
	}

	var matchedTemplate = make([]*UserTemplate, len(xmlStrt.UserTemplate))

	for _, v := range xmlStrt.UserTemplate {
		if v.Name == t.TemplateName {
			matchedTemplate[0] = v
		}
	}
	return matchedTemplate, nil
}

/**
 *
 * Update a template in OpenNebula
 *
 **/
func (t *TemplateReqs) UpdateTemplate() error {
	args := []interface{}{t.Client.Key, t.TemplateId, t.TemplateData, 0}
	if _, err := t.Client.Call(t.Client.RPCClient, TEMPLATE_UPDATE, args); err != nil {
		return fmt.Errorf("unable to update template:%s", err)
	}
	return nil
}
