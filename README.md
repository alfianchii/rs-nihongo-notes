# rs-nihongo-notes

> CLI tool to re-use Excalidraw study notes for Japanese learning each month.

This tool scans your `.excalidraw` file, finds all text elements that start with Day, Month, Gen, or something like that, sorts them starting at a user-defined value.  
That way you can easily **reset your study plan for the next month** without manually editing dozens of text boxes.

---

## ✨ Feature(s)

- Custom Days

---

## 🚀 Installation

Clone this repo and build the binary:

```bash
git clone https://github.com/yourusername/rs-nihongo-notes.git
cd rs-nihongo-notes
go build -o rsn ./cmd
```

## 📝 Usage

Suppose you have a file `my-notes.excalidraw` in the `./docs` folder.

```bash
./rsn -f my-notes -s 31 -docs my-documents -o cute-note -dry-run
```

### Options

- `-f` : input file (relative to `./docs/`).
- - You may omit the `.excalidraw` extension; it will be added automatically.
- `-s` : starting day number (default: `1`).
- `-docs` : Docs root directory (default: `./docs`). All input/output paths are resolved relative to this directory.
- `-o` : Output file name (relative to `-docs`, default: `RSN.excalidraw`).
- - Yes, you can omit the `.excalidraw` too.
- - Example: `-o cute-note.excalidraw` → writes `-o cute-note`.
- `-dry-run` : Preview only. Prints the renumbered text elements to stdout without writing the output file.

### Examples

```bash
# Preview changes only, don’t write file
./rsn -f my-notes -s 31 -dry-run

# Write output file "cute-note.excalidraw" into ./docs
./rsn -f my-notes -s 31 -o cute-note

# Use a custom docs directory
./rsn -f my-notes -s 10 -docs my-documents
```

## 📂 Project Structure

```bash
rs-nihongo-notes/
├── cmd/                 # CLI entrypoint (main program)
│   └── main.go
├── docs/                # Your Excalidraw notes (input/output root)
├── internal/
│   ├── app/             # Application orchestration (Run)
│   │   └── app.go
│   ├── cli/             # Flags & CLI options parsing
│   │   └── flags.go
│   ├── excalidraw/      # Excalidraw data model (I/O, types, and features)
│   │   ├── io.go
│   │   ├── model.go
│   │   └── services/
│   └── utils/           # Helpers
│       └── validation.go
├── go.mod
├── LICENSE
└── README.md
```

## 🔮 Roadmap

- [x] Dry-run mode to preview changes
- [x] Add `-o` flag to choose output filename (relative to `./docs`)
- [x] Add `-docs` flag to choose docs directory (where you would put `.excalidraw` files)
- [ ] **Interactive logs**: enhance CLI output with colors, progress indicators, or interactive prompts for a better user experience
- [ ] **Custom month**: use `-m` flag to insert the current learning month (e.g., "October")
- [ ] **Custom year**: use `-y` flag to insert the learning year (e.g., "2025")
- [ ] **Custom JLPT level**: allow tagging notes with `-n N5|N4|N3|N2|N1`
- [ ] **Custom Gen**: add `-g` flag to mark study generations (e.g., "Gen 1", "Gen 2", "Gen 3")
- [ ] **Daily Notes**: create daily notes template (color like header and its body, day, month, year, gen, N-level)
- [ ] **CSV to template's table**: add a feature to convert CSV files into the template's table format for easier note creation
- [ ] **One to Many**: move the daily notes into monthly notes

## 📜 License

[MIT](./LICENSE). Feel free to use, modify, and share.
