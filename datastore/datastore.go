package datastore

import (
	"encoding/xml"
	"fmt"
	"github.com/megamsys/opennebula-go/api"
)

type DatastoreTemplate struct {
	Template Datastore `xml:"DATASTORES"`
	T        *api.Rpc
}

type Datastore struct {
	Id          int    `xml:"ID"`
	Name        string `xml:"NAME"`
	Ds_mad      string `xml:"DS_MAD"`
	Tm_mad      string `xml:"TM_MAD"`
	Disk_type   string `xml:"DISK_TYPE"`
	Bridge_list string `xml:"BRIDGE_LIST"`
	Ceph_host   string `xml:"CEPH_HOST"`
	Type        string `xml:"TYPE"`
	Safe_dirs   string `xml:"SAFE_DIRS"`
	Pool_name   string `xml:"Pool_NAME"`
	Ceph_user   string `xml:"CEPH_USER"`
	Ceph_secret string `xml:"CEPH_SECRET"`
  Host        string `xml:"HOST"`
  Vg_name     string `xml:"VG_NAME"`
}

func (v *DatastoreTemplate) AllocateDatastore(id int) ([]interface{}, error) {
	finalXML := DatastoreTemplate{}
	finalXML.Template = v.Template
	finalData, _ := xml.Marshal(finalXML.Template)
	data := string(finalData)
	args := []interface{}{v.T.Key, data, id}
	res, err := v.T.Call(api.ONE_DATASTORE_ALLOCATE, args)
	if err != nil {
		return nil, err
	}
  fmt.Println(res)
	return res, nil
}


func (v *DatastoreTemplate) GetDATAs(a int) ([]interface{}, error) {
	args := []interface{}{v.T.Key, a}
	Datastores, err := v.T.Call(api.ONE_DATASTORE_INFO, args)
	if err != nil {
		return nil, err
	}

	return Datastores, nil

}
