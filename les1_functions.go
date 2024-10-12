package main

import (
	"errors"
	"fmt"
	"math"
)

var globalStr string = "Global Var"

const pi = math.Pi

func main() {
	var ph string = "check"
	a1 := 1
	a2 := "str str"

	fmt.Println(ph)
	fmt.Println(globalStr)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(pi)

	circleRadius := 2
	link := &circleRadius
	*link++
	circleRadius *= 2
	println(*link)

	iterate(link)

	fmt.Printf("Radius: %d \n", circleRadius)

	radius := float32(circleRadius)

	circleArea, err := calcCircleArea(radius)
	if err == nil {
		fmt.Printf("Area: %f32 \n", circleArea)
	}

	printCircleArea(radius)

	nRadius, nArea := circleAreaData(radius)

	fmt.Println(nRadius, nArea)

	var p *int
	println(p)
	fmt.Println(p)

	fFun := func(a int) {
		println(a)
	}

	fFun(3)

	a := 3
	b := 5

	println(pryamougolnikPerimetr(a, b))

	var prArea int = pryamougolnicArea(a, b)
	println(prArea)
	if prArea < 10 {
		println("Small")
	} else if 10 < prArea && prArea < 15 {
		println("Medium")
	} else {
		println("Large")
	}
}

func iterate(a *int) {
	*a++
}

func circleAreaData(radius float32) (float32, float32) {
	area, err := calcCircleArea(radius)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return radius, float32(0)
	}
	return radius, area
}

func calcCircleArea(radius float32) (float32, error) {
	if radius < 0 {
		return float32(0), errors.New("incorrect radius value")
	}

	return radius * radius * pi, nil
}

func printCircleArea(radius float32) {
	area, err := calcCircleArea(radius)
	if err == nil {
		fmt.Printf("Area: %f32 \n", area)
	}
}

func pryamougolnicArea(a int, b int) int {
	return a * b
}

func pryamougolnikPerimetr(a int, b int) int {
	return a*2 + b*2
}
