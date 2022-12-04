package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	largest := 0
	sum := 0
	for fileScanner.Scan() {
		calories := fileScanner.Text()
		if calories != "" {
			c, _ := strconv.Atoi(calories)
			if err != nil {
				fmt.Printf("error parsing calories: %v", err)
			}
			sum += c
		} else {
			if sum >= largest {
				largest = sum
			}
			sum = 0
		}
	}
	err = f.Close()
	if err != nil {
		return
	}
	fmt.Println(largest)
}
