package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}
type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp := [3]employee{{"Amir", 80000, "Full-Stack Developer"}, {"Ahmed", 70000, "Android Developer"}, {"Ali", 90000, "ios Developer"}}
	var comp company
	comp.companyName = "Tetra"
	comp.employees = emp[:]

	fmt.Println("Company Name", comp.companyName)
	fmt.Println("Employees", emp[:])
}
