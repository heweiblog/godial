package main

import "fmt"

type Peo struct {
	name string
	age  int
}

func main() {
	m := make(map[string]Peo, 5)
	p := Peo{
		name: "hww",
		age:  18,
	}
	m[p.name] = p
	fmt.Println(m["hww"])
}
