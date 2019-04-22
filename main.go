package main

import (
	"fmt"
	"graphUtils/graphs"
	"os"
)

func main() {
	fname := os.Args[1]
	fmt.Printf("file name %s \n", fname)
	f, err := os.Open(fname)
	if err != nil {
		fmt.Errorf("error opening file " )
	}
	g := graphs.ReadUgraph("/", f)
	ugraph := g.Ugraph()

	fmt.Println(g.Adj("'Breaker' Morant (1980)"))

	f.Close()
}
