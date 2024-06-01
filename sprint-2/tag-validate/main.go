package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

func Str2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return num
}

func Validate(object any) bool {
	const countValueTag = 2

	obj := reflect.ValueOf(object)
	objType := obj.Type()

	for i := range objType.NumField() {
		val, ok := obj.Field(i).Interface().(int)
		if !ok {
			continue
		}

		objField := objType.Field(i)
		tagValue, isLimit := objField.Tag.Lookup("limit")
		if !isLimit {
			continue
		}

		tags := strings.Split(tagValue, ",")

		if len(tags) > 2 || len(tags) == 0 {
			return false
		}

		if val < Str2Int(tags[0]) {
			return false
		}

		if len(tags) == countValueTag {
			if val > Str2Int(tags[1]) {
				return false
			}
		}
	}

	return true
}

func main() {
	user := User{"admin", 20, 88}
	log.Println(Validate(user))
}
