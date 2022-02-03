package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the order system")
	var cart []string
	var reviews []string
	for {

		fmt.Println("Please enter 1, 2 or 3 ")
		fmt.Println("1. Add dish to your cart")
		fmt.Println("2. See my cart")
		fmt.Println("3. Leave a review")
		fmt.Println("4. See my reviews")
		fmt.Println("--------------------")
		var n int
		fmt.Scan(&n)

		if n == 1 {
			fmt.Println("1. Mastery Fied Rice")
			fmt.Println("2. Famous Sauced Ribs")
			fmt.Println("3. Enter anything else to go back to main menue")
			var t1 int
			fmt.Scan(&t1)
			fmt.Println("--------------------")
			if n == 1 {
				cart = append(cart, "Mastery Fied Rice")
			} else if n == 2 {
				cart = append(cart, "Famous Sauced Ribs")
			} else {
				continue
			}
		} else if n == 2 {
			fmt.Println("Here is your cart")
			fmt.Println(cart)
			fmt.Println("--------------------")
		} else if n == 3 {
			fmt.Println("Plase enter your review")
			var review string
			fmt.Scan(&review)
			reviews = append(reviews, review)
			fmt.Println("--------------------")
			continue
		} else if n == 4 {
			fmt.Println("Here is your reviews")
			fmt.Println(reviews)
			fmt.Println("--------------------")
		}

	}

}
