package reverse

import "fmt"
import "strings"

func RevOne(input string) string {

	splitInput := strings.Split(input, "")

	n := len(splitInput)

	revInput := make([]string, n)

	for i, r := range splitInput {
		revInput[n-1-i] = r
	}

	finalString := strings.Join(revInput, "")
	fmt.Println(finalString)
	return finalString
}

func RevTwo(input string) string {
	n := len(input)

	revInput := make([]string, n)

	for i, r := range input {
		revInput[n-1-i] = string(r)
	}

	finalString := strings.Join(revInput, "")
	fmt.Println(finalString)
	return finalString

}

func RevThree(input string) string {
	var revString string

	for _, r := range input {
		revString = string(r) + revString
	}

	fmt.Println(revString)
	return revString

}
