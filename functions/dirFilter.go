//Joshua Vann
//COSC 3750
//2-23-26

package main

import (
	"io"
	"os"
)

//TODO:
//1. Enable sorting
//2. Figure out how to figure out what output is.
//3. Brighter colors?
//4. Fix memory exception error. Caused by trying to read lstat data in wrong directory.

func main() {
	l := os.Args[1:]
	var files = []string{}
	var directory = []string{}
	var directoryf = []string{}
	color := false
	if len(l) == 0 {
		l = append(l, ".")
	}
	for _, a := range l {
		info, _ := os.Lstat(a)
		if info.Mode().IsDir() {
			directory = append(directory, a)
		} else {
			files = append(files, a)
		}
	}
	if true {
		color = true
	} else {
		color = false
	}
	SimpleLS(io.Writer(os.Stdout), files, color)
	for _, dr := range directory {
		r, _ := os.ReadDir(dr)
		fr := dirFilter(r)
		for _, ent := range fr {
			directoryf = append(directoryf, dr+"/"+ent.Name())
		}
		w := io.Writer(os.Stdout)
		w.Write([]byte(dr))
		w.Write([]byte(":"))
		w.Write([]byte("\n"))
		SimpleLS(io.Writer(os.Stdout), directoryf, color)
		w.Write([]byte("\n"))
		directoryf = []string{}
	}
}

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





