package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
	for scanner.Scan() {
		text := scanner.Text()
		uInput := cleanInput(text)
		fmt.Println("Your command was:", uInput[0])
		fmt.Print("Pokedex > ")
	}
}
