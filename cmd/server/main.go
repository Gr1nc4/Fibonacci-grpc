package main

import (
	"google.golang.org/grpc"
	"grpcfibo/pkg/api"
	"grpcfibo/pkg/fibo"
	"log"
	"net"
)

func main(){
	//Launch grpc server
	s := grpc.NewServer()
	srv := &fibo.GRPCServer{}
	api.RegisterFiboNumServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil{
		log.Fatal(err)
	}
}