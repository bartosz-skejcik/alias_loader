package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"load_aliases/internal/load"
	"load_aliases/internal/printer"
	"load_aliases/internal/utils"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	yamlFile := os.Getenv("HOME") + "/.aliases.yaml"

	aliases, err := load.LoadAliases(yamlFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := os.Args[1:]
	var withDescriptions bool

	for i, arg := range args {
		if arg == "--get-random" {
			if i+1 < len(args) {
				numAliases, err := strconv.Atoi(args[i+1])
				if err != nil {
					fmt.Println("Invalid number of aliases provided. Please provide an integer")
					os.Exit(1)
				}
				randomAliases := utils.GetRandomAliases(utils.GetAllAliases(aliases), numAliases)
				printer.PrintAliasesTable(map[string]map[string]string{
					"Random Aliases": utils.ConvertStringListToMap(randomAliases),
				}, withDescriptions)
				return
			} else {
				fmt.Println("Please provide the number of random aliases to retrieve after the --get-random flag")
				os.Exit(1)
			}
		} else if arg == "--with-description" {
			withDescriptions = true
		}
	}

	// print if the args are not empty
	if len(args) > 0 {
		printer.PrintAliasesTable(aliases, withDescriptions)
		return
	}

	// Default behavior: register all aliases
	// fmt.Println("Registering all aliases...")
	// fmt.Println(aliases)
	for _, aliasMap := range aliases {
		for alias, cmd := range aliasMap {
			fmt.Printf("alias %s='%s'\n", alias, cmd)
			// utils.ExecuteAlias(alias, cmd)
		}
	}
}
