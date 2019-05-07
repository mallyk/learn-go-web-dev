package main

import "fmt"

type person struct {
	fName string
	lName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("I am", p.fName, p.lName)
}

func (sa secretAgent) saSpeak() {
	fmt.Println(sa.lName, ",", sa.fName, sa.lName)
}

func main() {
	p1 := person{"Miss", "Moneypenny"}
	sa1 := secretAgent{
		person:        person{"James", "Bond"},
		licenseToKill: true,
	}

	fmt.Println(p1.fName)
	p1.pSpeak()

	fmt.Println(sa1.licenseToKill)
	sa1.saSpeak()
	sa1.pSpeak()
}
