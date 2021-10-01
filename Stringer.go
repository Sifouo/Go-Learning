package main

import "fmt"

type persons struct {
	name    string
	familly string
}

type Stringer interface {
	String() string
}

func printstr() {
	me := persons{name: "Ali", familly: "Ahmadi"}
	fmt.Println(me)
}

func (p persons) String() string {
	return fmt.Sprintf("%v  (%v)", p.name, p.familly)
}
