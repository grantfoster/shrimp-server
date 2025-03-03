package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"reflect"
)

func main() {
	name := flag.String("n", "empty", "snake case name for your new migrations")
	flag.Parse()
	if *name == "empty" {
		slog.Error("flag n (name) must not be empty")
		os.Exit(0)
	}
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Println(err)
	}

	var versionNumbers []int
	for _, file := range files {
		versionNumber, err := firstCharToInt(file.Name())
		if err != nil {
			slog.Error("there was a problem converting the first character of the filename to an integer",
				"err", err.Error(),
			)
			os.Exit(0)
		}
		versionNumbers = append(versionNumbers, versionNumber)
	}

	largestNumber, err := findLargestNumber(versionNumbers)
	if err != nil {
		slog.Error("there was a problem finding the largest number in our list of files",
			"err", err.Error(),
		)
	}
	_, err = os.Create(fmt.Sprintf("migrations/%v_%s.up.sql", largestNumber+1, *name))
	if err != nil {
		slog.Error("there was an error creating your migration file",
			"err", err.Error(),
		)
	}
	_, err = os.Create(fmt.Sprintf("migrations/%v_%s.down.sql", largestNumber+1, *name))
	if err != nil {
		slog.Error("there was an error creating your migration file",
			"err", err.Error(),
		)
	}
}

func firstCharToInt(s string) (int, error) {
	// Check if string is empty
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}

	// Get the first character
	firstChar := s[0]

	// Check if it's a digit
	if firstChar < '0' || firstChar > '9' {
		return 0, fmt.Errorf("first character of string is not a digit")
	}

	// Convert to numeric value
	// In ASCII/UTF-8, digits are sequential, so we can subtract '0'
	return int(firstChar - '0'), nil
}

func findLargestNumber(numbers []int) (int, error) {
	highestNumber := numbers[0]
	for index, number := range numbers {
		if len(numbers) == 0 {
			return 0, errors.New("no file version numbers found")
		}
		if reflect.TypeOf(number).Kind() != reflect.Int {
			return 0, fmt.Errorf("encountered a chracter that is not a number at index: %v", index)
		}
		if number > highestNumber {
			highestNumber = number
		}
	}
	return highestNumber, nil
}
