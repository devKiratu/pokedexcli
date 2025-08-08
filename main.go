package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/devKiratu/pokedexcli/internal/pokecache"
)


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
				name: "exit",
				description: "Exit the Pokedex",
				callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the names of (next) 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback: commandMapb,
		},
	}
}

type cliCommand struct {
	name string
	description string
	callback func(*pagesNave) error
}

type pokeResult struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
		} `json:"results"`
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type pagesNave struct {
	next string
	previous string
}

var nav = &pagesNave{}
var cache *pokecache.Cache

func commandMap(nav *pagesNave) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if nav.next != "" {
		url = nav.next
	}
	err := makeApiRequest(url)
	if err != nil {
		return err
	}
	return nil
}

func makeApiRequest(url string) error {
// read values from the cache, if they exist
	if data, ok := cache.Get(url); ok {
		var apiResults pokeResult
		err := json.Unmarshal(data, &apiResults)
		if err != nil {
			return err
		}
		for _, item := range apiResults.Results {
		fmt.Println(item.Name)
	}
	return nil
	}
// prepare request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var apiResults pokeResult
	decoder:= json.NewDecoder(res.Body)
	err = decoder.Decode(&apiResults)
	if err != nil {
		return err
	}
	//cache results
	data, err := json.Marshal(apiResults)
	if err != nil {
		return err
	}
	cache.Add(url, data)
	// fmt.Println(apiResults)
	nav.next = apiResults.Next
	nav.previous = apiResults.Previous
	for _, item := range apiResults.Results {
		fmt.Println(item.Name)
	}
	return nil
}

func commandMapb(nav *pagesNave) error {
	if nav.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	err := makeApiRequest(nav.previous)
	if err != nil {
		return err
	}
	return nil
}

func commandExit(nav *pagesNave) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(nav *pagesNave) error {
	help := `
Welcome to the Pokedex!
Usage:
	`
	fmt.Println(help)
	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return  nil
}

func main() {
	cache = pokecache.NewCache(time.Second * 5)

  scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		userInput := cleanInput(text)

		if len(userInput) < 1 {
			fmt.Print("Pokedex > ")
			continue
		}

		if c, ok := getCommands()[userInput[0]]; ok {
			c.callback(nav)
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("Pokedex > ")
	}
}
