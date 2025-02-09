package main

import (
	"bufio"
	"fmt"
	"os"
)

type Config struct {
	Next     string
	Previous string
	Argument string
}

var Pokedex = make(map[string]PokeAPIPokemonStats)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	cfg := Config{}

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := CleanInput(scanner.Text())
		if len(input) > 1 {
			cfg.Argument = input[1]
		}
		if len(input) == 0 {
			fmt.Println("Please insert a command")
			continue
		}
		commandName := input[0]
		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
