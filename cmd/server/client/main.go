package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"grpcfibo/pkg/api"
	"log"
	"strconv"
)

func main(){
	//Add args launch
	flag.Parse()
	if flag.NArg() < 2{
		log.Fatal("Нужно 2 оргумента!")
	}
	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y,err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}
	//Set connection
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewFiboNumClient(conn)
	//Get fibonacci sequence
	for i := x; i <= y; i++ {
		res, err := c.Fibo(context.Background(), &api.FiboReq{X: int32(i)})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetResult())
	}
}
