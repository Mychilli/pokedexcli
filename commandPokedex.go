package main

import (
	"fmt"
)

func commandPokedex(c *Config) error {
	if len(Pokedex) == 0 {
		return fmt.Errorf("you have no pokemon")
	}
	fmt.Println("Your Pokedex: ")
	for poke := range Pokedex {
		fmt.Println(poke)
	}
	return nil
}
