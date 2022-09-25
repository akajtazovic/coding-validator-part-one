package main

import (
	"fmt"
	"os"
	"strings"
)

func (d Dir) GoFileCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go file"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
}

func (sd StrDir) GoFileCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go file"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file: ", gofile, " ", text)
	}
}
