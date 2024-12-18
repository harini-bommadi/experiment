package main

import "fmt"

func arr() {
	arr := make([]int, 5)
	arr[0] = 1
	arr[1] = 10
	fmt.Println(arr)
}

type Emp struct {
	EmpID     int
	firstname string
	lastname  string
	Age       int
}

func struct1() {
	var emp1 Emp
	emp1.EmpID = 1
	emp1.firstname = "harini"
	emp1.lastname = "bommadi"
	emp1.Age = 908
}

func main() {
	fmt.Println("Hello World")
	arr()
	struct1()
}
