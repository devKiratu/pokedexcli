package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var supportedCommands map[string]cliCommand = map[string]cliCommand{
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
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	help := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
	`
	fmt.Println(help)
	return  nil
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		userInput := cleanInput(text)

		if c, ok := supportedCommands[userInput[0]]; ok {
			c.callback()
		} else {
			fmt.Println("Unknown command")
		}

		fmt.Print("Pokedex > ")
	}
}
