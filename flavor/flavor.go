package flavor

import (
	"fmt"

	"github.com/megamsys/opennebula-go/api"
)

type FlavorOpts struct {
	TemplateName string
	TemplateId   int
}

/*
 * Given a templateId it would return the XML of that particular template
 *
 */

func (flavor *FlavorOpts) GetTemplate(endpoint string, key string) ([]interface{}, error) {

	client, err := api.RPCClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	args := []interface{}{key, -2, flavor.TemplateId, flavor.TemplateId}
	fmt.Println("final call ------------------------>>")
	res, err := api.Call(client, "one.templatepool.info", args)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, nil
}

func (flavor *FlavorOpts) GetTemplateByName(endpoint string, key string) ([]interface{}, error) {

	client, err := api.RPCClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}

	args := []interface{}{key, -2, -1, -1}
	templatePool, err := api.Call(client, "one.templatepool.info", args)

	//	for _, v := range templatePool {
	//		fmt.Println(v)
	//	}

	//	fmt.Println(res)
	return templatePool, nil
}
