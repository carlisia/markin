# Markin

A CLI tool for managing markdown notes, with a focus on Obsidian vaults.

## Features

- Add different types of note entries to markdown files
- Support for Obsidian vaults
- Configurable section names and positions
- Automatic section creation if missing
- Timestamp prefix for entries
- Silent operation with no terminal output

## Configuration

Create a configuration file at `$HOME/.config/markin/.markin.yaml`:

```yaml
project_dir: $VAULT_MAIN
daily_note_path: $DAILY_NOTES
daily_note_name: daily.md
section: "## 💡 🧠 🔥 Fleeting Ideas"
position: after-heading
create_section_if_missing: true
```

### Configuration Options

- `project_dir`: Directory containing your markdown files (can use environment variables)
- `daily_note_path`: Directory containing your daily notes (can use environment variables)
- `daily_note_name`: Name of the daily note file to modify
- `section`: Section name to add entries to (default: "## 💡 🧠 🔥 Fleeting Ideas")
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

## Installation

- Clone the repository:

```bash
git clone https://github.com/carlisia/markin.git
cd markin
```

- Build the executable:

```bash
go build -o markin cmd/markin/main.go
```

- Install the executable to your bin directory:

```bash
cp markin ~/go/bin/  # or any other directory in your PATH
```

## License

MIT
