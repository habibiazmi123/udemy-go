package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonRequest := `{"FirstName": "Muhamad", "MiddleName": "Habibi", "LastName": "Azmi", "Age": 10, "Married": true}`
	jsonBytes := []byte(jsonRequest)

	customer := &Customer{}
	json.Unmarshal(jsonBytes, customer)

	fmt.Println(customer)
}