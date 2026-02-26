//Joshua Vann
//COSC 3750
//2-23-26

package main

import (
	"io"
	"os"
	"slices"
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
	sort := func(i, j string) int {
		a := len(i)
		b := len(j)
		for k := 0; k < a && k < b; k++ {
			if i[k] == j[k] {
				continue
			}
			return int(byte(i[k]) - byte(j[k]))
		}
		return a - b
	}
	if IsTerminal(out) {
		color = true
	} else {
		color = false
	}
	if len(l) == 0 || (len(l) == 1 && l[0] == ".") {
		r, _ := os.ReadDir(".")
		fr := dirFilter(r)
		for _, ent := range fr {
			directoryf = append(directoryf, "./"+ent.Name())
		}
		slices.SortFunc(directoryf, sort)
		SimpleLS(io.Writer(os.Stdout), directoryf, color)
		directoryf = []string{}
		return
	}
	for _, a := range l {
		info, _ := os.Lstat(a)
		if info.Mode().IsDir() {
			directory = append(directory, a)
		} else {
			files = append(files, a)
		}
	}
	w := io.Writer(os.Stdout)
	slices.SortFunc(files, sort)
	SimpleLS(io.Writer(os.Stdout), files, color)
	if len(files) != 0 && len(directory) != 0 {
		w.Write([]byte("\n"))
	}
	slices.SortFunc(directory, sort)
	for i, dr := range directory {
		r, _ := os.ReadDir(dr)
		fr := dirFilter(r)
		for _, ent := range fr {
			directoryf = append(directoryf, dr+"/"+ent.Name())
		}
		slices.SortFunc(directoryf, sort)
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





