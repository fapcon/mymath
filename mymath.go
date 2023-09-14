package main

import (
	"fmt"
	"math"
)

func main() {
	x := 4.0
	y := 2.0
	Abs := Abs(x)
	Max := Max(x, y)
	Sqrt := Sqrt(x)
	Yn := Yn(int(x), y)
	fmt.Println(Abs)
	fmt.Println(Max)
	fmt.Println(Sqrt)
	fmt.Println(Yn)
}
func Sqrt(x float64) float64 {
	math.Sqrt(x)
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	math.Ceil(x)
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	math.Floor(x)
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	math.Pow(x, y)
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	math.Max(x, y)
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	math.Min(x, y)
	return math.Min(x, y)
}
func Abs(x float64) float64 {
	math.Abs(x)
	return math.Abs(x)
}
func Yn(x int, y float64) float64 {
	math.Yn(x, y)
	return math.Yn(x, y)
}
