package flavor

import (
	"fmt"

	"github.com/megamsys/opennebula-go/api"
)

type FlavorOpts struct {
	TemplateName string
	TemplateId   int
}

func (flavor *FlavorOpts) GetTemplate() {
	//templatepool.info, INFO_ALL, -1, -1

	client, err := api.RPCClient("http://localhost:2633/RPC2")
	if err != nil {
		fmt.Println(err)
	}
	args := []interface{}{key, -2, flavor.TemplateId, flavor.TemplateId}
	res := api.Call(client, "one.templatepool.info", args)
	if res != nil {
		fmt.Println(res)
	}

}
