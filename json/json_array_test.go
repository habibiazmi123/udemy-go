package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Muhamad",
		MiddleName: "Habibi",
		LastName: "Azmi",
		Hobbies: []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhamad","MiddleName":"Habibi","LastName":"Azmi","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if(err != nil) {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Azmi",
		Addresses: []Address{
			{
				Street: "Antapani", Country: "Bandung", PostalCode: "40291",
			},
			{
				Street: "Antapani 11", Country: "Bandung", PostalCode: "40292",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}