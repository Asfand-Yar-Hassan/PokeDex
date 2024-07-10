package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "display help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits the program",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowerString := strings.ToLower(str)
	words := strings.Fields(lowerString)
	return words
}
