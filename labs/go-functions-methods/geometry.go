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
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
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
        sides := []Point{}
        fmt.Println(sides)
	rand.Seed(time.Now().UnixNano())
	if s, err := strconv.Atoi(os.Args[1]); err == nil {
		//fmt.Printf("%T, %v", s, s)
		for i := 0; i < s; i++{
			fmt.Println(i)
			randomNumbers := []float64{}
			randomNumbers = randFloats(-100, 100, 2)
			//x := rand.Intn(101)
			//y := rand.Intn(101)
			//fmt.Println(randomNumbers[1])
			var p = Point{randomNumbers[0], randomNumbers[1]}
			fmt.Println(p)
        		sides = append(sides, p)
		}
	}
	fmt.Println(sides)
}

//!-path
