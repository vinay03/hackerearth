/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
	"fmt"
)

func main() {
	var T, N, count int
	var M [][]int

	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		fmt.Scanf("%d", &N)
		M = make([][]int, N)

		for j := 0; j < N; j++ {
			M[j] = make([]int, N)
			for k := 0; k < N; k++ {
				fmt.Scanf("%d", &M[j][k])
			}
		}

		count = 0

		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				for p := j; p < N; p++ {
					for q := k; q < N; q++ {
						if M[j][k] > M[p][q] {
							count++
						}
					}
				}
			}
		}
		fmt.Println(count)
	}
}
