package main

import (
	"fmt"
	"sort"
)

func binary_search(inputList []int, searchItem int) [2]int {
	// sort list of integers
	sort.Ints(inputList)
	fmt.Println(inputList)
	
	// initialize parameters 
	low := 0
	high := len(inputList) - 1
	steps := 0
	error := -1
		/* 
		note that Go automatically turns this calculation
		into a floor operation because integers are used.
		You have to convert all numbers involved to floats
		if you want to have the remainder included.
		*/
	for low <= high {
		
		steps++
		mid := low + high / 2
		guess := inputList[mid]
		if guess == searchItem {
			return [2]int{mid, steps}
		}else if guess > searchItem {
			high--
		} else {
			low++
		}
	}
	return [2]int{error, steps}
}

func main() {
	a := []int{3, 4, 510, 45, 60, 23, 303}
	i1 := binary_search(a, 303)
	i2 := binary_search(a, 460)
	fmt.Println(i1, i2)
}	
