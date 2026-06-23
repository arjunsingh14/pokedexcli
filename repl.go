package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/arjunsingh14/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

type config struct {
	client pokeapi.Client
	nextLocationsUrl *string
	previousLocationsUrl *string
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
		"explore": {
			name: "explore",
			description: "Lists pokemon in an area",
			callback: commandExplore,
		},
		 
	}
}

func startRepl(cfg *config) {
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
		commandName := input[0]
		val, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		args := input[1:]
		err := val.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}


	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input:", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}