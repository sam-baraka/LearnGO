package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonString := `{"name":"battery sensor","capacity":40,"time":"2020-04-05",
    "info":{
        "desc":"sensor reading"
    }}`

	var reading map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", reading)

}
