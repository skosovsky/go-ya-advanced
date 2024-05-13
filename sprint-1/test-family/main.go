package main

import (
	"errors"
	"log"
)

// Relationship определяет положение в семье.
type Relationship string

// Возможные роли в семье.
const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

// Family описывает семью.
type Family struct {
	Members map[Relationship]Person
}

// Person описывает конкретного человека в семье.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var (
	// ErrRelationshipAlreadyExists возвращает ошибку, если роль уже занята.
	ErrRelationshipAlreadyExists = errors.New("relationship already exists")
)

// AddNew добавляет нового члена семьи.
// Если в семье ещё нет людей, создаётся пустая мапа.
// Если роль уже занята, метод выдаёт ошибку.
func (f *Family) AddNew(relationship Relationship, person Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[relationship]; ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[relationship] = person

	return nil
}

func main() {
	family := Family{} //nolint:exhaustruct // empty
	err := family.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56, //nolint:mnd // example
	})
	log.Println(family, err)

	err = family.AddNew(Father, Person{
		FirstName: "Drug",
		LastName:  "Mishi",
		Age:       57, //nolint:mnd // example
	})
	log.Println(family, err)
}
