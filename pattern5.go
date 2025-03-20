package main

import "fmt"

func pattern5(n int) {
	for i := 1; i <= n; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	pattern5(5)
}
