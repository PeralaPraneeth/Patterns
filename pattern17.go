package main

import "fmt"

func pattern17(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		for j := i - 1; j >= 0; j-- {
			fmt.Printf("%c", 'A'+j)
		}

		fmt.Println()
	}
}

func main() {
	pattern17(5)
}
