package main

import (
	"io"
	"strings"
	"testing"
)

func TestGetCommentBeforeBreak(t *testing.T) {
	type args struct {
		r       io.Reader
		comment string
	}
	tests := []struct {
		name       string
		args       args
		wantString string
		wantIndex  int
	}{
		{
			name: "empty",
			args: args{
				r:       strings.NewReader(""),
				comment: "+",
			},
			wantString: "",
			wantIndex:  0,
		},

		{
			name: "one line",
			args: args{
				r:       strings.NewReader("+test"),
				comment: "+",
			},
			wantString: "test",
			wantIndex:  1,
		},

		{
			name: "wrong line",
			args: args{
				r:       strings.NewReader("-test"),
				comment: "+",
			},
			wantString: "",
			wantIndex:  0,
		},

		{
			name: "some lines",
			args: args{
				r:       strings.NewReader("+test1\n-test2"),
				comment: "+",
			},
			wantString: "test1\n",
			wantIndex:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString, gotIndex := GetCommentBeforeBreak(tt.args.r, tt.args.comment)
			if gotString != tt.wantString {
				t.Errorf("wrong text result '%s', want '%s'", gotString, tt.wantString)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("wrong line count = %d, want %d", gotIndex, tt.wantIndex)
			}
		})
	}
}
