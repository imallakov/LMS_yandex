package main

type Employee struct {
	name string
	position string
	salary int
	bonus int
}

func (p Employee) CalculateTotalSalary() int{
	return p.salary+p.bonus
}