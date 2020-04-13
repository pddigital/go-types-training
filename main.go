package main

import (
	"fmt"
	"go-data-types/organization"
)

func main() {
	p := organization.NewPerson("Paul", "Day", organization.NewEuropeanUnionIdentifier("123-45-7890", "Germany"))

	err := p.SetTwitterHandler("@pddigital")

	fmt.Printf("%T\n", organization.TwitterHandler("test"))

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s \n", err)
	}

	println(p.FullName())
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
	println(p.ID())
	println(p.Country())

}
