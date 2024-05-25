package soal

import (
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
		// fmt.Println(maxLoop)
	}

	// kunci stop rekursive buat mengembalikan hasil
	if currentIndex == maxLoop {
		return strings.Join(strArr, "")
	}

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

	if strArr[currentIndex] == strArr[rightSideIndex] {
		if strArr[currentIndex] == "9" {
			currentIndex++
			return HighestPalindromeRecursive(s, k, strArr, currentIndex, maxLoop)
		}

		// jika hanya mengubah 1 sisi saja, maka palindrom akan gagal
		if k == 1 {
			currentIndex++
			return HighestPalindromeRecursive(s, k, strArr, currentIndex, maxLoop)
		}

		if k > 0 && k%2 == 0 {
			strArr[currentIndex] = "9"
			strArr[rightSideIndex] = "9"
			currentIndex++
			k = k - 2
			return HighestPalindromeRecursive(s, k, strArr, currentIndex, maxLoop)
		} else {
			currentIndex++
			return HighestPalindromeRecursive(s, k, strArr, currentIndex, maxLoop)
		}
	}

	return "-1"
}
