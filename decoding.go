package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func decode() {
	layoutISO := "2006-01-02"
	dob, _ := time.Parse(layoutISO, "2010-01-18")
	Mahdi := Custumer{
		FirstName: Name{FirstName: "Mahdi"},
		Family:    "Rezaei",
		Address: address{
			Line1: "Isfahan",
			Line2: "Zeinabieh",
			Line3: "Soodan",
		},
		Dob: dob,
	}

	mahdiJson, err := json.MarshalIndent(Mahdi, "", "    ")
	if err == nil {
		fmt.Println(string(mahdiJson))
	} else {
		fmt.Println("oops")
	}
}

type Name struct {
	FirstName string
}

type address struct {
	Line1 string
	Line2 string
	Line3 string
}

type Custumer struct {
	FirstName Name
	Family    string
	Address   address
	Dob       time.Time
}
