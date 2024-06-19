package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1000000
	str := "badges"

	m1 := make(map[string]*string)

	printAlloc()

	for i := 0; i < n; i++ { // Adds 1 million elements
		key := fmt.Sprintf("%s%d", str, i) // Ensure unique keys
		val := ""
		m1[key] = &val
	}

	printAlloc()
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
}
