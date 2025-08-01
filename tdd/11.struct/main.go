package main

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary int
}

func NewPerson(name string, age int, job string, salary int) Person {
	return Person{
		Name:   name,
		Age:    age,
		Job:    job,
		Salary: salary,
	}
}
