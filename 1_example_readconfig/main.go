package main

import (
	"fmt"

	"1_example_readconfig.com/helper"
)

func main() {
	config := helper.ReadConfig("helper/example_config.json")
	fmt.Println(config)
	fmt.Println(config.Name)
	fmt.Println(config.Nickname)
	fmt.Println(config.Age)
	fmt.Println(config.Info)
}
