package main

// this could be done with import "package" line by line, but why type that out all the time? 
// use a "factored" import statement
import (
	"fmt"
	"math"
)

func main () {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
