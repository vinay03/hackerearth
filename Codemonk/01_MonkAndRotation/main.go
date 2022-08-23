package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var T, N, K, a int64
	fmt.Scanf("%d", &T)

	var text string
	reader := bufio.NewReader(os.Stdin)
	var A []string

	for a = 0; a < T; a++ {
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		parts := strings.Split(text, " ")
		N, _ = strconv.ParseInt(parts[0], 10, 64)
		K, _ = strconv.ParseInt(parts[1], 10, 64)

		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		A = strings.Split(text, " ")
		if K > N {
			K = K % N
		}

		fmt.Print(strings.Trim(strings.Join(A[N-K:], " ")+" "+strings.Join(A[:N-K], " "), " "), "\n")
	}
}
