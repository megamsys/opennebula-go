package clusters

import (
 "github.com/megamsys/opennebula-go/api"
)

const (
  GETCLUSTERS  = "one.clusterpool.info"
  GETCLUSTER   = "one.cluster.info"
)


type Clusters struct {
    Cluster   []*cluster `xml:"CLUSTER"`
  	T            *api.Rpc
}

type cluster struct {
	ID         string       `xml:"ID"`
	Name       string       `xml:"NAME"`
	Hosts      *Host        `xml:"HOSTS"`
	Datastores *Datastore   `xml:"DATASTORES"`
	Vnets      *Vnet        `xml:"VNETS"`
}

type Host struct {
  ID    string `xml:"ID"`
}

type Datastore struct {
   ID    string `xml:"ID"`
}

type Vnet struct {
   ID    string `xml:"ID"`
}

func (c *Clusters) ClusterPoolinfo() ([]interface{}, error) {
  args := []interface{}{c.T.Key}
	res, err := c.T.Call(GETCLUSTERS, args)
	//close connection
	defer c.T.Client.Close()
	if err != nil {
		return nil, err
	}

	return res, nil

}
