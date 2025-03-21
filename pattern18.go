package main

import "fmt"

func pattern18(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf("%c", 'A'+n-1-j)
		}
		fmt.Println()
	}
}

func main() {
	pattern18(5)
}
