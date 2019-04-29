package main

import "fmt"

func main() {
	if a := 3; a < 1 || a > 5 {
		fmt.Println("a<1||a>5")
	} else {
		fmt.Println("a>=1&&a<=5", a)
	}
}
