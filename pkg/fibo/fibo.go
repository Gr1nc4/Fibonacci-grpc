package fibo

import (
	"context"
	"grpcfibo/pkg/api"
)

type GRPCServer struct {}

func (s *GRPCServer) Fibo(ctx context.Context, req *api.FiboReq)(*api.FiboResp, error){
	return &api.FiboResp{Result: FiboSlice( req.GetX(), req.GetY())}, nil
}
var (
	cache map[int64]int64
)
func init(){
	cache = make(map[int64]int64)
}
//Calculate fibonacci number
func Fib(n int64) int64{
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
func FiboSlice(x,y int64) []int64 {
	var arr []int64
	for i := x; i <y ; i++ {
		arr = append(arr, Fib(i))
	}
	return arr
}
