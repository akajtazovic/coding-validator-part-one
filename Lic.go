package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

func (d Dir) Lic() {
	fmt.Println("")
	fmt.Println("LICENSE FILE CHECK STARTS HERE:")
	fmt.Println("test lic", d)
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files[5].Name())
	fmt.Println("pwd:", d)
	contents, err := os.Open("LICENSE")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.ReadFile("LICENSE")
	fmt.Println(file)
	if err == nil {
		fmt.Println("PASS")
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			inside := scanner.Text()
			fmt.Println("Inside: ", inside)
		}
		if err != nil {
			pc, filename, line, _ := runtime.Caller(1)
			log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
		}
	}
}

func (sd StrDir) Lic() {
	fmt.Println("")
	fmt.Println("LICENSE FILE CHECK STARTS HERE:")
	fmt.Println("test slice", sd)
	sdi := sd.strDirect
	files, err := os.ReadDir(sdi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pwd:", sd)
	fmt.Println(files[5].Name())
	contents, err := os.Open("LICENSE")
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.ReadFile("LICENSE")
	fmt.Println(file)
	if err == nil {
		fmt.Println("PASS")
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			inside := scanner.Text()
			fmt.Println("Inside: ", inside)
		}
		if err != nil {
			pc, filename, line, _ := runtime.Caller(1)
			log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
		}
	}

}
