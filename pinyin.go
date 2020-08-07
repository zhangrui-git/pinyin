package pinyin

import (
	"strings"
	"unicode/utf8"
)

var dict = make(map[string][]string)

type Pinyin struct {
	string
	FullPinyin []string
	Initials []string
}

func init()  {
	parts := strings.Split(dictData, "|")
	for _, i := range parts {
		idx := strings.Split(i, ":")
		dict[idx[0]] = strings.Split(idx[1], ",")
	}
}

func New(s string) *Pinyin {
	full, initials := make([]string, utf8.RuneCountInString(s)), make([]string, utf8.RuneCountInString(s))
	for _, word := range s {
		ph := " "
		rn := utf8.RuneLen(rune(word))
		if rn > 1 {
			for idx, chars := range dict {
				for _, char := range chars {
					if string(word) == char {
						ph = idx
					}
				}
			}
		} else if rn == 1 {
			ph = string(word)
		}
		full = append(full, ph)
		initials = append(initials, string([]rune(ph)[0]))
	}
	return &Pinyin{string: s, FullPinyin: full, Initials: initials}
}