package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Information Info `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}

func main() {
	fmt.Println("JSON tutorial")

	jsonString := `{"name":"battery sensor", "capacity": 40, "info": {
		"desc": "a sensor reading"
		}}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
