package soal

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

var charWeights = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5,
	'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
	'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15,
	'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20,
	'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25,
	'z': 26,
}

// WeightedString calculates the total weight of the given string.
func WeightedString(s string) int {
	totalWeight := 0
	for _, char := range s {
		if weight, exists := charWeights[char]; exists {
			// fmt.Println("char:", char, "weight:", weight)
			totalWeight += weight
		}
	}
	return totalWeight
}

func WeightedStringMultipliedCompareNumberQuery(s string, nums []int) (result []bool, err error) {
	lowercaseStrings := strings.ToLower(s)

	// hanya boleh alfabet
	for _, r := range lowercaseStrings {
		if !unicode.IsLetter(r) {
			// return error jika bukan huruf ditemukan
			fmt.Println("karakter ", r, " itu bukan huruf")
			return result, errors.New("parameter kata hanya boleh mengandung huruf")
		}
	}

	arrOfStrings := strings.Split(lowercaseStrings, "")
	sort.Strings(arrOfStrings)

	tempChar := ""
	sumChar := 0
	indexNum := 0
	// direct compare queries
	for i := 0; i < len(arrOfStrings); i++ {
		// init char here
		if tempChar == "" {
			tempChar = arrOfStrings[i]
		}

		// if next char is different with previous char, compare with query and set the boolean result
		if tempChar != arrOfStrings[i] {
			// fmt.Println(tempChar, " vs ", arrOfStrings[i], " - ", "sumChar: ", sumChar, " vs ", nums[indexNum])
			if sumChar == nums[indexNum] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
			tempChar = arrOfStrings[i]
			sumChar = 0
			indexNum++
		}

		// start sum the weight for same char
		sumChar += WeightedString(tempChar)
		if i == len(arrOfStrings)-1 {
			if sumChar == nums[indexNum] {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
		}
	}

	return result, nil
}
