package main

import (
	"fmt"
)

func main() {
	var N, Sum, i int64
	var A []int64

	fmt.Scanf("%d", &N)
	A = make([]int64, N)

	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
		Sum += A[i]
	}

	fmt.Print(Sum)
}
