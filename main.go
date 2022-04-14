package main

import (
	"github.com/kjhch/go-greeting/service"
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(service.AlbumService), "")
	s.Serve("tcp", ":9091")
}
