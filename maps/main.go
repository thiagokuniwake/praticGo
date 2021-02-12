package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	fmt.Println(m["c"])

	_, exists := m["c"]
	fmt.Println(exists)

	value, exists := m["a"]
	fmt.Println(value)

	var x = map[string]int{}
	fmt.Println(x)

	newMap := map[string]int{"x": 2, "y": 5}
	fmt.Println(newMap)

	if value, exists := m["c"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("ops!")
	}
}
