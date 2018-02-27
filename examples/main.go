package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pupfile"
)

func main() {
	// create test
	pf := pupfile.NewPupFile()
	err := pf.Create("test.zip")
	if err != nil {
		fmt.Println("create file fail! err is", err)
	}
	const ImgFileName = "image/01.jpg"
	pf.AddEmptyPage()
	pf.SetPageImage(0, ImgFileName, readFileBytes(ImgFileName))
	pf.Close()

	// open test
	pf2 := pupfile.NewPupFile()
	err = pf2.Open("test.zip")
	if err != nil {
		fmt.Println("open file fail! err is", err)
	}
	pf2.Close()

	imgData := pf2.GetPageImage(0)
	ioutil.WriteFile("out.jpg", imgData, os.ModePerm)
}

func readFileBytes(filename string) []byte {
	bytes, _ := ioutil.ReadFile(filename)
	return bytes
}
