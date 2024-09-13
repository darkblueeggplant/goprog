package main

import (
	"fmt"

	"2_example_nested.com/nested_helper"
)

func main() {
	config := nested_helper.ReadConfig("nested_helper/example_config.json")
	fmt.Println(config)
}
