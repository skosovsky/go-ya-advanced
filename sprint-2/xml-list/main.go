package main

import (
	"encoding/xml"
	"log"
)

type (
	List struct {
		Persons []Person `xml:"Person"`
	}

	Person struct {
		ID     string   `xml:"id,attr"`
		Name   string   `xml:"Name"`
		Phones []string `xml:"Phones>Phone"`
		Email  string   `xml:"Email,omitempty"`
	}
)

func main() {
	var list List

	data := `
    <List>
        <Person id="1">
            <Name>Carla Mitchel</Name>
            <Phones>
                <Phone>123-45-67</Phone>
                <Phone>890-12-34</Phone>
            </Phones>
        </Person>
        <Person id="2">
            <Name>Michael Smith</Name>
            <Email>msmith@example.com</Email>
        </Person>
    </List>
`

	err := xml.Unmarshal([]byte(data), &list)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range list.Persons {
		log.Println(person.ID, person.Name, person.Email, person.Phones)
	}
}
