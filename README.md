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
./rsn -in my-notes.excalidraw -start 31
```

- `-f` : input file (relative to `./docs/`)
- `-s` : starting day number (default: `1`)

The updated file will be saved as:

```bash
./docs/RSN.excalidraw
```

## 📂 Project Structure

```bash
rs-nihongo-notes/
├── cmd/              # CLI entrypoint
├── internal/
│   ├── models/       # Excalidraw data structures
│   └── utils/        # Helpers for error handling, etc.
└── docs/             # Your Excalidraw notes (input/output)
```

## 🔮 Roadmap

- [ ] Dry-run mode to preview changes
- [ ] Add `-o` flag to choose output filename
- [ ] **Custom month**: use `-m` flag to insert the current learning month (e.g., "October")
- [ ] **Custom year**: use `-y` flag to insert the learning year (e.g., "2025")
- [ ] **Custom JLPT level**: allow tagging notes with `-n N5|N4|N3|N2|N1`
- [ ] **Custom generation**: add `-g` flag to mark study generations (e.g., "Gen 1", "Gen 2", "Gen 3")
- [ ] **CSV to template's table**: add a feature to convert CSV files into the template's table format for easier note creation

## 📜 License

[MIT](./LICENSE). Feel free to use, modify, and share.
