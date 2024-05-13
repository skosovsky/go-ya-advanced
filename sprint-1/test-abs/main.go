package main

import (
	"fmt"
	"math"
)

func main() {
	v := Abs(3)    //nolint:mnd // example
	fmt.Println(v) //nolint:forbidigo // example
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
func Abs(value float64) float64 {
	return math.Abs(value)
}
