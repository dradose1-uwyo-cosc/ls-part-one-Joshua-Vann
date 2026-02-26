//Joshua Vann
//COSC 3750
//2-23-26

package functions

import "os"

//Checks if outputting to a terminal or being piped/redirected
func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
