package soal

import (
	"fmt"
	"strings"
)

func HighestPalindromeRecursive(
	s string, k int, strArr []string, currentIndex int, maxLoop int,
) string {
	// k harus minimal 0 dan positif
	if k < 0 {
		return "-1"
	}

	// string tidak boleh kosong
	if len(s) <= 0 {
		return "-1"
	}

	// panjang string harus minimal 1
	if len(s) == 1 {
		if k > 0 {
			return "9"
		}
		return s
	}

	// split to array when string length > 1
	if len(strArr) == 0 {
		strArr = strings.Split(s, "")

		// no need loop all index because of palindrome only need half of the string
		maxLoop = len(strArr) / 2
		fmt.Println(maxLoop)
	}

	// kunci stop rekursive buat mengembalikan hasil
	if currentIndex == maxLoop {
		return strings.Join(strArr, "")
	}

	// fmt.Println("currentIndex: ", currentIndex, "-> ", strArr[currentIndex], ",rightSideIndex: ", len(strArr)-1-currentIndex, "-> ", strArr[len(strArr)-1-currentIndex])

	// menggunakan angka terbesar dari kiri atau kanan
	rightSideIndex := len(strArr) - 1 - currentIndex
	if strArr[currentIndex] != strArr[rightSideIndex] {
		if strArr[currentIndex] > strArr[rightSideIndex] {
			strArr[rightSideIndex] = strArr[currentIndex]
		} else {
			strArr[currentIndex] = strArr[rightSideIndex]
		}
		currentIndex++
		k--
		return HighestPalindromeRecursive(s, k, strArr, currentIndex, maxLoop)
	}

	return "-1"
}
