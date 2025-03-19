package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type PokemonList struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }

	var pokemonList PokemonList
	err = json.Unmarshal(body, &pokemonList)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, pokemon := range pokemonList.Results {
		fmt.Printf("Name: %s, URL: %s\n", pokemon.Name, pokemon.URL)
	}
}
