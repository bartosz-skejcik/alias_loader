package utils

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"sync"
)

var descriptionCache = make(map[string]string)
var cacheMutex sync.Mutex

func GetAllAliases(aliases map[string]map[string]string) []string {
	var allAliases []string
	for _, commands := range aliases {
		for alias, command := range commands {
			allAliases = append(allAliases, fmt.Sprintf("%s='%s'", alias, command))
		}
	}
	return allAliases
}

func GetRandomAliases(aliasList []string, count int) []string {
	if len(aliasList) == 0 {
		return []string{}
	}
	randAliases := make([]string, 0, count)
	for i := 0; i < count; i++ {
		randAliases = append(randAliases, aliasList[rand.Intn(len(aliasList))])
	}
	return randAliases
}

func ExecuteAlias(alias, command string) {
	cmd := fmt.Sprintf("alias %s='%s'", alias, command)
	exec.Command("zsh", "-c", cmd).Run()
}

func GetRandomCategories(aliases map[string]map[string]string, count int) map[string]map[string]string {
	if count <= 0 || len(aliases) == 0 {
		return nil
	}

	result := make(map[string]map[string]string)
	categories := make([]string, 0, len(aliases))

	for category := range aliases {
		categories = append(categories, category)
	}

	// Get random categories
	for i := 0; i < count && i < len(categories); i++ {
		randIndex := rand.Intn(len(categories))
		category := categories[randIndex]
		result[category] = aliases[category]

		// Remove selected category
		categories[randIndex] = categories[len(categories)-1]
		categories = categories[:len(categories)-1]
	}

	return result
}

func ConvertStringListToMap(list []string) map[string]string {
	result := make(map[string]string)
	for _, aliasString := range list {
		// Split the alias string at the '=' character
		parts := strings.SplitN(aliasString, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}
