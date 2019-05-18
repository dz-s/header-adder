package main

import (
	"bufio"
	"io"
	"strings"
)

// GetCommentBeforeBreak returns commented strings before first string that starts without comment symbol
func GetCommentBeforeBreak(r io.Reader, commentSign string) (string, int) {
	// check parameters validity
	if r == nil || commentSign == "" {
		return "", 0
	}

	var (
		buf, str string
		err      error
		n        int
	)

	// // concatenate strings respecting first line
	// addString := func(base, apd string) (res string) {
	// 	if base == "" {
	// 		res = apd
	// 	} else {
	// 		res = base + "\n" + apd
	// 	}
	// 	return
	// }

	sr := bufio.NewReader(r)

	for err == nil {
		str, err = sr.ReadString('\n')
		if str == "" {
			continue
		}
		if !strings.HasPrefix(str, commentSign) {
			break
		}
		// buf = addString(buf, strings.TrimPrefix(str, commentSign))
		buf += strings.TrimPrefix(str, commentSign)
		n++
	}

	return buf, n
}
