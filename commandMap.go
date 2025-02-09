package main

import (
	"encoding/json"
	"fmt"
)

func commandMap(c *Config) error {
	if c.Next == "" {
		c.Next = "https://pokeapi.co/api/v2/location-area/?offset=0"
	}
	data, err := jsonGetter(c.Next)
	if err != nil {
		return nil
	}
	locationArea := PokeAPILocationArea{}
	goop := json.Unmarshal(data, &locationArea)
	if goop != nil {
		return goop
	}
	if locationArea.Previous != "" {
		c.Previous = locationArea.Previous
	}
	c.Next = locationArea.Next
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(c *Config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	data, err := jsonGetter(c.Previous)
	if err != nil {
		return nil
	}
	locationArea := PokeAPILocationArea{}
	goop := json.Unmarshal(data, &locationArea)
	if goop != nil {
		return goop
	}
	c.Previous = locationArea.Previous
	for _, area := range locationArea.Results {
		fmt.Println(area.Name)
	}

	return nil
}

type PokeAPILocationArea struct {
	ID       int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
