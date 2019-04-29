package main

/*
#include <stdio.h>

int print()
{
	printf("hello\n");
	return 111;
}

struct Stu
{
	int age;
	int high;
};

int test(int i)
{
	printf("hello %d\n",i);
	return 110;
}

void test2(struct Stu s)
{
	printf("hello %d\n",s.age);
}
*/
import "C"

import (
	"fmt"
)

type Stu struct {
	age  int
	high int
}

func main() {
	n := C.print()
	fmt.Println("return = ", n)
	m := C.test(123)
	fmt.Println(m)
	//s := Stu{age: 1, high: 2}
	//C.test2(s) //error
}
