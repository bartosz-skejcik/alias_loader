# Alias Loader

Alias Loader is a utility written in Go for managing, loading and printing shell aliases from YAML files.

## Features

- Load aliases from a YAML file.
- Print aliases in a formatted table with optional descriptions.
- Use AI to generate descriptions for aliases.

## Installation

Clone the repository and build the binary using the following commands:

```bash
git clone https://github.com/bartosz-skejcik/alias_loader.git
cd alias_loader
go build -o alias_loader
```

Copy the built file to you home directory:

```bash
cp alias_loader ~/alias_loader
```

## Usage

### Exporting the Groq API Key

Add this line to your `.bashrc` or `.zshrc` file:

```bash
export GROQ_API_KEY="your api key"
```

### Loading Aliases

Add the following to your `.bashrc` or `.zshrc` file:

```bash
# Load aliases
eval "$(~/alias_loader)"
```

### Printing Aliases

```bash
# With description
~/load_aliases --with-description --get-random 3

# Without description
~/load_aliases --get-random 3
```

## Example YAML File

Create a file named `aliases.yaml` with the following content:

```yaml
general:
  ll: "ls -la"
  c: "clear"
  h: "history"
docker:
  dcu: "docker-compose up"
  dcd: "docker-compose down"
```

This file should be placed in your home directory.

## Additional Information

This project uses the Groq AI model for querying alias descriptions and caches the responses for later use.

Fell free to explore the code and modify it to suit your needs.

```

```
