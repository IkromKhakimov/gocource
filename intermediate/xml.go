package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type PersonX struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	City    string   `xml:"city,omitempty"`
	Email   string   `xml:"-"`
	Address AddressX `xml:"address"`
}

type AddressX struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := PersonX{Name: "John", Age: 30, City: "London", Email: "email@exampleemail.com", Address: AddressX{City: "Oakland", State: "CA"}}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println("XML Data:", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println("XML Data:", string(xmlData1))

	//xmlRaw := `<person><name>John</name><age>25</age></person>`
	xmlRaw := `<person><name>John</name><age>25</age><address><city>San Francisco</city><state>CA</state></address></person>`
	var personxml PersonX
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling data into XML:", err)
	}
	fmt.Println("XML Data:", personxml)
	fmt.Println("Local String:", personxml.XMLName.Local)
	fmt.Println("Namespace:", personxml.XMLName.Space)

	book := BookX{
		ISBN:   "abc123",
		Title:  "Go Bootcamp",
		Author: "Ashish",
		Pseudo: PseudoX{
			Value:      "value",
			PseudoAttr: "blue",
		},
	}
	xmldataattr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println("XML Data:", string(xmldataattr))
}

type BookX struct {
	XMLName xml.Name `xml:"book"`
	ISBN    string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	Pseudo  PseudoX
}

type PseudoX struct {
	Value      string `xml:",chardata"`
	PseudoAttr string `xml:"pseudo,attr"`
}

// <book isbn="abc123" color="blue">
