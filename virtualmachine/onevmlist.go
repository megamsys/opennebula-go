package virtualmachine

import (
	"encoding/xml"
	"github.com/megamsys/opennebula-go/api"
	"strconv"
	"strings"
)

const (
	//VmState starts at 0
	INIT VmState = iota
	PENDING
	HOLD
	ACTIVE
	STOPPED
	SUSPENDED
	DONE
	UNKNOWNSTATE
	POWEROFF
	UNDEPLOYED
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
	UserTemplate   UserTemplate    `xml:"USER_TEMPLATE"`
	HistoryRecords *HistoryRecords `xml:"HISTORY_RECORDS"`
}

type VmTemplate struct {
	Graphics *Graphics `xml:"GRAPHICS"`
	Context  *Context  `xml:"CONTEXT"`
	Nics      []Nic     `xml:"NIC"`
}

type Nic struct {
  IPaddress string `xml:"IP"`
	Mac       string `xml:"MAC"`
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

type UserTemplate struct {
	Description        string `xml:"DESCRIPTION"`
	Error              string `xml:"ERROR"`
	Sched_Requirements string `xml:"SCHED_REQUIREMENTS"`
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

func (v *VM) StateString() string {
	return VmStateString[VmState(v.State)]
}

func (v *VM) Nics() []Nic {
	return v.VmTemplate.Nics
}

func (v *VM) LcmStateString() string {
	return LcmStateString[LcmState(v.LcmState)]
}

func (v *VM) IsFailure() bool {
  return strings.Contains(v.LcmStateString(), "failure")
}

func (v *VM) IsSnapshotReady() bool {
	return (v.State == int(ACTIVE) && v.LcmState == int(RUNNING)) || (v.State == int(POWEROFF) && v.LcmState == int(LCM_INIT))
}
