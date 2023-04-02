package main

import (
	"fmt"
)

func main(){
	//1.define
	myArray := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("%v\n", myArray)

	//2.loop
	for index,value := range myArray {
		switch value {
		case "stupid":
			myArray[index] = "smart"
		case "weak":
			myArray[index] = "strong"
		}
	}
	fmt.Printf("%v\n", myArray)
}