package main

import "fmt"

// This file take a look at 'labels' in Golang and how it is used for control flow in Golang applications

func InitControlFlowExample() {
begin: // this is called a 'label' which is not available in other languages
	var start, end int
	var confirmation string
	fmt.Println("Enter start number")
	fmt.Scanln(&start)
	fmt.Println("Enter end number")
	fmt.Scanln(&end)
	fmt.Printf("The numbers entered are %d and %d\n", start, end)

	fmt.Println("Running for loop with the given ranges")
	for i := start; i <= end; i++ {
		fmt.Println("Iteration number: ", i)
	}

	fmt.Println("Do you want to choose again?")
	fmt.Scanln(&confirmation)
	if confirmation == "Y" || confirmation == "Yes" {
		goto begin
	} else {
		fmt.Println("Invalid input received, ending the loop")
	}
}
