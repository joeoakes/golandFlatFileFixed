package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define a struct to hold the student information
type Student struct {
	Name    string
	Age     string
	Subject string
}

// Function to parse a line based on fixed widths
func parseLine(line string) Student {
	return Student{
		Name:    strings.TrimSpace(line[0:12]),
		Age:     strings.TrimSpace(line[12:15]),
		Subject: strings.TrimSpace(line[15:]),
	}
}

func main() {
	// Open the flat file
	file, err := os.Open("students.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		student := parseLine(line)
		fmt.Printf("Name: %s, Age: %s, Subject: %s\n", student.Name, student.Age, student.Subject)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
