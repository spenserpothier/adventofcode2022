package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day03/input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	items := []string{}

	for fileScanner.Scan() {
		ruckSack := fileScanner.Text()
		compartmentA := ruckSack[:len(ruckSack)/2]
		compartmentB := ruckSack[len(ruckSack)/2:]

		for o := 0; o < len(compartmentB); o++ {
			found := false
			for i := 0; i < len(compartmentA); i++ {
				if string(compartmentA[i]) == string(compartmentB[o]) {
					items = append(items, string(compartmentA[i]))
					found = true
					break
				}

			}
			if found {
				break
			}
		}
	}

	var sum uint32
	for _, item := range items {
		itemRune := rune(item[0])
		if itemRune >= 'a' && itemRune <= 'z' {
			sum += uint32(itemRune) - 96
		} else if itemRune >= 'A' && itemRune <= 'Z' {
			sum += uint32(itemRune) - 65 + 27
		}
	}
	fmt.Printf("%d\n", sum)
	err = f.Close()
	if err != nil {
		return
	}
}
: