package main

import (
	"log"
	"reflect"
	"strings"
	"time"
)

type User struct {
	User      string    `json:"user,omitempty" example:"Doe"`
	CreatedAt time.Time `json:"createdAt"`
}

const (
	targetTag   = "json"
	targetField = "User"
)

func main() {
	var obj User

	objType := reflect.TypeOf(obj)
	log.Println(objType) // main.User

	field, ok := objType.FieldByName(targetField)
	if !ok {
		log.Println("field not found: ", targetField)

		return
	}

	log.Println(field) // {User  string json:"user,omitempty" example:"Doe" 0 [0] false}

	tagValue, ok := field.Tag.Lookup(targetTag)
	if !ok {
		log.Println("tag not found: ", targetTag)
	}

	log.Println("tag:", tagValue)             // tag: user,omitempty
	log.Println("tags:", field.Tag)           // tags: json:"user,omitempty" example:"Doe"
	log.Println(strings.Split(tagValue, ",")) // [user omitempty]

	for i := range objType.NumField() {
		objField := objType.Field(i)

		log.Println(objField.Tag, objField.Tag.Get("json"), objField.Name)
		// json:"user,omitempty" example:"Doe" // user,omitempty // User
		// json:"createdAt" // createdAt // CreatedAt
	}
}
