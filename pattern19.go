package main

import "fmt"

func pattern19(n int) {
	space := 0
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print("*")
		}
		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}
		for j := n; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
		space += 2
	}

	spaces := 2*n - 2
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		spaces -= 2
	}

}

func main() {
	pattern19(5)
}
