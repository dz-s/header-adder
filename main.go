package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

const (
	description = `                                             
+_____           _            _____   _   _         
|  |  |___ ___ _| |___ ___   |  _  |_| |_| |___ ___ 
|     | -_| .'| . | -_|  _|  |     | . | . | -_|  _|
|__|__|___|__\|___|___|_|    |__|__|___|___|___|_|                                               

by Ilya Kaznacheev (c) 2019 via MIT license.`
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
	p := flags.NewParser(&opts, flags.Default)
	// p.Usage = usage
	p.LongDescription = description
	_, err = p.Parse()
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
