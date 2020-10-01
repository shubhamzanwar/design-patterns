package main

import "fmt"

// Female encapsulates all traits of a female employee
type Female struct {
	name       string
	age        int
	department Department
}

func (f *Female) describePerson() string {
	description := fmt.Sprintf("%s is a %d years old woman\n", f.name, f.age)
	description = description + fmt.Sprintf("She works in the %s department", f.department.getDepartmentName())
	return description
}
