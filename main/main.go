package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Animal struct {
	NoLegs int	`json:{noLegs}`
	Name string
}

func (a Animal) CanTalk() bool {
	return false
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal{%d %s}", a.NoLegs, a.Name)
}

type Talker interface {
	CanTalk() bool
}

func main() {

	var creature Talker
	creature = Animal{4, "Pig"}
	fmt.Println(creature)
	//_ -ignores the variable
	values, err := ioutil.ReadFile("main/animals.json")

	if(err != nil){
		fmt.Println("cannot open the file")
		panic(err)
	}

	// You have to convert the json to string
	fmt.Println(string(values));

	var animal []Animal // Slice, not array
	err = json.Unmarshal(values, &animal)
	if(err != nil){
		fmt.Println("cannot open the file")
		panic(err)
	}

	//Check values after deserialization
	fmt.Println(animal)

	serialized, err := json.Marshal(animal)
	if(err != nil){
		fmt.Println("cannot open the file")
		panic(err)
	}
	fmt.Println(string(serialized))

}
