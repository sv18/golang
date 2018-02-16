package main

import "fmt"

type p struct {
	fname string
	lname string
	loc   string
}

type q struct {
	fullname string
	loc      string
}

func main() {
	p1 := p{fname: "s", lname: "v", loc: "usa"}
	q1 := p2q(p1)
	fmt.Printf("%#v\n", q1)
}

func p2q(in p) q {
	return q{fullname: in.fname + " " + in.lname, loc: in.loc}
}
