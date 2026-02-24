//Joshua Vann
//COSC 3750
//2-23-26

package main

import "os"

func main() {
	l := os.Args[1:]
	if len(l) == 0 {
		l = append()
	} else {
		print("b")
	}
}

//Removes any hidden files from the dir listing, note not exported
func dirFilter(entries []os.DirEntry) []os.DirEntry {
	var ret = []os.DirEntry{}
	for _, e := range entries {
		if e.Name()[0] == 0x2E && e.Name() != "." {
			print(e.Name()[0])
			continue
		}
		ret = append(ret, e)
	}
}
