package main

import (
	"bufio"
	"fmt"
	"os"
)

// Validation Interface
type Validation interface {
	Lic()
	ReadMe()
	GoFileCheck()
	LineCheck()
	CommentCheck()
}

type Dir string
type StrDir struct {
	strDirect string
}

func rootdir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

// main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var userinput = ""

	for userinput != "val" {
		scanner.Scan()
		userinput = scanner.Text()
		fmt.Println("In: ", userinput)
		if userinput == "val detail" {
			fmt.Println("DETAILS: ")
			fmt.Println("This is the License Validation")
			var v Validation
			v = Dir(rootdir())
			fmt.Println("This is the go file Validation")
			v.GoFileCheck()
			fmt.Println("This is the license validation")
			v.Lic()
			fmt.Println("This is the ReadMe Validation")
			v.ReadMe()
			fmt.Println("This is the Comment Check Validation")
			fmt.Println("If this section contains at least one comment, it is validated")
			v.CommentCheck()
			fmt.Println("This is the Line Check Validation")
			v.LineCheck()
		}
		if userinput == "val help" {
			fmt.Println("Help:")
			fmt.Println("")
			fmt.Println("Lic() checks if you have a license file.")
			fmt.Println("ReadMe() checks if you have a readme file.")
			fmt.Println("LineCheck() checks if your line characters are under 100.")
			fmt.Println("CommentCheck() checks if you have comments")
			fmt.Println("GoFileCheck() checks if your file is a .go file")
			fmt.Println("")
			fmt.Println("You can find a summary of each validation after it is completed,")
			fmt.Println("or you can use val detail for a summary of the total errors,")
			fmt.Println("and where they are located.")
		}
		if userinput == "val" {
			var v Validation
			dir := Dir(rootdir())
			v = dir
			v.GoFileCheck()
			v.Lic()
			v.LineCheck()
			v.CommentCheck()
			v.ReadMe()
			str := StrDir{rootdir()}
			v = str
			v.GoFileCheck()
			v.Lic()
			v.LineCheck()
			v.CommentCheck()
			v.ReadMe()
		}
	}
}
