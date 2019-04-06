# Header Adder

The tool help to automate header adding into project source files

## Roadmap

- [x] Read from pipe
- [x] Read from file
- [ ] Header update
- [x] Recursive directory traversal
- [ ] Binaries
- [ ] Installation
- [ ] Templates

## Usage

Install package

```bash
go get -u github.com/ilyakaznacheev/header-adder
```

then run

```bash
header-adder -f <file with header> -e <extension>
```

where `<file with header>` is a header text in a file with a proper formatting (e.g. exactly as you want the text to be in a file, with comment sighs, newlines, etx.).

### Example

```bash
header-adder -f header.txt -e go
```

will add text from `header.txt` to any `.go` files in the current directory.

```bash
echo "test header" | go run *.go -r -e txt
```

will add text "test header" o any `.go` files in the current directory and al subdirectories.
