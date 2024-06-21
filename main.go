package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func count_bytes(file io.Reader) int {
	byts, _ := io.ReadAll(file)
	return len(byts)
}

func count_words(line string) int {
	words_slice := strings.Split(line, " ")
	return len(words_slice)
}

func ProcessFlags(flag string, file io.Reader) int {
	count := 0
	if flag == "-c" {
		count = count_bytes(file)
	}
	return count
}

func ReadFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		println("Error opening file\n")

		println("Usage: wcg -flag(optional) <filename>")
		println("Available flags: \n -c: count the number of bytes \n -l: count the number of lines \n -w: count the number of words \n -m: count the number of characters \n no flags: will output the equivalent of using -c, -l and -w flags. ")
	}
    return file, err
}

func main() {
	args := os.Args
	if len(args) < 2 {
		println("Usage: wcg -flag(optional) <filename>")
		println("Available flags: \n -c: count the number of bytes \n -l: count the number of lines \n -w: count the number of words \n -m: count the number of characters \n no flags: will output the equivalent of using -c, -l and -w flags. ")
		return
	}
	filename := args[len(args)-1]
    file, _ := ReadFile(filename)

	count := ProcessFlags(args[1], file)
	output := fmt.Sprintf("%d %s", count, filename)

	println(output)
}
