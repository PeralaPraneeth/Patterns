package main

import "fmt"

func pattern13(n int) {
	sum := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(sum)
			sum = sum + 1
		}
		fmt.Println()
	}

}

func main() {
	pattern13(5)
}
