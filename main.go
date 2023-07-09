package main

import (
	"fmt"
	"kubeconstructor/kinds"
)

func main() {
	menu()
}

func menu() {
	var res string
	fmt.Println("What type of kind you want to create?")
	fmt.Println("0. Exit")
	fmt.Println("1. Ingress")
	fmt.Println("2. Pod")

	fmt.Printf(">")
	fmt.Scanln(&res)

	switch res {
	case "0":
		fmt.Println("Exiting...")
	case "1":
		kinds.Ingress()

	}

}
