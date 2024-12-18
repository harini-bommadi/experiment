package main

import "fmt"

func arr() {
	arr := make([]int, 5)
	arr[0] = 1
	arr[1] = 10
	fmt.Println(arr)
}

func main() {
	fmt.Println("Hello World")
	arr()
	arr()
}
