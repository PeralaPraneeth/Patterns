package main

import "fmt"

func pattern12(n int) {
	spaces := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}

		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()

		spaces -= 2
	}
}

func main() {
	pattern12(5)
}
