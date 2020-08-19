package main

//type Employee struct {
//	Id        int
//	Name      string
//	Address   string
//	DoB       time.Time
//	Position  string
//	Salary    int
//	ManagerId int
//}

type tree struct {
	value int
	left, right *tree
}