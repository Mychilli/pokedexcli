package main

import (
	//"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	//"fmt"
	"strings"
)

func CleanInput(text string) []string {
	result := strings.ToLower(text)
	words := strings.Fields(result)
	return words
}
func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *Config) error {
	for outside, inside := range GetCommands() {
		fmt.Printf("%v: %v\n", outside, inside.description)
	}
	return nil
}

func jsonGetter(URL string) ([]byte, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "pull up the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "goto previous page",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "check for pokemon",
			callback:    commmandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "show pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "display your pokedex",
			callback:    commandPokedex,
		},
	}
}
