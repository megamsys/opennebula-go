package api

import (
	"errors"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/kolo/xmlrpc"
	"github.com/megamsys/libgo/cmd"
)

const (
	ENDPOINT = "endpoint"
	USERID   = "userid"
	PASSWORD = "password"
	TEMPLATE = "template"
	IMAGE    = "image"
	VCPU_PERCENTAGE ="vcpu_percentage"

	VMPOOL_ACCOUNTING = "one.vmpool.accounting"
	VMPOOL_INFO       = "one.vmpool.info"
	TEMPLATEPOOL_INFO = "one.templatepool.info"
	TEMPLATE_UPDATE   = "one.template.update"
	VM_INFO           = "one.vm.info"
)

var (
	ErrArgsNotSatisfied = errors.New("[" + ENDPOINT + "," + USERID + "," + PASSWORD + "] one (or) more args missing!")
)

type Rpc struct {
	Client xmlrpc.Client
	Key    string
}

func NewClient(config map[string]string) (*Rpc, error) {
	log.Debugf(cmd.Colorfy("  > [one-go] connecting", "blue", "", "bold"))

	if !satisfied(config) {
		return nil, ErrArgsNotSatisfied
	}

	client, err := xmlrpc.NewClient(config[ENDPOINT], nil)
	if err != nil {
		return nil, err
	}

	log.Debugf(cmd.Colorfy("  > [one-go] connected", "blue", "", "bold")+" %s", config[ENDPOINT])

	return &Rpc{
		Client: *client,
		Key:    config[USERID] + ":" + config[PASSWORD]}, nil
}

func (c *Rpc) Call(command string, args []interface{}) ([]interface{}, error) {
	log.Debugf(cmd.Colorfy("  > [one-go] ", "blue", "", "bold")+"%s", command)
	log.Debugf(cmd.Colorfy("\n> args   ", "cyan", "", "bold")+" %v\n", args)

	result := []interface{}{}
	if err := c.Client.Call(command, args, &result); err != nil {
		return nil, err
	}
	//log.Debugf(cmd.Colorfy("\n> response ", "cyan", "", "bold")+" %v", result)
	log.Debugf(cmd.Colorfy("  > [one-go] ( ´ ▽ ` ) SUCCESS", "blue", "", "bold"))
	return result, nil
}

func satisfied(c map[string]string) bool {
	return len(strings.TrimSpace(c[ENDPOINT])) > 0 &&
		len(strings.TrimSpace(c[USERID])) > 0 &&
		len(strings.TrimSpace(c[PASSWORD])) > 0
}
