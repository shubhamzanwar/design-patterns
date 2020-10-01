package main

// Gender describes and handles everything related to employee gender
type Gender interface {
	describePerson() string
}

// Department describes and handles everything related
// to a particular function in the org
type Department interface {
	getDepartmentName() string
}
