# Markin

A CLI tool for inserting small blurbs, snippets, text updates into markdown notes. It pairs well with Obsidian, but can serve as a note tracker for any set of markdown files.

## Features

- Insert lines into specific sections of markdown files
- Support for Obsidian vaults
- Configurable section names and positions
- Automatic section creation if missing
- Timestamp prefix for inserted lines
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
- `section`: Section name to insert lines into (default: "## 💭 ✍️ ✨ Notes")
- `position`: Where to insert lines in the section ("after-heading" or "before-end")
- `create_section_if_missing`: Whether to create the section if it doesn't exist

## Usage

Initialize the configuration:

```bash
markin init
```

Insert a line:

```bash
markin insert "Your note here"
```

This will insert a line like:

```markdown
- ⚡ _06:33:45 pm:_ **Fleeting**:: Your note here
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
