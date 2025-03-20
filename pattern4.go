package main

import "fmt"

func pattern4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func main() {
	pattern4(5)
}
