package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func newPerson(name string) *person {
	p := person{Name: name}
	p.Age = 42

	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{Name: "Alice", Age: 30})
	fmt.Println(person{Name: "Fred"})
	fmt.Println(&person{Name: "Ann", Age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	sp := &s
	fmt.Println(sp.Age)

	sp.Age = 51
	fmt.Println(sp.Age)

	dog := struct {
		Name   string
		isGood bool
	}{
		Name:   "Fido",
		isGood: true,
	}

	fmt.Println(dog)
}
