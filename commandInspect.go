package main

import "fmt"

func commandInspect(c *Config) error {
	if pokemon, ok := Pokedex[c.Argument]; ok {
		fmt.Println("Name: ", pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		fmt.Println(" -hp: ", pokemon.Stats[0].BaseStat)
		fmt.Println(" -attack: ", pokemon.Stats[1].BaseStat)
		fmt.Println(" -defense: ", pokemon.Stats[2].BaseStat)
		fmt.Println(" -special-attack: ", pokemon.Stats[3].BaseStat)
		fmt.Println(" -special-defense: ", pokemon.Stats[4].BaseStat)
		fmt.Println(" -speed: ", pokemon.Stats[5].BaseStat)
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Println(" - ", pokemonType.Type.Name)
		}

	} else {
		return fmt.Errorf("you do not have that pokemon")
	}
	return nil
}
