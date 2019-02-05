package main

import (
	"strings"
)

func FindEnglishWords() (res []string) {
	var startIndex int
	var endIndex int
	for i, b := range data {
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') {
			if startIndex == 0 {
				startIndex = i
			} else {
				endIndex = i
			}
		} else {
			if startIndex > 0 {
				res = append(res, data[startIndex:endIndex+1])
				startIndex = 0
			}
		}
	}
	return res
}

func ReplaceGo() (res string) {
	bStr := []rune(data)
	var prevRune rune
	for i, r := range bStr {
		if prevRune == 0 {
			prevRune = r
			continue
		} else {
			if (prevRune == 'G' || prevRune == 'g') && (r == 'o' || r == 'O') {
				bStr[i-1] = 'O'
				bStr[i] = 'G'
				prevRune = 0
			}
		}
	}
	return string(bStr)
}

func StringReplaceGo() (res string) {
	res = strings.Replace(data, "GO", "OG", -1)
	res = strings.Replace(res, "gO", "OG", -1)
	res = strings.Replace(res, "Go", "OG", -1)
	res = strings.Replace(res, "go", "OG", -1)
	return res
}

func DumbReplaceGo() (res string) {
	var prevRune rune
	for _, r := range data {
		if prevRune == 0 {
			prevRune = r
			continue
		} else {
			if (prevRune == 'G' || prevRune == 'g') && (r == 'o' || r == 'O') {
				res += "OG"
			} else {
				res += string([]rune{prevRune, r})
			}
			prevRune = 0
		}
	}
	if prevRune != 0 {
		res += string(prevRune)
	}
	return res
}
