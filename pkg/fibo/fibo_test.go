package fibo

import (
	"testing"
)

func TestFib(t *testing.T){
	var n int32 = 10
	var out int32 = 55

	result := Fib(n)
	if result != out{
		t.Error("Incorrect result", result,  " not equal ", out)
	}
}
