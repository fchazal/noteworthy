package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func ToPath(name string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	transformed, _, _ := transform.String(t, name)
	lowered := strings.ToLower(transformed)
	nospace := strings.ReplaceAll(lowered, " ", "_")
	noquote := strings.ReplaceAll(nospace, "'", "")
	return noquote
}

func InSlice(list []string, item string) bool {
	for _, b := range list {
		if b == item {
			return true
		}
	}
	return false
}

func RemoveFromSlice(list []string, item string) []string {
	index := 0
	for _, i := range list {
		if i != item {
			list[index] = i
			index++
		}
	}
	return list[:index]
}
