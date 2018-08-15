package main

import (
	"fmt"
)

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op func(float64) float64) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}

func Substract(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc - n
	}
}

func Multiply(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc * n
	}
}

func Divide(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc / n
	}
}

func main() {
	var c Calculator
	fmt.Printf("The sum is %d", c.Do(Add(5)))
	fmt.Printf("The substraction is %d", c.Do(Substract(1)))
	fmt.Printf("The multiplication is %d", c.Do(Multiply(2)))
	fmt.Printf("The division is %d", c.Do(Divide(4)))
}
