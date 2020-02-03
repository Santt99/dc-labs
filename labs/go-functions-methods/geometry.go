package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	sides := os.Args[1]
	nOfPoints, err := strconv.Atoi(sides)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("- Generating a [%d] sides figure\n", nOfPoints)
	points := make(Path, nOfPoints)
	fmt.Println("- Figure's verices")
	for i := 0; i < nOfPoints; i++ {
		point := Point{rand.Float64() * 100, rand.Float64() * 100}
		fmt.Printf("   - ( %f, %f )\n", point.X(), point.Y())
		points[i] = point
	}

	fmt.Println("- Figure's Perimeter")
	fmt.Printf("   -")
	distance := points.Distance()
	fmt.Printf(" = %f\n", distance)
}

type Point struct{ x, y float64 }

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

func (p Point) X() float64 {
	return p.x
}
func (p Point) Y() float64 {
	return p.y
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {

			tempDistance := path[i-1].Distance(path[i])
			if i < len(path)-1 {
				fmt.Printf(" %f +", tempDistance)
			} else {
				fmt.Printf(" %f ", tempDistance)
			}

			sum += tempDistance
		}
	}
	return sum
}
