package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

// Options are available command-line arguments
type Options struct {
	Recursive bool    `short:"r" description:"Recursive directory traversal"`
	FileName  *string `short:"f" long:"file" description:"Path to file with header text" value-name:"FILE"`
	Extension string  `short:"e" long:"ext" description:"Target files extension" required:"true"`
}

// main run the app
func main() {
	var (
		opts Options
		err  error
	)
	_, err = flags.Parse(&opts)
	if err != nil {
		return
	}

	if opts.FileName != nil {
		// read from file
		err = AddHeaderFromFile("./", *opts.FileName, opts.Extension, opts.Recursive)
	} else {
		// read from unix pipe
		err = AddHeaderFromPipe("./", os.Stdin, opts.Extension, opts.Recursive)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
