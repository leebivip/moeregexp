// moeregexp.go
package moeregexp

import (
	"regexp"
)

func IsMatch(regexstr string, input string) bool {
	reg := regexp.MustCompile(regexstr)
	return reg.MatchString(input)
}

func GetMatchCollection(regexstr string, input string) []string {
	reg := regexp.MustCompile(regexstr)
	var mc []string
	mc = reg.FindAllString(input, -1)
	return mc
}
