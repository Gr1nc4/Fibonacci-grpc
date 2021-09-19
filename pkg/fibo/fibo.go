package fibo

import (
	"context"
	"grpcfibo/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Fibo(ctx context.Context, req *api.FiboReq)(*api.FiboResp, error){
	return &api.FiboResp{Result: Fib(req.GetX())}, nil
}

func Fib(n int32) int32{
	if n <= 1{
		return n
	}
	sum := Fib(n-1) + Fib(n-2)
	return sum
}