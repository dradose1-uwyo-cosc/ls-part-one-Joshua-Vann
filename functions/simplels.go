//Joshua Vann
//COSC 3750
//2-23-26

package main

import (
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
	green := color{start: []byte("\x1b[34m"), end: []byte("\x1b[0m")}
	//def := color{start: []byte("\x1b[37m"), end: []byte("\x1b[0m")}
	for _, a := range args {
		if !useColor {
			w.Write([]byte(a))
		} else {
			fi, _ := os.Lstat(a)
			if fi.Mode().IsDir() {
				blue.ColorPrint(w, a)
			} else if fi.Mode().IsRegular() && (fi.Mode()&0111) != 0 {
				green.ColorPrint(w, a)
			} else {
				w.Write([]byte(a))
			}
		}
	}
}
