//Joshua Vann
//COSC 3750
//2-23-26

package main

import (
	"io"
	"os"
)

func main() {
	l := os.Args[1:]
	var files []string
	var directory []string
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
		var directoryf []string
		SimpleLS(io.Writer(os.Stdout), files, true)
		for _, dr := range directory {
			r, _ := os.ReadDir(dr)
			fr := dirFilter(r)
			for _, ent := range fr {
				directoryf = append(directoryf, ent.Name())
			}
			w := io.Writer(os.Stdout)
			w.Write([]byte(dr))
			w.Write([]byte(":"))
			w.Write([]byte("\n"))
			SimpleLS(io.Writer(os.Stdout), directoryf, true)
			w.Write([]byte("\n"))
		}

	} else {
		var directoryf []string
		SimpleLS(io.Writer(os.Stdout), files, false)
		for _, dr := range directory {
			r, _ := os.ReadDir(dr)
			fr := dirFilter(r)
			for _, ent := range fr {
				directoryf = append(directoryf, ent.Name())
			}
			w := io.Writer(os.Stdout)
			w.Write([]byte(dr))
			w.Write([]byte(":"))
			w.Write([]byte("\n"))
			SimpleLS(io.Writer(os.Stdout), directoryf, false)
			w.Write([]byte("\n"))
		}
	}
}

// Removes any hidden files from the dir listing, note not exported
func dirFilter(entries []os.DirEntry) []os.DirEntry {
	var ret = []os.DirEntry{}
	for _, e := range entries {
		if e.Name()[0] == 0x2E {
			print(e.Name()[0])
			continue
		}
		ret = append(ret, e)
	}
	return ret
}
