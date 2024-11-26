package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoderStream(t *testing.T) {
	
	create, _ := os.Create("sample_output.json")
	encoder := json.NewEncoder(create)

	customer := Customer{
		FirstName: "Muhamad",
		MiddleName: "Habibi",
		LastName: "Cumi",
	}
	_ = encoder.Encode(customer)

	fmt.Println(customer)

}