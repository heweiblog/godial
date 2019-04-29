package main

import (
	"fmt"
	"gitee.com/johng/gf/g/container/gmap"
)

type Stu struct {
	name string
	age  int
}

func main() {
	p := gmap.NewStringInterfaceMap()
	s1 := Stu{"hww", 18}
	p.Set(s1.name, s1)
	tmp := p.Get("hww")
	v, ok := tmp.(Stu)
	if ok {
		fmt.Println(v.name, v.age)
		v.age = 33
		p.Set("mnn", v)
	} else {
		fmt.Println("no")
	}

	s1 = Stu{"sss", 22}
	p.Set(s1.name, s1)
	s1 = Stu{"hdd", 12}
	p.Set(s1.name, s1)
	var l bool

	p.Iterator(func(k string, v interface{}) bool {
		task, ok := v.(Stu)

		if ok && task.age == 9 {
			fmt.Println(task)
			l = true
			return false
		}
		fmt.Println(task)
		return true
	})

	fmt.Println(l)

	r := p.Get("ssis")
	if r == nil {
		fmt.Println("no")
	} else {
		fmt.Println("have", r)
	}

}
