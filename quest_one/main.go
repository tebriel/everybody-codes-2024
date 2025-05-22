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

func main() {
	data, err := utils.ReadFile("./data/quest-1-part-1.txt")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Total potions needed:", PartOne(data))
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
