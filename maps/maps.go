package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"FaceBook": "https://www.facebook.com",
	}
	fmt.Println("Printing whole map", websites)
	fmt.Println("Printing value stored for the key Google:", websites["Google"])
	//add a new key value pair:
	websites["AWS"] = "https://www.aws.com"
	fmt.Println("Printing whole map after adding a new pair: ", websites)
	//delete a pair
	delete(websites, "AWS")
	fmt.Println("Printing the map after deleting AWS key value pair: ", websites)
}
