package virtualmachine

import (
"encoding/xml"
   "fmt"
	"github.com/megamsys/opennebula-go/api"
)

type Apiclient struct {
   VmId  int
T      *api.Rpc
VM     *VM `xml:"VM"`
}


type VM struct {

	Id   int    `xml:"ID"`
  Name string `xml:"NAME"`
	VmTemplate *VmTemplate `xml:"TEMPLATE"`
	HistoryRecords *HistoryRecords  `xml:"HISTORY_RECORDS"`
}

type VmTemplate struct {
Graphics   *Graphics `xml:"GRAPHICS"`
}

type HistoryRecords struct {
	 History *History  `xml:"HISTORY"`
}
type History struct {
	HostName string `xml:"HOSTNAME"`
}

type Graphics struct {
		Port string `xml:"PORT"`
}


func (v *Apiclient) GetVm() (*VM, error) {

	args := []interface{}{v.T.Key, v.VmId}
	onevm, err := v.T.Call(api.VM_INFO, args)
	fmt.Println("********************************")
	fmt.Println(onevm)

	if err != nil {
		return nil, err
	}

  	xmlVM := &VM{}
  	assert, _ := onevm[1].(string)
  	if err = xml.Unmarshal([]byte(assert), xmlVM); err != nil {
  		return nil, err
  	}


fmt.Println("**************ip")

fmt.Println(*xmlVM)
fmt.Println(xmlVM.GetPort())
fmt.Println(xmlVM.GetHostIp())
return xmlVM, err
}


func (u *VM) GetPort() string {

	fmt.Println("*********port")
	fmt.Println(u.VmTemplate.Graphics.Port)
	return u.VmTemplate.Graphics.Port
}

func (u *VM) GetHostIp() string {
	fmt.Println("*********ip")
	fmt.Println(u.HistoryRecords.History.HostName)
	return u.HistoryRecords.History.HostName
}
