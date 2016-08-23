package datastore

import (
  "github.com/megamsys/opennebula-go/api"
  "encoding/xml"
  //"fmt"
)

type DatastoreTemplate struct {
  Template Datastore `xml:"DATASTORES"`
  T          *api.Rpc
}

type Datastore struct {
  Id             int     `xml:"ID"`
  Name           string  `xml:"NAME"`
  Ds_mad         string  `xml:"DS_MAD"`
  Tm_mad         string `xml:"TM_MAD"`
  Disk_Type      string `xml:"DISK_TYPE"`
  Bridge_List    string `xml:"BRIDGE_LIST"`
  Ceph_host      string `xml:"CEPH_HOST"`
  Pool_name      string `xml:"Pool_NAME"`
  Ceph_user      string `xml:"CEPH_USER"`
  Ceph_secret    string  `xml:"CEPH_SECRET"`
}
