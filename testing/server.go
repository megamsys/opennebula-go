package testing

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type OneServer struct{}
type one int

type template struct {
}

func (t *one) template(args int) *template {
	fmt.Println(args)
	return nil
}

func NewServer() (*OneServer, error) {
	one := new(one)
	rpc.Register(one)
	rpc.HandleHTTP()

	/*server := OneServer{
		listener:       listener,
		imgIDs:         make(map[string]string),
		failures:       make(map[string]string),
	}
	*/
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil, nil
}
