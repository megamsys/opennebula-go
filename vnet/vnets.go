package vnet

import (
  "github.com/megamsys/opennebula-go/api"
  "encoding/xml"
  "fmt"
)

type VNETemplate struct {
  Template  Vnet `xml:"TEMPLATE"`
  T   *api.Rpc
}

type Vnet struct {
  Id             int    `xml:"ID"`
  Name           string `xml:"NAME"`
  Type           string  `xml:"TYPE"`
  Description    string `xml:"DESCRIPTION"`
  Bridge         string `xml:"BRIDGE"`
  Network_addr   string `xml:"NETWORK_ADDRESS"`
  Network_mask   string `xml:"NETWORK_MASK"`
  Dns            string `xml:"DNS"`
  Gateway        string `xml:"GATEWAY"`
  Vn_mad         string  `xml:"VN_MAD"`
  Addrs          []*Address  `xml:"AR"`
}

type Address struct {
  Type        string `xml:"TYPE"`
  StartIP     string `xml:"IP"`
  Size        string `xml:"SIZE"`
}

func (v *VNETemplate) CreateVnet(id int) ([]interface{} ,error)  {

  finalXML := VNETemplate{}
	finalXML.Template = v.Template
	finalData, _ := xml.Marshal(finalXML.Template)
	data := string(finalData)
	args := []interface{}{v.T.Key,data, id}
	res, err := v.T.Call(api.VNET_CREATE, args)
	if err != nil {
		return nil, err
	}
    return res, nil
}

func (v *VNETemplate) VnetAddIps() ([]interface{} ,error)  {
  finalXML := VNETemplate{}
	finalXML.Template.Addrs = v.Template.Addrs
	finalData, _ := xml.Marshal(finalXML.Template)
	data := string(finalData)
	args := []interface{}{v.T.Key,data, v.Template.Id}
	res, err := v.T.Call(api.VNET_ADDIP, args)
	if err != nil {
		return nil, err
	}
  return res, nil
}

func (v *VNETemplate) VnetInfos(id int) ([]interface{}, error) {
	args := []interface{}{v.T.Key, id}
	res, err := v.T.Call(api.VNET_SHOW, args)
	if err != nil {
		return nil, err
	}
  return res, nil
}


func (v *VNETemplate) VnetsInfos(id int) ([]interface{}, error) {
	args := []interface{}{v.T.Key, id, -1, -1}
	res, err := v.T.Call(api.VNET_LIST, args)
	if err != nil {
		return nil, err
	}
  return res, nil
}
