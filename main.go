package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Hello welcome to go conference here in 2023")
	arraysName := []string{"John Doe", "Jack man ", "Bruzlli Khan"}
	result, result2 := PrintSum(arraysName)
	fmt.Println("Surname is ", result, result2)
	// initalizing the variable for the while loops here
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++

	}
}

func PrintSum(a []string) (c string, b []string) {
	holdSurname := ""
	firstNames := []string{}
	for _, v := range a {
		firstName := strings.Fields(v) //we used fields here because we want to split the string on white spaces
		firstNames = append(firstNames, firstName[0])
		holdSurname += string(v[0])
	}
	return holdSurname, firstNames
}
