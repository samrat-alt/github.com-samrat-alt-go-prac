package main

import "fmt"

func showName(names []string) {

	for i, showname := range names {
		fmt.Printf("%d %s\n", i+1, showname)
	}

}

func main() {

	// name := "samrat pandey"
	// for i := 0; i < 20; i++ {
	// 	fmt.Printf("%d. my name is %s\n", i+1, name)
	// }

	names := []string{"John", "Emma", "Michael", "Sophia", "William", "Olivia", "James", "Amelia",
		"Benjamin", "Isabella", "Daniel", "Mia", "Alexander", "Charlotte", "Ethan", "Emily", "Jacob", "Ava", "Ryan", "Evelyn"}

	showName(names)
	// for i, name := range names {
	// 	fmt.Printf("%d %s\n", i+1, name)
	// }

}
