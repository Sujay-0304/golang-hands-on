package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Animal struct {
	XMLName     xml.Name `xml:"animal" json:"-"`
	AnimalName  string   `xml:"animal_name" json:"animal_name"`
	AnimalBreed string   `xml:"AnimalBreed" json:"AnimalBreed"`
}


type Personal struct {
	XMLName       xml.Name `xml:"personal" json:"-"`
	FavouriteFood string   `xml:"FavouriteFood" json:"FavouriteFood"`
	Shop          string   `xml:"shop" json:"shop"`
}

type Pet struct {
	XMLName  xml.Name `xml:"pet" json:"-"`
	Animal   Animal   `xml:"animal" json:"animal"`
	Personal Personal `xml:"personal" json:"personal"`
}

type Pets struct {
	XMLName xml.Name `xml:"pets" json:"-"`
	Pet     []Pet    `xml:"pet" json:"pet"`
}

func main() {
	pets := Pets{
		Pet: []Pet{
			{
				Animal: Animal{AnimalName: "Dog", AnimalBreed: "Labrador"},
				Personal: Personal{
					FavouriteFood: "Bones",
					Shop:          "Pet Store",
				},
			},
		},
	}


	jsonData, _ := json.MarshalIndent(pets, "", "  ")
	fmt.Println("JSON Output:")
	fmt.Println(string(jsonData))

	xmlData, _ := xml.MarshalIndent(pets, "", "  ")
	fmt.Println("\nXML Output:")
	fmt.Println(string(xmlData))
}
