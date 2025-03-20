package main

import "fmt"

func pattern3(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {
	pattern3(5)
}
