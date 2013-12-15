package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "main.go"
	outputFile := "main_copy"

	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))

	err = ioutil.WriteFile(outputFile, buf, 0x644)

	if err != nil {
		panic(err.Error())
	}
}
