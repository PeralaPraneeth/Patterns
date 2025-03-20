package main

import "fmt"

func pattern15(n int) {
	for i := n; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", 'A'+j)
		}
		fmt.Println()
	}
}

func main() {
	pattern15(4)
}
