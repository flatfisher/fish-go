package main

import (
	"fmt"
)

type Book struct {
	Name       string
	Ingredient string
}

func (b Book) read() {
	fmt.Println(b.Name)
}

type Cook interface {
	cook()
}

func (b Book) cook() {
	fmt.Printf("Cooked %s\n", b.Ingredient)
}

func main() {
	b := Book{Name: "Learing Go", Ingredient: "fish"}
	b.read()
	b.cook()

	var c Cook = Book{Name: "Learing Go", Ingredient: "fish"}
	c.cook()
}
