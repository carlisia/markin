# Markin

A CLI tool for managing markdown notes, with a focus on Obsidian vaults.

## Features

- Add different types of note entries to markdown files
- Support for Obsidian vaults
- Configurable section names and positions
- Automatic section creation if missing
- Timestamp prefix for entries
- Silent operation with no terminal output

## Installation

```bash
go install github.com/carlisia/markin@latest
```

## Configuration

Create a configuration file at `$HOME/.config/markin/.markin.yaml`:

```yaml
markdown_dir: $VAULT_MAIN
markdown_file: daily.md
section: "## 💭 ✍️ ✨ Notes"
position: after-heading
create_section_if_missing: true
```

### Configuration Options

- `markdown_dir`: Directory containing the markdown file (can use environment variables)
- `markdown_file`: Name of the markdown file to modify
- `section`: Section name to add entries to (default: "## 💭 ✍️ ✨ Notes")
- `position`: Where to add entries in the section ("after-heading" or "before-end")
- `create_section_if_missing`: Whether to create the section if it doesn't exist

## Usage

Initialize the configuration:

```bash
markin init
```

Add a fleeting note entry:

```bash
markin fl "Your fleeting thought here"
```

This will add an entry like:

```markdown
- ⚡ *06:33:45 pm:* **Fleeting**:: Your fleeting thought here
```

## Development

Build the project:

```bash
go build -o markin cmd/markin/main.go
```

Run tests:

```bash
go test ./...
```

## License

MIT
