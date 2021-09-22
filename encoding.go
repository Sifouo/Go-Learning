package main

import (
	"encoding/json"
	"fmt"
)

func encode() {
	jsonstring := `{
		"success": true,
		"timestamp": 1588779306,
		"base": "EUR",
		"date": "2020-05-06",
		"rates": {
		"AUD": 1.683349,
		"CAD": 1.528643,
		"GBP": 0.874757,
		"SGD": 1.534513,
		"USD": 1.080054
		}
	}`

	var result map[string]interface{}

	json.Unmarshal([]byte(jsonstring), &result)

	rates := result["rates"]
	fmt.Println(rates)
	currencies := rates.(map[string]interface{})
	fmt.Println(currencies["SGD"])

}
