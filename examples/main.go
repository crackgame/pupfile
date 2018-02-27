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

	// open from byte test
	pf3 := pupfile.NewPupFile()
	data, _ := ioutil.ReadFile("test.zip")
	err = pf3.OpenFromStream(data)
	if err != nil {
		fmt.Println("open file fail! err is", err)
	}
	pf3.Close()

	imgData3 := pf3.GetPageImage(0)
	ioutil.WriteFile("out3.jpg", imgData3, os.ModePerm)
}

func readFileBytes(filename string) []byte {
	bytes, _ := ioutil.ReadFile(filename)
	return bytes
}
