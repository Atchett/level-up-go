package main

import (
	"fmt"
	"sort"
)

func main() {

	// map of star wars films
	starWarsYears := map[string]int{
		"A New Hope":              1977,
		"The Empire Strikes Back": 1980,
		"Return of the Jedi":      1983,
		"Attack of the Clones":    2002,
		"Revenge of the Sith":     2005,
	}

	// to be able to sort the map
	// put the ints into a slice
	var years []int
	for _, year := range starWarsYears {
		years = append(years, year)
	}

	// sort the slice
	sort.Ints(years)

	for _, year := range years {
		t, ok := getKeyUsingValue(starWarsYears, year)
		if !ok {
			fmt.Println("Year not found")
		}
		fmt.Println(t, "was released in", year)
	}
}

// get the key from the map using the value
func getKeyUsingValue(m map[string]int, val int) (key string, ok bool) {
	for k, v := range m {
		if v == val {
			key = k
			ok = true
			return
		}
	}
	return
}
