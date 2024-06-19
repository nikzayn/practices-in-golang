package main

import (
	"fmt"
)

func main() {
	// type valStruct struct{}
	// valInt := 0
	// valString := ""
	// valBool := true

	// fmt.Println("Int", unsafe.Sizeof(valInt))
	// fmt.Println("String", unsafe.Sizeof(valString))
	// fmt.Println("Bool", unsafe.Sizeof(valBool))
	// fmt.Println("Struct", unsafe.Sizeof(valStruct{}))

	// Creating a set using a map with empty struct values
	set := make(map[string]struct{})

	// Adding elements to the set
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["cherry"] = struct{}{}

	// Checking membership
	if _, exists := set["banana"]; exists {
		fmt.Println("banana is in the set")
	}

	// Printing the set
	for key := range set {
		fmt.Println(key)
	}
}
