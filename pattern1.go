package main

import "fmt"

func patterns(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	patterns(5)
}
