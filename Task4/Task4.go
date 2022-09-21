package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	rollNumber string
	name       string
	address    string
	subjects   []string
}

func NewStudent(rollNo string, name string, address string, subject []string) *Student {
	s := new(Student)
	s.rollNumber = rollNo
	s.name = name
	s.address = address
	s.subjects = subject

	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollNo string, name string, address string, subject []string) *Student {
	st := NewStudent(rollNo, name, address, subject)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := 0; i <= len(ls.list)-1; i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))

		fmt.Println("student", "roll no", "     ", ls.list[i].rollNumber)
		fmt.Println("student", "name", "     ", ls.list[i].name)
		fmt.Println("student", "address", "     ", ls.list[i].address)
		fmt.Println("student", "subjects", "     ", ls.list[i].subjects)
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent("24", "Asim", "AAAAAAAA", []string{"blockchain", "webdevelopment"})
	student.CreateStudent("25", "Ahmed", "BBBBBBB", []string{"devops", "spm"})
	student.CreateStudent("26", "Asad", "CCCCCCCC", []string{"deeplearning", "smd"})

	student.Print()
	calculateHash(*student)

}

func calculateHash(ls StudentList) {
	fmt.Println("Data received")
	for i := 0; i <= len(ls.list)-1; i++ {

		fmt.Println("student", i, ls.list[i])
		fmt.Println("address hash", sha256.Sum256([]byte(ls.list[i].address)))
		fmt.Println("roll no hash", sha256.Sum256([]byte(ls.list[i].rollNumber)))
		fmt.Println("name hash", sha256.Sum256([]byte(ls.list[i].name)))
		for j := 0; j <= len(ls.list[i].subjects)-1; j++ {
			fmt.Println("hash of subject", j+1, sha256.Sum256([]byte(ls.list[i].subjects[j])))
		}

	}

}
