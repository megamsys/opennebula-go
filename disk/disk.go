package disk

import (
	"encoding/xml"
	"github.com/megamsys/opennebula-go/api"
)

type VmDisk struct {
	VmId int `xml:"ID"`
	Vm   Vm  `xml:"VM"`
	T    *api.Rpc
}

type Vm struct {
	Disk Disk `xml:"DISK"`
}
type Disk struct {
	Disk_Type  string `xml:"TYPE"`
	Dev_Prefix string `xml:"DEV_PREFIX"`
	Size       string `xml:"SIZE"`
	Target     string `xml:"TARGET"`
}

func (v *VmDisk) AttachDisk() ([]interface{}, error) {
	if v.Vm.Disk.Dev_Prefix == "" {
		v.Vm.Disk.Dev_Prefix = "vd"
	}
	if v.Vm.Disk.Disk_Type == "" {
		v.Vm.Disk.Disk_Type = "FS"
	}
	finalXML := VmDisk{}
	finalXML.Vm = v.Vm
	finalData, _ := xml.Marshal(finalXML.Vm)
	data := string(finalData)
	args := []interface{}{v.T.Key, v.VmId, data}
	res, err := v.T.Call(api.DISK_CREATE, args)
	defer v.T.Client.Close()
	if err != nil {
		return nil, err
	}
	return res, err

}
