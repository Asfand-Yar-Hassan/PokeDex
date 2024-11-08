package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		err := command.callback(cfg)

		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "list all the area locations",
			callback:    callBackMap,
		},
		"map --prev": {
			name:        "map --prev",
			description: "list all previous locations",
			callback:    callBackMapPrev,
		},
	}
}

func cleanInput(str string) []string {
	lowerString := strings.ToLower(str)
	words := strings.Fields(lowerString)
	return words
}
