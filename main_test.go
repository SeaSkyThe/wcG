package main

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "test.txt"
	_, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("The function read_file is not working: %q", err)
	}
}

func TestProcessing(t *testing.T) {
	file, _ := ReadFile("test.txt")
	flags := []string{"-c", "-l", "-w", "-m"}

	expected_count := 0
	for _, flag := range flags {
        count := ProcessFlag(flag, file)
		if flag == "-c" {
			expected_count = 342190
			if count != expected_count {
				t.Fatalf("The processing returned a wrong number of bytes. expected %d. current %d.", expected_count, count)
			}
            fmt.Println("Flag -c worked correctly")
		} else if flag == "-l" {
			expected_count = 7145
			if count != expected_count {
				t.Fatalf("The processing returned a wrong number of lines. expected %d. current %d.", expected_count, count)
			}
            fmt.Println("Flag -l worked correctly")
		} else if flag == "-w" {
			expected_count = 58164
			if count != expected_count {
				t.Fatalf("The processing returned a wrong number of words. expected %d. current %d.", expected_count, count)
			}
            fmt.Println("Flag -w worked correctly")
		} else if flag == "-m" {
			expected_count = 339292
			if count != expected_count {
				t.Fatalf("The processing returned a wrong number of characters. expected %d. current %d.", expected_count, count)
			}
            fmt.Println("Flag -m worked correctly")
        }
	}
}
