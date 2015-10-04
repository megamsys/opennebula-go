package testing

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type OneServer struct {
	listener net.Listener
}

type one int

type template struct {
}

func (t *one) template(args int) *template {
	fmt.Println(args)
	return nil
}

func NewServer(host string) (*OneServer, error) {
	one := new(one)
	rpc.Register(one)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", host)

	if err != nil {
		return nil, err
	}

	server := OneServer{
		listener: l,
	}
	go http.Serve(l, nil)
	return &server, nil
}
