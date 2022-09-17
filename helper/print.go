package helper

import (
	"encoding/json"
	"fmt"
)

// PP - PrettyPrint
func PP(v any) {
	b, _ := json.MarshalIndent(v, "", "	")
	fmt.Printf("\n%s\n", string(b))
}
