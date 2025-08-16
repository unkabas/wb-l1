package main

import "fmt"

// структура Human
type Human struct {
	Name   string
	Gender string
	Age    int64
}

// структура Action cо встроенным структурой Human
type Action struct {
	Human
	Job string
}

// методы Human
func (h Human) PrintGender() {
	fmt.Printf("%s is a %s\n", h.Name, h.Gender)
}
func (h Human) PrintAge() {
	fmt.Printf("%s is %d years old\n", h.Name, h.Age)
}

// метод Action
func (a Action) PrintJob() {
	fmt.Printf("%s is a(an) %s\n", a.Name, a.Job)

}

// реализация
func main() {
	Peter := Action{Human: Human{Name: "Peter", Gender: "male", Age: 30}, Job: "fireman"}
	Peter.PrintGender()
	Peter.PrintAge()
	Peter.PrintJob()
}
