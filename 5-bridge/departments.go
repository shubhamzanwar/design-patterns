package main

// Developer struct for dev dept
type Developer struct{}

// PM struct for project management department
type PM struct{}

// HR struct for human resources department
type HR struct{}

func (d Developer) getDepartmentName() string {
	return "software development"
}

func (pm PM) getDepartmentName() string {
	return "product management"
}

func (hr HR) getDepartmentName() string {
	return "human resources"
}
