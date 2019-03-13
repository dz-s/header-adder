// Package adder adds copyright (or any other text) to all files of certain kind in directory
package adder

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AddHeaderFromFile(fromDir, pathToFile string, ext string, recursive bool) error {
	text, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return err
	}

	return AddHeader(fromDir, text, ext, recursive)
}

// AddHeader adds header text to project files
func AddHeader(fromDir string, text []byte, ext string, recursive bool) error {

	filesToHandle := make([]os.FileInfo, 0)
	if recursive {
		err := filepath.Walk(fromDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
				filesToHandle = append(filesToHandle, info)
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
				filesToHandle = append(filesToHandle, info)
			}
		}
	}

	if len(filesToHandle) == 0 {
		return fmt.Errorf("no files to change")
	}

	for _, f := range filesToHandle {
		f, err := os.OpenFile(f.Name(), os.O_RDWR, 0600)
		if err != nil {
			return err
		}
		fb, _ := ioutil.ReadAll(bufio.NewReader(f))

		text = append(text, append([]byte("\n\n"), fb...)...)

		_, err = f.WriteAt(text, 0)
		f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
