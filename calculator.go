package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Dog struct {
	Name string
	Age  int
}

type Cat struct {
	Name string
	Age  int
}

func (h Human) GetName() string {
	return h.Name
}

func (d Dog) GetName() string {
	return d.Name
}

func (c Cat) GetName() string {
	return c.Name
}

type Namer interface {
	GetName() string
}

func (h Human) SayHello() string {
	return fmt.Sprintf("hello %s", h.Name)
}

func (h *Human) AddAge() {
	h.Age = h.Age + 1
}

func Test(item Namer) {
	fmt.Println("Погладил", item.GetName())
}
