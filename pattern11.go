package main

import "fmt"

func pattern11(n int) {
	start := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}
		for j := 0; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}

func main() {
	pattern11(5)
}
