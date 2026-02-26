//Joshua Vann
//COSC 3750
//2-23-26

package functions

import (
	"os"
)

// Removes any hidden files from the dir listing, note not exported
func dirFilter(entries []os.DirEntry) []os.DirEntry {
	var ret = []os.DirEntry{}
	for _, e := range entries {
		if e.Name()[0] == 0x2E {
			continue
		}
		ret = append(ret, e)
	}
	return ret
}





