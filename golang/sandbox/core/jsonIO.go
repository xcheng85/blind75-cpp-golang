package core

import "fmt"

var TestJson = `
{
	"First": "Giannis",
	"Last": "Antetokounmpo",
	"Dob": {
		"year": 1994, 
		"month": 12,
		"day": 6
	}
}

`

func Probe(js map[string]interface{}) {
	for k, v := range js {
		// type assertion, then recursive
		// dfs
		obj, ok := v.(map[string]interface{})
		if ok {
			fmt.Println(obj)
			Probe(obj)
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func TypeSwitch(js map[string]interface{}) {
	for k, v := range js {
		// type switch syntax
		switch c := v.(type) {
		case string:
			fmt.Printf("%v field is string %v\n", k, c)
		case float32:
			fmt.Printf("%v field is float32 %v\n", k, c)
		case float64:
			fmt.Printf("%v field is float64 %v\n", k, c)
		case int64:
			fmt.Printf("%v field is int64 %v\n", k, c)
		case int:
			fmt.Printf("%v field is int %v\n", k, c)
		case int32:
			fmt.Printf("%v field is int32 %v\n", k, c)
		case map[string]interface{}:
			// type assert from interface{} to map
			TypeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}
