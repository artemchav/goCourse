package main

import (
	"fmt"
	"structs"
)

func main() {

	queue := &structs.Queue{}
	queue.Put(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(queue.Len())

	fmt.Println(queue.GetQueueContent())

	fmt.Println(queue.TakeOneTask())
	fmt.Println(queue.Len())

	fmt.Println(queue.GetQueueContent())

	queue.Put(321, 123)
	fmt.Println(queue.Len())

	fmt.Println(queue.TakeOneTask())
	fmt.Println(queue.GetQueueContent())
	fmt.Println(queue.Len())

}
