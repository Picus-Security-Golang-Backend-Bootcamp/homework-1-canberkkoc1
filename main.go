package main

import (
	"fmt"
	"helper/helper"
	"os"
	"strings"
)

func main() {

	bookSlice := []string{"Pale Fire", "Antigone", "As I Lay Dying", "The Color Purple", "The Possessed", "Gone With the Wind", "Lord of the Flies"}

	// read command line
	args := os.Args

	// get first argument lowercase

	firstArg := helper.Lowercase(args[1])

	if len(args) == 1 || (firstArg != "list" && firstArg != "search") {

		fmt.Println("Please enter 'search' <Book_name> or 'list' ")

	} else {

		if firstArg == "list" {

			for i := 0; i < len(bookSlice); i++ {

				fmt.Println(bookSlice[i])
			}

		}

		if firstArg == "search" {

			// get book name

			secondArg := strings.Join(args[2:], " ")

			secondArg = helper.Lowercase(secondArg)

			var result bool = false

			for _, v := range bookSlice {

				if helper.Lowercase(v) == secondArg {
					result = true
					break
				}

			}

			if result {
				fmt.Printf("Book Found: %v", secondArg)
			} else {
				fmt.Printf("Book did not Found: %v", secondArg)

			}
		}

	}

}
