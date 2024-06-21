package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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
			_ = fmt.Errorf("Error when counting lines: %q", err)
			return 0
		}
	}
}

func count_words(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    count := 0
    for scanner.Scan() {
        count = count + 1
    }
	return count
}

func count_characters(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanRunes)
    count := 0
    for scanner.Scan(){
        count = count + 1
    }
    return count
}

func ProcessFlag(flag string, file *os.File) int {
	file.Seek(0, io.SeekStart)
	count := 0

	switch flag {
	case "-c":
		count = count_bytes(file)
	case "-l":
		count = count_lines(file)
    case "-w":
        count = count_words(file)
    case "-m":
        count = count_characters(file)
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
    defer file.Close()

    flags := args[1:len(args)-1]

    if len(flags) == 0 {
        flags = []string{"-l", "-w", "-c"}
    }

    output := " "
    for _, flag := range flags {
        output = output + fmt.Sprintf(" %d", ProcessFlag(flag, file))
    }
    output += fmt.Sprintf(" %s", filename)

	println(output)
}
