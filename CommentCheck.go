package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (d Dir) CommentCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
		comment := "//"
		if strings.ContainsAny(lines, comment) == true {
			fmt.Println("Comment Located")
		}
		if strings.ContainsAny(lines, comment) == false {
			fmt.Println("Comment not found")
		}
	}
}

func (sd StrDir) CommentCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File"
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
		comment := "//"
		if strings.ContainsAny(lines, comment) == true {
			fmt.Println("Comment Located")
		}
		if strings.ContainsAny(lines, comment) == false {
			fmt.Println("No comment")
		}
	}
}
