package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func count_bytes(file io.Reader) int {
	byts, _ := io.ReadAll(file)
	return len(byts)
}

func count_lines(file io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			fmt.Errorf("Error when counting lines: %q", err)
			return 0
		}
	}
}

func count_words(line string) int {
	words_slice := strings.Split(line, " ")
	return len(words_slice)
}

func ProcessFlags(flag string, file *os.File) int {
    file.Seek(0, io.SeekStart)
	count := 0
	if flag == "-c" {
		count = count_bytes(file)
	} else if flag == "-l" {
		count = count_lines(file)
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
