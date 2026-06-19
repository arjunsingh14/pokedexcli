package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	next string
	previous string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
    	"exit": {
        	name:        "exit",
        	description: "Exit the Pokedex",
        	callback:    commandExit,
    	},
		"help": {
			name: "help",
			description: "Display a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Lists location areas",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists previous locations",
			callback: commandMapb,
		},
	}
}

func startRepl() {
	cfg := config{next: "https://pokeapi.co/api/v2/location-area/", previous: ""}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		val, ok := getCommands()[input[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		val.callback(&cfg)
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input:", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}