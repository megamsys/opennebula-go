package virtualmachine

import (
	"encoding/xml"
	"github.com/megamsys/opennebula-go/api"
	"strconv"
)

type Vnc struct {
	VmId string
	T    *api.Rpc
	VM   *VM `xml:"VM"`
}

type VM struct {
	Id             string          `xml:"ID"`
	Name           string          `xml:"NAME"`
	State          int             `xml:"STATE"`
	LcmState       int             `xml:"LCM_STATE"`
	VmTemplate     *VmTemplate     `xml:"TEMPLATE"`
	HistoryRecords *HistoryRecords `xml:"HISTORY_RECORDS"`
}

type VmTemplate struct {
	Graphics *Graphics `xml:"GRAPHICS"`
	Context  *Context  `xml:"CONTEXT"`
}

type Context struct {
	VMIP string `xml:"ETH0_IP"`
}

type HistoryRecords struct {
	History *History `xml:"HISTORY"`
}
type History struct {
	HostName string `xml:"HOSTNAME"`
}

type Graphics struct {
	Port string `xml:"PORT"`
}

func (v *Vnc) GetVm() (*VM, error) {
	intstr, _ := strconv.Atoi(v.VmId)
	args := []interface{}{v.T.Key, intstr}
	onevm, err := v.T.Call(api.VM_INFO, args)
	if err != nil {
		return nil, err
	}

	xmlVM := &VM{}
	if err = xml.Unmarshal([]byte(onevm), xmlVM); err != nil {
		return nil, err
	}
	return xmlVM, err
}

func (u *VM) GetPort() string {
	return u.VmTemplate.Graphics.Port
}

func (u *VM) GetState() int {
	return u.State
}

func (u *VM) GetLcmState() int {
	return u.LcmState
}

func (u *VM) GetHostIp() string {
	return u.HistoryRecords.History.HostName
}

func (u *VM) GetVMIP() string {
	return u.VmTemplate.Context.VMIP
}
