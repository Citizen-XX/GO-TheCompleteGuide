package main

import "fmt"

type product struct{
	title string
	id int
	price float64
}

func main() {
	hobbies := [3]string{"Programming","Gaming","Workout"}
	goals := []string{"Improving","Challenge","Learning"}
	products := []product{
		{
			title: "The Interview",
			id: 0,
			price: 6.99,
		},
		{
			title: "Walter Mitty",
			id: 1,
			price: 4.99,
		},
	}
	// var products []product = []product{
	// 	{
	// 		title: "The Interview",
	// 		id: 0,
	// 		price: 6.99,
	// 	},
	// 	{
	// 		title: "Walter Mitty",
	// 		id: 1,
	// 		price: 4.99,
	// 	},
	// }
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	someHobbies := hobbies[:2]
	fmt.Println(someHobbies)
	someHobbies = someHobbies[1:3]
	fmt.Println(someHobbies)
	fmt.Println(goals)
	goals[1] = "ForFun"
	goals=append(goals,"ThirdAdded")
	fmt.Println(goals)
	products = append(products, product{
		title: "LettersToJuliette",
		id: 2,
		price: 9.99,
	},)
	fmt.Println(products[1:])
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
