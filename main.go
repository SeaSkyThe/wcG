package main

import (
	"bufio"
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
	for scanner.Scan() {
		count = count + 1
	}
	return count
}

func ProcessFlag(flag string, file []byte) int {
	reader := bytes.NewReader(file)
	count := 0

	switch flag {
	case "-c":
		count = count_bytes(reader)
	case "-l":
		count = count_lines(reader)
	case "-w":
		count = count_words(reader)
	case "-m":
		count = count_characters(reader)
	default:
		println("Flag not recognized: ", flag)
		println("Usage: wcg -flag(optional) <filename>")
		println("Available flags: \n -c: count the number of bytes \n -l: count the number of lines \n -w: count the number of words \n -m: count the number of characters \n no flags: will output the equivalent of using -c, -l and -w flags. ")
	}

	return count
}

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		println("Error opening file\n")
		println("Usage: wcg -flag(optional) <filename>")
		println("Available flags: \n -c: count the number of bytes \n -l: count the number of lines \n -w: count the number of words \n -m: count the number of characters \n no flags: will output the equivalent of using -c, -l and -w flags. ")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ReadStdin() []byte {
	var content []byte
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		content, _ = io.ReadAll(os.Stdin)
	} else {
		println("Usage: wcg -flag(optional) <filename>")
		println("Available flags: \n -c: count the number of bytes \n -l: count the number of lines \n -w: count the number of words \n -m: count the number of characters \n no flags: will output the equivalent of using -c, -l and -w flags. ")
		return nil
	}
	return content
}

func main() {
	args := os.Args
	var filename string = ""
	var content []byte
	flags := []string{"-l", "-w", "-c"}

	// Getting Flags (generally all the arguments except the last)
	if len(args) >= 2 {
		flags = args[1 : len(args)-1]
		// Check if the last is a flag
		last_argument := args[len(args)-1]
		// If it is, goes to the flag list
		if strings.Contains(last_argument, "-") && len(last_argument) == 2 {
			flags = append(flags, last_argument)
            content = ReadStdin()
		} else {
			filename = last_argument
			content, _ = ReadFile(filename)

            if len(flags) == 0{
	            flags = []string{"-l", "-w", "-c"}
            }
		}
	} else {
        content = ReadStdin()
	}

	output := " "
	for _, flag := range flags {
		output = output + fmt.Sprintf("%8d", ProcessFlag(flag, content))
	}
	output += fmt.Sprintf(" %s", filename)
	println(output)
}
