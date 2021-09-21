package fibo

import (
	"testing"
)

func TestFib(t *testing.T){
	var n int64 = 10
	var out int64 = 55

	result := Fib(n)

	if result != out{
		t.Error("Incorrect result", result, "not equal", out)
	}
}

func TestFiboSlice(t *testing.T) {
	var x, y int64
	x, y = 0, 3
	var arr = FiboSlice(x, y)
	var arrTest = []int64{0, 1, 1, 2}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arrTest[i] {
			t.Error("Incorrect result", arr[i], "not equal", arrTest[i])
		}
	}
}
