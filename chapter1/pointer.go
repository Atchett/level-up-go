package main

import (
	"fmt"
)

type counter struct {
	Count int
}

func (c counter) increment() {
	c.Count++
}

func (c *counter) incrementWithPointer() {
	c.Count++
}

func main() {

	counter := &counter{}
	fmt.Println("Initial count:", counter.Count)

	counter.increment()
	fmt.Println("Count after increment:", counter.Count)

	counter.incrementWithPointer()
	fmt.Println("Count after pointer increment:", counter.Count)

}
