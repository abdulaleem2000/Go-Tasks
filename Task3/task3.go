package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollNumber int
	name       string
	address    string
}

func NewStudent(rollNo int, name string, address string) *Student {
	s := new(Student)
	s.rollNumber = rollNo
	s.name = name
	s.address = address

	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollNo int, name string, address string) *Student {
	st := NewStudent(rollNo, name, address)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := 0; i <= len(ls.list)-1; i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))

		fmt.Println("student", "roll no", "     ", ls.list[i].rollNumber)
		fmt.Println("student", "name", "     ", ls.list[i].name)
		fmt.Println("student", "address", "     ", ls.list[i].address)
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAAAA")
	student.CreateStudent(25, "Ahmed", "BBBBBBB")
	student.CreateStudent(26, "Asad", "CCCCCCCC")

	student.Print()

}
