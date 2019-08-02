package main

import (
	"fmt"
	"math/rand"

	"./binary_tree"
)

func main() {
	headNode := new(binary_tree.Node)
	mappy := make(map[int]int)
	isDone := make(chan bool)
	go func() {

		for i := 0; i < 2500000; i++ {
			headNode.Insert(rand.Intn(1000000))
		}
		//time.Sleep((10 * time.Second))
		isDone <- true
	}()

	<-isDone
	headNode.InOrderTraversalHash(mappy)
	fmt.Println("Frequencies:")

	for key, value := range mappy {
		fmt.Printf("%d: %d\n", key, value)
	}
}
