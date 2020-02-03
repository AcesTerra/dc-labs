// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main
//package geometry

import (
	"math"
	"math/rand"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Point struct{
	x, y float64
}

func (p *Point) SetX(x float64) {
 	p.x = x
}

func (p Point) GetX() float64 {
 	return p.x
}

func (p *Point) SetY(y float64) {
        p.y = y
}

func (p Point) GetY() float64 {
        return p.y
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.GetX()-p.GetX(), q.GetY()-p.GetY())
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.GetX()-p.GetX(), q.GetY()-p.GetY())
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	//fmt.Println("Flag")
	for i:= range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	//fmt.Println(sum)
	return sum
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func main(){
        sides := 0
	rand.Seed(time.Now().UnixNano())
	if s, err := strconv.Atoi(os.Args[1]); err == nil {
		sides = s
	}
	fmt.Printf("Generating a %v sides figure", sides)
	fmt.Println("Figure's vertices:")
	//fmt.Printf("%T, %v", s, s)
	//fmt.Println(sides)
	var path Path
	for i := 0; i < sides; i++{
		//fmt.Println(i)
		randomNumbers := []float64{}
		randomNumbers = randFloats(-100, 100, 2)
		//x := rand.Intn(101)
		//y := rand.Intn(101)
		//fmt.Println(randomNumbers[1])
		var randP = Point{}
		randP.SetX(randomNumbers[0])
		randP.SetY(randomNumbers[1])
		fmt.Println(randP)
        	path = append(path, randP)
		//path := Path{points}
		//path = append(path, points)
		//fmt.Println(path.Distance())
		//fmt.Println(d)
	}
	fmt.Println("Figure's perimeter:")
	fmt.Println(path.Distance())
}

//!-path
