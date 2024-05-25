package main

import (
	"fmt"

	"github.com/fredy-bambang/bank-bri/soal"
)

func main() {
	fmt.Println("----------------SOAL 1 Weighted Strings------------------------")
	// SOAL 1
	hasilSoal11, _ := soal.WeightedStringMultipliedCompareNumberQuery("abbcccd", []int{1, 3, 9, 8})
	hasilSoal12, _ := soal.WeightedStringMultipliedCompareNumberQuery("abcd", []int{1, 3, 9, 8})
	hasilSoal13, _ := soal.WeightedStringMultipliedCompareNumberQuery("aaabbbcccddd", []int{1, 3, 9, 12})
	fmt.Println(hasilSoal11)
	fmt.Println(hasilSoal12)
	fmt.Println(hasilSoal13)

	fmt.Println("")
	fmt.Println("----------------SOAL 2------------------------")

	// SOAL 2
	testCases := []string{
		"{ [ ( ) ] }",
		"{ [ ( ] ) }",
		"{ ( ( [ ] ) [ ] ) [ ] }",
		"{ [ }",
		"[ ( { ) } ]",
		"([(({()}))])",
		"[()]",
	}

	for _, test := range testCases {
		fmt.Printf("Input: %s\n", test)
		if soal.IsValidBrackets(test) {
			fmt.Println("Output: Valid")
		} else {
			fmt.Println("Output: Invalid")
		}
	}

	fmt.Println("")
	fmt.Println("----------------SOAL 3------------------------")

	// SOAL 3
	fmt.Println(soal.HighestPalindromeRecursive("3943", 1, []string{}, 0, 0))
	fmt.Println(soal.HighestPalindromeRecursive("932239", 2, []string{}, 0, 0))
	fmt.Println(soal.HighestPalindromeRecursive("12345678", 10, []string{}, 0, 0))
}
