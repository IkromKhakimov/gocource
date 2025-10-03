package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type PersonJ struct {
	FirstName    string   `json:"first_name"`
	Age          int      `json:"age,omitempty"`
	EmailAddress string   `json:"email"`
	Address      AddressJ `json:"address"`
}

type AddressJ struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := PersonJ{FirstName: "John", EmailAddress: "sample.@emample.com"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := PersonJ{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakeemail.com", Address: AddressJ{City: "New York", State: "State"}}

	jsondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsondata1))

	jsonData2 := `{
		"full_name": "Jenny Doe",
		"emp_id": "0008",
		"age": 30,
		"address": {
			"city": "San Jose",
			"state": "CA"
		}}`

	var employeeFromJson EmployeeJ
	json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	fmt.Println(employeeFromJson.Age + 5)

	listOfCityState := []AddressJ{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "San Jose", State: "CA"},
		{City: "San Jose", State: "CA"},
		{City: "San Jose", State: "CA"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to JSON:", err)
	}

	fmt.Println("JSON List:", string(jsonList))
	fmt.Println("JSON List:", jsonList)

	// Handling unknown json structures
	jsonData3 := `{"name": "John", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
}

type EmployeeJ struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
