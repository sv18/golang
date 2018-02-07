package main

import (
	"fmt"
)

func main() {
	s := []string{"hello", "world"}
	m := slicectomap(s)
	fmt.Printf("%v", m)
}

func slicectomap(s []string) map[string]struct{} {
	m := map[string]struct{}{}
	for _, v := range s {
		m[v] = struct{}{} //note: duplicates will be eliminated
	}
	return m
}
