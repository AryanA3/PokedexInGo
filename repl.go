package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	var inp string
	var cleaned_inp []string

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		inp = reader.Text()
		cleaned_inp = cleanInput(inp)

		if len(cleaned_inp) == 0 {
			continue
		}

		cmd, ok := getCommands()[cleaned_inp[0]]
		if ok {
			err := cmd.callback()
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

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	return strings.Split(text, " ")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
