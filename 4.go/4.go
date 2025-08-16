package main

import "fmt"

func pattern(n int) {
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Printf("%c", 'A'+i)
		}
		fmt.Println("")
	}
}
func main() {
	pattern(26)
}
