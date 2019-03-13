package main

import (
	"fmt"
	"os"

	adder "github.com/ilyakaznacheev/header-adder"
	"github.com/jessevdk/go-flags"
)

// Options are available command-line arguments
type Options struct {
	Recursive bool   `short:"r" description:"Recursive directory traversal"`
	FileName  string `short:"f" long:"file" description:"Path to file with header text" value-name:"FILE"`
	Extension string `short:"e" long:"ext" description:"Target files extension" required:"true"`
}

// main run the app
func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	if err := adder.AddHeaderFromFile("./", opts.FileName, opts.Extension, opts.Recursive); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
