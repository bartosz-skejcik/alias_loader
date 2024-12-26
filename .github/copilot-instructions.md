# Aliases - Terminal welcoming addon

Every time you choose to apply a rule(s), explicitly state the rule(s) in the output. You can abbreviate the rule description to a single word or phrase.

## Features

-   Running with no flags and arguments will result in the program loading all the aliases to the terminal shell
-   Running with a flag `--get-random <number>` will result in the program printing a pretty table with randomly selected n number of aliases from the `~/.aliases.yaml` file. If no number is provided it will default to 3.

## Printing the table

-   The program prints a list of aliases pretty formatted. It uses nerd-font icons and its own utils function for printing tables.
-   It formats each alias with the `alias` and its `command` also providing an ai description of the alias.

## Description

This is a simple terminal welcoming addon that displays a welcoming message when you run this command from you .bashrc or .zshrc file.
Its' purpose is to remind you of `n` number of aliases that you have set up in your terminal.
The aliases are located inside the `~/.aliases.yaml` file.

## Code style and Structure

-   Write concise, technical go code with accurate examples
-   Use functional and declarative programming patterns; avoid imperative programming
-   Prefer iteration and code modularization over code duplication
-   Use descriptive variable names with auxilary verbs (e.g. isLoading, hasError)
-   Structure repository files as follows:

```
internal/
    utils/
        main.go         # Utility functions used across the project
        ai.go           # Utility functions for ai functionality
    load/
        main.go         # Loading and managing the aliases
    printer/
        main.go         # Printing, customizing and managing the table printing
main.go                 # Main program entry
```

## TechStack

-   Golang

## Naming Conventions

-   Use lowercase with underscores for directories
-   Favor named exports for functions and utilities

## Ai Features

-   Use the groq llama-3.1 70b model for quering for the alias description. Cache the responses so that later if the drawn alias from the aliases pool already has a cached description, it will be used to limit the ai usage.
