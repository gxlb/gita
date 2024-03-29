package gitlib

import (
	"encoding/json"
	"fmt"
)

// PrettyShow show something pretty
func PrettyShow(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err, v)
	}
	fmt.Println(string(b))
}
