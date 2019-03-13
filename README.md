# Header Adder

The tool help to automate header adding into project source files

## Roadmap

- [ ] Read from pipe
- [x] Read from file
- [ ] Header update
- [x] Recursive directory traversal
- [ ] Binaries
- [ ] Installation

**Now works for Golang files only!**

## Usage

Add header text in a file with a proper formatting (e.g. exactly as you want the text to be in a file, with comment sighs, newlines, etx.).

Then run

```bash
go run cmd/adder/adder.go -f <file with header> -e <extension>
```

Example:

```bash
go run cmd/adder/adder.go -f header.txt -e go
```