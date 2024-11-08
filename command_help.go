package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("PokeDex Help Menu")
	fmt.Println("Available commands:")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" -%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
