package reflect

import "strings"

func ToLower(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += strings.ToLower(string(v))
		}
		if v >= 'a' && v <= 'z' {
				n += string(v)
		}
	}
	return n
}

