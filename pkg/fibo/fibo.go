package fibo

import (
	"context"
	"grpcfibo/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Fibo(ctx context.Context, req *api.FiboReq)(*api.FiboResp, error){
	return &api.FiboResp{Result: Fib(req.GetX())}, nil
}
var (
	cache map[int32]int32
)
func init(){
	cache = make(map[int32]int32)
}
//Calculate fibonacci number
func Fib(n int32) int32{
	if n <= 1{
		return n
	}
	if r, ok := cache[n];ok{
		return r
	}
	sum := Fib(n-1) + Fib(n-2)
	cache[n] = sum
	return sum
}