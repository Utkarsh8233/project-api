package main

import "fmt"

func main() {
	hobbies := [3]string{"football", "singing", "dancing"}
	first_element_hobbies := hobbies[0]
	fmt.Println("Printing the first element: ", first_element_hobbies)
	fmt.Println("Printing the second and the third element: ", hobbies[1:])
	fmt.Println("Printing the first and second element: ", hobbies[:2])
	//using another method to extract element 1 and 2:
	hobby_slice := hobbies[0:2]
	fmt.Println("Printing the first and the second element: ", hobby_slice)
	hobby_slice[0] = hobbies[1]
	hobby_slice[1] = hobbies[2]
	fmt.Println("Prnting the manipulated slice: ", hobby_slice)
	//dynamic array
	goals := []string{"Car", "Home"}
	fmt.Println("My goals: ", goals)
	goals[1] = "Restaurant"
	fmt.Println("Modified goals: ", goals)
	goals = append(goals, "Home")
	fmt.Println("Adding third goal (appending)", goals)

	//append 2 arrays
	array1 := []string{"lenovo", "apple", "hp"}
	array2 := []string{"sony", "dell", "asus"}
	newArray := append(array1, array2...)
	fmt.Println("Printing array after appending two arrays: ", newArray)

	//structs
	type Product struct {
		product string
		id      int
		price   float64
	}

	products := []Product{
		{
			"Car",
			1,
			500000,
		},
		{
			"Bike",
			2,
			200000,
		},
	}
	fmt.Println("Printing struct of product: ", products)
	newProduct := Product{
		"Laptop",
		3,
		50000,
	}
	products = append(products, newProduct)
	fmt.Println("Printing struct after appending: ", products)
}
