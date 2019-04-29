package main

import "fmt"

type Stu struct {
	name string
	age  int
}

func (s Stu) Set1(name string, age int) {
	s.name = name
	s.age = age
}

func (s *Stu) Set2(name string, age int) {
	s.name = name
	s.age = age
}

func print1(s Stu) {
	fmt.Println(s)
}

func print2(s *Stu) {
	fmt.Println(*s)
}

func main() {
	s := Stu{name: "hww", age: 18}
	ps := &s
	s.Set1("aaa", 1)
	fmt.Println(s)
	s.Set2("bbb", 2)
	fmt.Println(s)
	ps.Set1("ccc", 3)
	fmt.Println(s)
	ps.Set2("ddd", 4)
	fmt.Println(s)
	print1(s)
	print1(*ps)
	print2(&s)
	print2(ps)
}
