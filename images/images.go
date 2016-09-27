package images

import (
  "encoding/xml"
  "github.com/megamsys/opennebula-go/api"
)

const (
  LOCKED = 4
  READY = 1
  USED = 2
  FAILURE = 5

)

type Images struct {
  Images []Image `xml:"IMAGE"`
  T    *api.Rpc
}

type Image struct {
Id   int    `xml:"ID"`
Uid  int    `xml:"UID"`
Name string `xml:"NAME"`
Type string `xml:"TYPE"`
Size int    `xml:"SIZE"`
State int   `xml:"STATE"`
DatastoreID int `xml:"DATASTORE_ID"`
Datastore string `xml:"DATASTORE"`
FsType   string  `xml:"FSTYPE"`
RunningVMs Vms   `xml:"VMS"`
T    *api.Rpc
}

type Vms struct {
  Id  []int `xml:"ID"`
}

func (v *Image) ImageShow() (*Image, error) {
	args := []interface{}{v.T.Key, v.Id}
	res, err := v.T.Call(api.ONE_IMAGE_SHOW, args)
	if err != nil {
		return nil,err
	}
  xmlImage := &Image{}
  if err = xml.Unmarshal([]byte(res), xmlImage); err != nil {
    return nil, err
  }
	return xmlImage, err
}

func (v *Image) ImageList() (*Images, error) {
  first := -1 // -1 for default smaller ID
  last  := -1 //-1 for default last ID
	args := []interface{}{v.T.Key, -1,first,last}
	res, err := v.T.Call(api.ONE_IMAGE_LIST, args)
	defer v.T.Client.Close()
	if err != nil {
		return nil,err
	}
  xmlImages := &Images{}
  if err = xml.Unmarshal([]byte(res), xmlImages); err != nil {
    return nil, err
  }
	return xmlImages, err
}

func (v *Image) State_string() string {
  switch v.State {
  case LOCKED:
    return "locked"
  case READY:
    return "ready"
  case USED:
    return "used"
  case FAILURE:
    return "failure"
  default:
    return "unknown"
  }
}
