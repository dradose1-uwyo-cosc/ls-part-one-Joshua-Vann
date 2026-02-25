//Joshua Vann
//COSC 3750
//2-23-26

package main

import "io"

//Prints appropriate entries in color, or colorless if regular file.
func (c color) ColorPrint(w io.Writer, s string) {
	w.Write(c.start)
	w.Write([]byte(s))
	w.Write(c.end)
}
