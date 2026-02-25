//Joshua Vann
//COSC 3750
//2-23-26

package main

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
	//blue := color{start: []byte("\x1b[34m"), end: []byte("\x1b[0m")}
	//green := color{start: []byte("\x1b[34m"), end: []byte("\x1b[0m")}
	def := color{start: []byte("\x1b[37m"), end: []byte("\x1b[0m")}
	for _, a := range args {
		if !useColor {
			def.ColorPrint(w, a)
		} else {
			fi, err := os.Lstat(a)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error with file")
				return
			} else {
				print(fi)
			}
		}
	}
}
