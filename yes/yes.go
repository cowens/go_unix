package main

import (
	f "fmt"
	"os"
	s "strings"
)

const bufsize = 1024*64

func main() {
	atom := "y\n"

	if len(os.Args) > 1 {
		atom = s.Join(os.Args[1:], " ") + "\n"
	}

	buf := s.Repeat(atom, bufsize / len(atom))

	for {
		f.Print(buf)
	}
}
