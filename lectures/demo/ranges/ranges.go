package main

import "fmt"

func main() {
	slice := []string{"one", "uno", "emu", "!"}

	for i, element := range slice {
		fmt.Println(i, element, ":")

		for _, alp := range element {
			fmt.Printf("   %q\n",alp)
		}
	}
}
