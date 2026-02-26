//Joshua Vann
//COSC 3750
//2-23-26

package functions

import (
	"fmt"
	"io"
	"os"
)

type color struct {
	start []byte
	end   []byte
}

// A simple ls for when no flags provided
func SimpleLS(w io.Writer, args []string, useColor bool) {
	blue := color{start: []byte("\x1b[34m"), end: []byte("\x1b[0m")}
	green := color{start: []byte("\x1b[32m"), end: []byte("\x1b[0m")}
	for _, a := range args {
		if !useColor {
			for i, let := range a {
				if let == rune('/') {
					a = a[i+1:]
				}
			}
			w.Write([]byte(a))
			w.Write([]byte("\n"))
		} else {
			fi, err := os.Lstat(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "gols: Could not access %s, no such file or directory\n", a)
				continue
			}
			for i, let := range a {
				if let == rune('/') {
					a = a[i+1:]
				}
			}
			if fi.Mode().IsDir() {
				blue.ColorPrint(w, a)
			} else if fi.Mode().IsRegular() && (fi.Mode()&0111) != 0 {
				green.ColorPrint(w, a)
			} else {
				w.Write([]byte(a))
				w.Write([]byte("\n"))
			}
		}
	}
}






