package jsonf

import (
	"encoding/json"
	"fmt"
)

// your JSON structure as a byte slice
var j = []byte(`{"foo":1,"bar":2,"baz":[3,4]}`)

// Jsonf :
func Jsonf(instr []byte) {

	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(instr, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		fmt.Println(k[i])
		i++
	}

	// output result to STDOUT
	for _, arr := range k {
		fmt.Println(arr)
	}
	fmt.Printf("%#v\n", k)

}
