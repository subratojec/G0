package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Go BANK")
	fmt.Println("What do you want to do ?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice:")
	fmt.Scan(&choice)

	fmt.Println("You entered : ", choice)

}
