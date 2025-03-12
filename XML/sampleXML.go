package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Pet struct
// pets
// 	pet
// 	    animal
// 			animal_name
// 			AnimalBreed
// 		animal
// 		personal
// 			FavouriteFood
// 			shop
// 		personal
// 	pet
// pets

type Animal struct {
	XMLName     xml.Name `xml:"animal"`
	AnimalName  string   `xml:"animal_name"`
	AnimalBreed string   `xml:"AnimalBreed"`
}

type Personal struct {
	XMLName       xml.Name `xml:"personal"`
	FavouriteFood string   `xml:"FavouriteFood"`
	Shop          string   `xml:"shop"`
}

type Pet struct {
	XMLName  xml.Name `xml:"pet"`
	Animal   Animal   `xml:"animal"`
	Personal Personal `xml:"personal"`
}

type Pets struct {
	XMLName xml.Name `xml:"pets"`
	Pet     []Pet    `xml:"pet"`
}


func main() {
	writeXML()
	// ReadXML()
}



func writeXML(){
	pets := Pets{
		Pet: []Pet{
			{
				Animal: Animal{
					AnimalName:  "Dog",
					AnimalBreed: "Labrador",
				},

				Personal: Personal{
					FavouriteFood: "Pedigree",
					// Shop:          "PetShop",
				},
			},
			{
				Animal: Animal{
					AnimalName:  "cat",
					// AnimalBreed: "Labrador",
				},

				Personal: Personal{
					FavouriteFood: "Whiskers",
					// Shop:          "PetShop",
				},
			},
		},
	}
	fmt.Println(pets)
	XMLdata,err := xml.MarshalIndent(pets,"","   ")
	if err != nil {
		fmt.Println("error in marshalling")
		return 
	} 
	fmt.Println(string(XMLdata))

	for ind, pet := range pets.Pet {
		fmt.Println("Pet: ", ind+1)
		fmt.Println("Animal Name: ", pet.Animal.AnimalName)
		fmt.Println("Animal Breed: ", pet.Animal.AnimalBreed)
		fmt.Println("Favourite Food: ", pet.Personal.FavouriteFood)
		fmt.Println("Shop: ", pet.Personal.Shop)
	}

	// Writing to file
	writeErr := os.WriteFile("pets.xml", XMLdata, 0644) 
	if writeErr != nil {
		fmt.Println("Error in writing to file")
		return
	}
	fmt.Println("Data written to file")
}

func ReadXML(){
	file, err := os.ReadFile("pets.xml")
	if err != nil {
		fmt.Println("Error in reading file")
		return
	}

	var pets Pets 
	fmt.Println(pets) 

	err = xml.Unmarshal(file, &pets)
	if err != nil {
		fmt.Println("Error in unmarshalling")
		return
	}
	fmt.Println(pets)

	formatData,err := xml.MarshalIndent(pets,"","   ")
	if err != nil {
		fmt.Println("error in marshalling")
		return 
	}
	fmt.Println(string(formatData))

}
