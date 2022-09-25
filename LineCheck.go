package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func lineNumber(l int) error {
	if l > 100 {
		return errors.New("line number too long")
	}
	return nil
}

func (d Dir) LineCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go File"
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
	error := 0
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("total lines: ", lineCount)
		lc := lineNumber(lineCount)
		fmt.Println("LINE CHECK:")
		fmt.Println("")
		if lc == nil {
			fmt.Println("PASS")
		}
		if lc != nil {
			error++
			pc, filename, line, _ := runtime.Caller(1)
			log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("FAIL")
		}
	}
	fmt.Println("total errors: ", error)
}

func (sd StrDir) LineCheck() {
	fmt.Println("GO FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a go File"
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
	error := 0
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("total lines: ", lineCount)
		lc := lineNumber(lineCount)
		fmt.Println("LINE CHECK STARTS HERE")
		fmt.Println("")
		if lc == nil {
			fmt.Println("")
			fmt.Println("PASS")
		}
		if lc != nil {
			error++
			pc, filename, line, _ := runtime.Caller(1)
			log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
		}
	}
	fmt.Println("total errors: ", error)
}
