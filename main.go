package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	return strings.Split(text, " ")
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var inp string
	var cleaned_inp []string

	for true {
		fmt.Print("Pokedex > ")
		reader.Scan()
		inp = reader.Text()
		cleaned_inp = cleanInput(inp)
		fmt.Printf("Your command was: %s\n", cleaned_inp[0])
	}
}
