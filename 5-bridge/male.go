package main

import "fmt"

// Male encapsulates all traits of a male employee
type Male struct {
	name       string
	age        int
	department Department
}

func (m *Male) describePerson() string {
	description := fmt.Sprintf("%s is a %d years old man\n", m.name, m.age)
	description = description + fmt.Sprintf("He works in the %s department", m.department.getDepartmentName())
	return description
}
