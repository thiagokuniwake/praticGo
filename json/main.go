package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"car"`
	Year  int    `json:"-"` //Not create tag json
	Color string
}

func main() {
	//Marshal
	car := Car{"Gol", 2019, "Black"}
	result, _ := json.Marshal(car)
	fmt.Println(string(result))

	//Unmarshal
	car2 := Car{}
	err := json.Unmarshal(result, &car2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(car2)
}
