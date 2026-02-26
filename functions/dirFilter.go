//Joshua Vann
//COSC 3750
//2-23-26

package main

import (
	"io"
	"os"
	"sort"
)

//TODO:
//1. Enable sorting, looks like it works?

func main() {
	l := os.Args[1:]
	var files = []string{}
	var directory = []string{}
	var directoryf = []string{}
	color := false
	out := os.Stdout
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
	if IsTerminal(out) {
		color = true
	} else {
		color = false
	}
	w := io.Writer(os.Stdout)
	sort.Strings(files)
	SimpleLS(io.Writer(os.Stdout), files, color)
	if len(files) != 0 && len(directory) != 0 {
		w.Write([]byte("\n"))
	}
	sort.Strings(directory)
	for i, dr := range directory {
		r, _ := os.ReadDir(dr)
		fr := dirFilter(r)
		for _, ent := range fr {
			directoryf = append(directoryf, dr+"/"+ent.Name())
		}
		sort.Strings(directoryf)
		w.Write([]byte(dr))
		w.Write([]byte(":"))
		w.Write([]byte("\n"))
		SimpleLS(io.Writer(os.Stdout), directoryf, color)
		if i != len(directory)-1 {
			w.Write([]byte("\n"))
		}
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





