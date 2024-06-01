package main

import (
	"fmt"
	"log"
)

func main() {
	values := []string{"test1", "test2"}

	if ConcatSprintf(values) == ConcatPlus(values) {
		log.Println("result ok")
	}
}

func ConcatSprintf(values []string) string {
	return fmt.Sprintf("%s:%s", values[0], values[1])
}

func ConcatPlus(values []string) string {
	return values[0] + ":" + values[1]
}
