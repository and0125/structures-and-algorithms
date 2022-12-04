package main

import (
	"fmt"
)

func main() {
	listIntegers := []int64{3, 4, 510, 45, 60, 23, 303}
	listNames := []string{"Andrew", "Jake", "Russell", "Aaron", "Darius"}

	fmt.Println(binary_search(listIntegers, 5))
	fmt.Println(binary_search(listIntegers, 303))
	fmt.Println(binary_search(listNames, "Terry"))
	fmt.Println(binary_search(listNames, "Aaron"))

}

func binary_search(inputList interface{}, searchItem interface{}) [2]int64 {
	/* In Go, in order to handle multiple data types in a single function, you need to do switch
	statements based on the type, then define the function*/
	typedList, ok := inputList.([]int64)
	if ok {
		/* code for binary search for a int64 list */
		return [2]int64{0,0} 
	} else if typedList, ok := inputList.([]string); ok{
		/* code for binary search for a string list */
		return [2]int64{0,0}
	} else {
		return [2]int64{404,404} 
	}
}