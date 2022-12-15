package main

import "fmt"

func main() {
	queue := make([]int, 0)
	fmt.Println(queue == nil)

	queue = append(queue, 1)
	fmt.Println(queue)

	queue = queue[1:]
	fmt.Println(queue)
	fmt.Println(queue == nil)
	queue = nil
	fmt.Println(queue == nil)
}
