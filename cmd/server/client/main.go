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
		log.Fatal("нет аргументов, нужно 2 оргумента!")
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
		res, err := c.Fibo(context.Background(), &api.FiboReq{X: int64(x), Y: int64(y)})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetResult())
}
