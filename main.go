package main

import "fmt"

func main() {
	config, err := readFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = scheduler(config.Reminders, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	select {}

}
