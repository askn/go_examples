package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output", os.O_CREATE|os.O_WRONLY, 0666)
	// bit düzeyinde or | birden fazla kullanmak için
	// os.O_RDONLY

	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	outputString := "Hello world!\n"

	for i := 0; i < 10; i++ {

		outputWriter.WriteString(outputString)
	}

	outputWriter.Flush()
	// yazmayı bitiriyor

	// stdout'a yaz
	os.Stdout.WriteString("hello, world\n")
}
