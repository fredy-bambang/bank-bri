package main

import (
	"fmt"

	"github.com/fredy-bambang/bank-bri/soal"
)

func main() {
	hasilSoal11, _ := soal.WeightedStringMultipliedCompareNumberQuery("abbcccd", []int{1, 3, 9, 8})
	hasilSoal12, _ := soal.WeightedStringMultipliedCompareNumberQuery("abcd", []int{1, 3, 9, 8})
	hasilSoal13, _ := soal.WeightedStringMultipliedCompareNumberQuery("aaabbbcccddd", []int{1, 3, 9, 12})
	fmt.Println(hasilSoal11)
	fmt.Println(hasilSoal12)
	fmt.Println(hasilSoal13)
}
