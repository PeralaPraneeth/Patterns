package main

import "fmt"

func pattern6(n int) {
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {
	pattern6(5)
}
