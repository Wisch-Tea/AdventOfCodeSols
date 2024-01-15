package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// AOC Day 1 input
var INPUT_FILE = "../day_1_input.txt"

var NUM_ARRAY = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

/**
 * This is an adventofcode.com solution for the 2023 advent calender day 1
 * The input is in a format of jchfsasix4lkhc where the first and last numbers (in this case 6 and 4)
 * are combined into a two-digit number (i.e. 64) based on whichever number appears first outboard to inboard
 * for both left and right.
 *
 * This solution could undoubtedly be improved as it has numerous manual checks to ensure that string slices
 * stay in bounds and don't run past 5 characters for the numerical string searches. Though it does have an
 * O(n)/linear time complexity per line.
 */
func main() {
	inputFile, err := os.Open(INPUT_FILE)
	if err != nil {
		log.Fatal("Failed to load the input file: "+INPUT_FILE, err)
		return
	}
	defer inputFile.Close()

	var firstNum string
	var secondNum string
	var sum int
	firstAssigned := false
	secondAssigned := false
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lineText := string(scanner.Text())
		lineArraySize := len(lineText) - 1
		for i := 0; i <= lineArraySize; i++ {
			// First check if the current index from the left is a number
			if !firstAssigned && isANumber(lineText[i]) {
				firstNum = string(lineText[i])
				firstAssigned = true
			} else if !firstAssigned && (i+2) <= lineArraySize {
				endSlice := (i + 5)
				beginSlice := i
				if endSlice-i > 5 {
					beginSlice = endSlice - 5
				}
				if endSlice > lineArraySize {
					endSlice = lineArraySize + 1
				}
				num, containsNum := containsNumString(lineText[beginSlice:endSlice], true)
				if containsNum {
					firstNum = strconv.Itoa(num)
					firstAssigned = true
				}
			}
			if !secondAssigned && isANumber(lineText[lineArraySize-i]) {
				secondNum = string(lineText[lineArraySize-i])
				secondAssigned = true
			} else if !secondAssigned && (lineArraySize-(i+2) >= 0) {
				beginSlice := lineArraySize - (i + 5)
				endSlice := (lineArraySize + 1) - i
				if endSlice > lineArraySize {
					endSlice = lineArraySize + 1
				}
				if beginSlice+endSlice > 5 {
					beginSlice = endSlice - 5
				}
				if beginSlice < 0 {
					beginSlice = 0
				}
				num, containsNum := containsNumString(lineText[beginSlice:endSlice], false)
				if containsNum {
					secondNum = strconv.Itoa(num)
					secondAssigned = true
				}
			}
			if firstAssigned && secondAssigned {
				lineValue, _ := strconv.Atoi(firstNum + secondNum)
				sum = lineValue + sum
				firstAssigned = false
				secondAssigned = false
				firstNum = ""
				secondNum = ""
				break
			}
		}
	}
	fmt.Printf("The total for "+INPUT_FILE+" is: %d", sum)
	// fmt.Println(string(content))
}

/**
 * A helper function that'll check if a passed in character is a number and return true if it is
 */
func isANumber(char byte) bool {
	if char < '0' || char > '9' {
		return false
	}
	return true
}

/**
 * A helper function that'll check if a number in text form exists within a passed in string slice
 * and return the number in integer form and a true value to indicate that a number was found
 */
func containsNumString(slice string, start bool) (int, bool) {
	if start && strings.ContainsAny(slice[:2], "123456789") {
		return 0, false
	} else if !start && strings.ContainsAny(slice[len(slice)-3:], "123456789") {
		return 0, false
	}
	for index, number := range NUM_ARRAY {
		if strings.Contains(slice, number) {
			// Increment index on return to match enum value
			return index + 1, true
		}
	}
	return 0, false
}
