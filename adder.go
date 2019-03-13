// Package adder adds copyright (or any other text) to all files of certain kind in directory
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// AddHeaderFromFile adds header text from a text file
//
// fromDir: root directory for action
// pathToFile: path to header text file
// ext: file extension
// recursive: flag that enables recursive directory traversal
func AddHeaderFromFile(fromDir, pathToFile string, ext string, recursive bool) error {
	text, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return err
	}

	return AddHeader(fromDir, text, ext, recursive)
}

// AddHeaderFromPipe adds header text from a text file
//
// fromDir: root directory for action
// r: text reader interface
// ext: file extension
// recursive: flag that enables recursive directory traversal
func AddHeaderFromPipe(fromDir string, r io.Reader, ext string, recursive bool) error {
	text, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return AddHeader(fromDir, text, ext, recursive)
}

// AddHeader adds header text to project files
func AddHeader(fromDir string, text []byte, ext string, recursive bool) error {

	filesToHandle := make([]string, 0)
	if recursive {
		err := filepath.Walk(fromDir, func(filepath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
				filesToHandle = append(filesToHandle, path.Join(fromDir, filepath))
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else {
		files, err := ioutil.ReadDir(fromDir)
		if err != nil {
			return err
		}
		for _, info := range files {
			if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
				filesToHandle = append(filesToHandle, path.Join(fromDir, info.Name()))
			}
		}
	}

	if len(filesToHandle) == 0 {
		return fmt.Errorf("no files to change")
	}

	for _, path := range filesToHandle {
		fmt.Println(path)
		f, err := os.OpenFile(path, os.O_RDWR, 0600)
		if err != nil {
			return err
		}
		fb, _ := ioutil.ReadAll(bufio.NewReader(f))

		fileText := append(text, append([]byte("\n\n"), fb...)...)

		_, err = f.WriteAt(fileText, 0)
		f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
