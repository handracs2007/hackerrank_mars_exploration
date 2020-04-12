// https://www.hackerrank.com/challenges/mars-exploration/problem

package main

import (
	"fmt"
)

func marsExploration(s string) int {
	var alteredCount = 0
	var message = "SOS"

	for i := 0; i < len(s); i += 3 {
		for j := 0; j < 3; j++ {
			if s[i+j] != message[j] {
				alteredCount++
			}
		}
	}

	return alteredCount
}

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
}
