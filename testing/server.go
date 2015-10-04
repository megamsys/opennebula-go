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

type One int

type template struct {
}

func (t *One) template(args int) *template {
	fmt.Println(args)
	return nil
}

// Stop stops the server.
func (s *OneServer) Stop() {
	if s.listener != nil {
		s.listener.Close()
	}
}

// URL returns the HTTP URL of the server.
func (s *OneServer) URL() string {
	if s.listener == nil {
		return ""
	}
	return "http://" + s.listener.Addr().String() + "/"
}

func NewServer(host string) (*OneServer, error) {
	on := new(One)
	rpc.Register(on)
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
