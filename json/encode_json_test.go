package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJSON("Azmi")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Muhamad", "Habibi", "Azmi"})
}