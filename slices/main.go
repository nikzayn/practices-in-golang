package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	data []string
}

func main() {
	foos := make([]Foo, 1000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			data: make([]string, 1024*1024),
		}
	}
	printAlloc()

	val := keepFirstTwoElementsOnlyCopy(foos)
	runtime.GC() //triggers the garbage collector to run
	printAlloc()
	runtime.KeepAlive(val) //prevents a specific object from being collected
}

func firstTwoElements(foo []Foo) []Foo {
	return foo[:10]
}

func keepFirstTwoElementsOnlyCopy(foos []Foo) []Foo {
	res := make([]Foo, 10)
	copy(res, foos)
	return res
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
}
