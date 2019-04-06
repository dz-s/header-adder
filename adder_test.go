// Package adder adds copyright (or any other text) to all files of certain kind in directory
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestAddHeaderPlain(t *testing.T) {
	type args struct {
		fromDir string
		text    string
		ext     string
	}
	type textfile struct {
		name     string
		ext      string
		filename string
		text     string
	}
	tests := []struct {
		name    string
		files   []textfile
		args    args
		wantErr bool
	}{
		{
			name: "single file",
			files: []textfile{
				textfile{
					name: "file 1",
					ext:  "txt",
					text: "some text",
				},
			},
			args: args{
				text: "test header",
				ext:  "txt",
			},
			wantErr: false,
		},

		{
			name: "many files",
			files: []textfile{
				textfile{
					name: "file 1",
					ext:  "txt",
					text: "some text",
				},
				textfile{
					name: "file 2",
					ext:  "txt",
					text: "other text",
				},
				textfile{
					name: "file 3",
					ext:  "txt",
					text: "different text",
				},
			},
			args: args{
				text: "test header",
				ext:  "txt",
			},
			wantErr: false,
		},

		{
			name: "wrong extension",
			files: []textfile{
				textfile{
					name: "file 1",
					ext:  "xml",
					text: "some text",
				},
				textfile{
					name: "file 2",
					ext:  "html",
					text: "other text",
				},
				textfile{
					name: "file 3",
					ext:  "json",
					text: "different text",
				},
			},
			args: args{
				text: "test header",
				ext:  "txt",
			},
			wantErr: true,
		},

		{
			name: "different extensions",
			files: []textfile{
				textfile{
					name: "file 1",
					ext:  "txt",
					text: "some text",
				},
				textfile{
					name: "file 2",
					ext:  "yml",
					text: "other text",
				},
			},
			args: args{
				text: "test header",
				ext:  "txt",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromDir := fmt.Sprintf("%s/%d", os.TempDir(), time.Now().UnixNano())

			if _, err := os.Stat(fromDir); os.IsNotExist(err) {
				if err := os.Mkdir(fromDir, 0700); err != nil {
					t.Fatal("failed to create temporary directory:", err)
				}
			}
			for idx, f := range tt.files {
				tmpFile, err := ioutil.TempFile(fromDir, "headertest_*."+f.ext)
				if err != nil {
					t.Fatal("cannot create temporary file:", err)
				}

				tt.files[idx].filename = tmpFile.Name()

				text := []byte(f.text)
				if _, err = tmpFile.Write(text); err != nil {
					t.Fatal("failed to write to temporary file:", err)
				}
			}

			if err := AddHeader(fromDir, []byte(tt.args.text), tt.args.ext, false); (err != nil) != tt.wantErr {
				t.Errorf("unexpected error behavior %v, want %v", err, tt.wantErr)
			}

			for _, f := range tt.files {
				text, err := ioutil.ReadFile(f.filename)
				if err != nil {
					t.Errorf("failed on file read: %s", err.Error())
					continue
				}

				textStr := string(text)
				textWant := fmt.Sprintf("%s\n\n%s", tt.args.text, f.text)
				if tt.args.ext == f.ext && textStr != textWant {
					t.Errorf("wrong file text %s want %s", textStr, textWant)
				}
				if tt.args.ext != f.ext && textStr == textWant {
					t.Errorf("wrong file text %s want %s", textStr, textWant)
				}
				// os.Remove(f.filename)
			}

			if err := os.RemoveAll(fromDir); err != nil {
				t.Fatal("failed to remove temporary directory:", err)
			}
		})
	}
}
