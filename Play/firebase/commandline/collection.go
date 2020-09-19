package main

import "fmt"

func setcollection(arguments []string) {

	if len(arguments) == 0 {
		fmt.Printf("no collection set - use %q\n", "setcollection <collection>")
		return
	}
	collection = arguments[0]
}

func getcollection(arguments []string) {

	fmt.Printf("collection: %q\n", collection)
}
