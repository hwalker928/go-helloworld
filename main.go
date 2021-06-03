package main

// Imports the required packages
import (
	// fmt is used for writing to the console
	"fmt"
	// errors is used for logging errors
	"errors"
	// math for sqrt function
	"math"
)

func main() {

	fmt.Println("------------")

	// Declares the variable x as an integer with the value of 5
	var x int = 5
	// Declares the variable y as an integer with the value of 7
	// We can use := as go can detect the type of variable for us
	y := 7
	// Adds the variables x and y together and declares it as z
	var z int = x + y
	// Outputs the value of z to the console
	fmt.Println(z)

	fmt.Println("------------")

	// If variable z is greater than y, then
	// The condition does NOT need brackets
	if z > y {
		// Output information to console
		fmt.Println("z is bigger than y.")
		// Else if statement
	} else if y == 5 {
		fmt.Println("y is equal to 5.")
		// Else statement
	} else {
		fmt.Println("No conditions matched.")
	}

	fmt.Println("------------")

	// Declares the variable a as an integer array which can hold 5 values
	// All integers by default are 0 - unless specified
	var a [5]int
	fmt.Println(a)
	// Sets the 3rd value of the array a to 7
	a[2] = 7
	fmt.Println(a)

	// Creates an integer array, which can hold 4 values, called b
	// The values are predefined by the {4, 3, 2, 1}
	// The length of the array is part of the type, meaning the number of elements is fixed
	b := [4]int{4, 3, 2, 1}
	fmt.Println(b)

	// This is called a slice.
	// It is like an array, but does not have a fixed length of elements.
	c := []int{4, 3, 2, 1}
	// Adds 13 to the end of the slice array, as it is not fixed.
	c = append(c, 13)
	fmt.Println(c)

	fmt.Println("------------")

	// Defines the map variable as vertices
	vertices := make(map[string]int)
	// Adds these following values to the map
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["hexagon"] = 6
	// Removes the value of hexagon from the map
	delete(vertices, "hexagon")
	// Outputs the map as a whole
	fmt.Println(vertices)
	// Outputs only the value of square
	fmt.Println(vertices["square"])

	fmt.Println("------------")

	// Starting at 0, looping 5 times in total, increasing i
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Same as above code, just using a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------------")

	// Creates the array/slice with predefined strings
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Using a map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println("------------")

	// Using our function defined below
	fmt.Println(sum(2, 3))
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("------------")

	// Uses the struct of person, with the values Harry and 14
	p := person{name: "Harry", age: 14}
	// These values can be accessed by p.name and p.age
	fmt.Println("Struct:", p, "Name:", p.name, "Age:", p.age)

	fmt.Println("------------")

	d := 7
	// Gives the variable to the function
	inc(&d)
	// Outputs the memory value of d
	fmt.Println("Mem val:", &d, "True val:", d)
}

// Takes 2 ints, and returns an int
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

// Creates a struct, with name as a string and age as an it
type person struct {
	name string
	age  int
}

// Uses the actual variable, not the value!
func inc(x *int) {
	*x++
}
