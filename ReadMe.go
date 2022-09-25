package main

import (
	"fmt"
	"os"
)

func (d Dir) ReadMe() {
	fmt.Println("")
	fmt.Println("README FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(string(d))
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(files[8].Name())
	readme, err := os.Open("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fil, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readme)
	fmt.Println(fil)

}

func (sd StrDir) ReadMe() {
	fmt.Println("")
	fmt.Println("README FILE CHECK STARTS HERE:")
	files, err := os.ReadDir(sd.strDirect)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(files[8].Name())
	readme, err := os.Open("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fil, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readme)
	fmt.Println(fil)

}
