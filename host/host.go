package host

import (
  "fmt"
	"github.com/megamsys/opennebula-go/api"
)

const (
  SHOW = "show"
)
type HQuery struct {
	HostID  string
	T      *api.Rpc
}

type Host struct {
  Id  int         `xml:ID`
  HostName string     `xml:NAME`
	Temp    *Template `xml:"TEMPLATE"`
}

type Template struct {
  VMs []*VM        `xml:"VM"`
  Wildvms string `xml:"WILDS"`
}

type VM struct {
	Id   int          `xml:"ID"`
	Deploy_id  int    `xml:"DEPLOY_ID"`
	VMName string     `xml:"VM_NAME"`
}

// Given a name, this function will return the VM
func (v *HQuery) HostInfos(a int) ([]interface{}, error) {
  args := []interface{}{v.T.Key,a}
	fmt.Println("**********************88")
	hostinfo, err := v.T.Call(api.ONE_HOST_INFO, args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return hostinfo, nil

}
