package main

import "fmt"

type Human struct {
	Name string
	Age int
}
type worker struct {
	Human
	Company string
}
type  doctor struct {
	worker
	Stack[]string
}
type  Driver struct {
	worker
	DrivingExperience string
}
type  Policeman struct {
	worker
	Rank string
}
type  Teacher struct {
	worker
	Degree string
}
type  Firefighter struct {
	worker
	Rank2 string
}

func main(){
	prof1 := Firefighter{Rank2: "Дежурный"}
	human1 := Human{"Петя", 49}
	comp := worker{Company: "Рога и копыта"}
	human2 := Human{Name: "Ираклий", Age: 25}
	prof2:= Driver{DrivingExperience: "Опытный Драйвер"}
	human3 := Human{"Иван", 35}
	p:=&human3
	p.Name = "Анатолий"
	p.Age = 41

	fmt.Println(human1.Name,human1.Age,prof1.Rank2)
	fmt.Println(human3.Age,human3.Name)
	fmt.Println(human2.Name,human2.Age,prof2.DrivingExperience)
	fmt.Println(comp.Company)
	fmt.Println(human3.Name)
	fmt.Println(prof2.DrivingExperience)
}




