package main

type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func (e Employee) String() string {
	return e.FirstName + " " + e.LastName
}
