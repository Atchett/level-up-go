package main

import (
	"fmt"

	"bitbucket.org/johnpersonal/LevelUpGo/chapter2/constants"
)

const (
	_ int = iota //ignore the first value (0 for iota)
	one
	two
	three
	four
	five
)

// Constants for testing
const (
	KB constants.ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
	fmt.Println(five)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

}
