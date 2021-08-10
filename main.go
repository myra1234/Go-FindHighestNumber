package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Enter some numbers with space between each number, hit ENTER when done")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Some error occured while reading your input, please try again", err)
		return
	}

	input = strings.ReplaceAll(input, "\n", "")
	strSlice := strings.Split(input, " ")

	intSlice := make([]int, len(strSlice))
	for k, j := range strSlice {
		intSlice[k], _ = strconv.Atoi(j)
	}

	fmt.Println("Highest number is:", findHighestNumber(intSlice))
}

func findHighestNumber(input []int) int {
	highest := input[0]

	for _, i := range input {
		if i > highest {
			highest = i
		}
	}
	return highest
}
