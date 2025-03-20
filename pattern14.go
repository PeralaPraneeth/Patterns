package main

import "fmt"

func pattern14(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func main() {
	pattern14(5)
}
