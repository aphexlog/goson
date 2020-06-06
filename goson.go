// GoLang styling documentation located here: https://golang.org/s/style
package main

import (
	"encoding/json"
	"fmt"
)

// Ancestor : JSON object to be parsed
type Ancestor struct {
	Species     string
	Description string
}

func main() {
	ancestorJSON := `{"species": "Orrorin","description": "The Orrorin tugenensis is a postulated early species of Homininae, estimated at 6.1 to 5.7 million years (Ma) and discovered in 2000."}`
	var ancestor Ancestor
	json.Unmarshal([]byte(ancestorJSON), &ancestor)
	fmt.Printf("Species: %s, Description: %s", ancestor.Species, ancestor.Description)
	// `&ancestor` is the address of the variable we want to store our parsed data in
}
