package ChannelSync

import (

	"fmt"
)

type Salutation struct {
	Name string
	Greeting string

}

type Salutations []Salutation

type Printer func(string)()

func Greeting(){

	salutations := Salutations{
		{"Mary", "Hi"},
		{"Joe","Howdy"},
		{"Nancy","What'up"},
	}
	//slice = append(slice[:1], slice[2:]... )
	salutations.Greet( CreatePrintFunction("?"),  6)

}
func CreatePrintFunction(custom string) Printer {
	return func(s string){
		fmt.Println(s + custom)
	}

}

func (salutations Salutations) Greet( do Printer, times int){
	for _,s := range salutations{
		message := s.Name + " " + s.Greeting
		do(message)
	}

}
