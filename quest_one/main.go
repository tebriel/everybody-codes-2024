// Package questone encapsulates the first quest tasks
package questone

import (
	"fmt"
	"os"

	"github.com/tebriel/everybody-codes-2024/utils"
)

// PotionData contains information about the potions we need
type PotionData struct {
	Needed int
}

// RunAll runs all parts of the quest
func RunAll() {
	data, err := utils.ReadFile("./data/quest-1-part-1.txt")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Total potions needed:", PartOne(data))

	data, err = utils.ReadFile("./data/quest-1-part-2.txt")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Total potions needed:", PartTwo(data))
}

// PartOne calculates potions based on our first input
func PartOne(data string) int {
	creatures := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
	}

	results := PotionData{}
	runes := []rune(data)
	for _, r := range runes {
		results.Needed += creatures[r]
	}
	return results.Needed
}

// PartTwo calculates potions based on our second input
func PartTwo(data string) int {
	creatures := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
		'x': 0,
	}
	results := PotionData{}
	runes := []rune(data)
	for i := 0; i < len(data); i += 2 {
		first := runes[i]
		second := runes[i+1]
		results.Needed += creatures[first] + creatures[second]
		if first != 'x' && second != 'x' {
			results.Needed += 2
		}
	}
	return results.Needed
}
