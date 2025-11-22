package scale

import (
	"strings"
	"unicode"
)

func Scale(tonic string, interval string) []string {
	var scale []string
	var keys []string
	var i int = 0

	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	sharp_scales := []string{"C", "G", "f#", "a", "e", "A"}
	flat_scales := []string{"F", "bb", "Eb", "g", "d", "Db"}

	// 2 octave maps for sharp and flat scales
	sharp_scale_keys := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	flat_scale_keys := []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}

	if Find(sharp_scales, tonic) < len(sharp_scales) {
		keys = sharp_scale_keys
	} else if Find(flat_scales, tonic) < len(sharp_scales) {
		keys = flat_scale_keys
	}

	i = Find(keys, UcFirst(tonic))

	for _, v := range strings.Split(interval, "") {

		scale = append(scale, keys[i])

		if v == "M" {
			i = i + 2
		} else if v == "A" {
			i = i + 3
		} else {
			i = i + 1
		}
	}

	return scale
}

// REF: https://yourbasic.org/golang/find-search-contains-slice/
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// REF: https://blog.csdn.net/weixin_30258901/article/details/95072251
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
