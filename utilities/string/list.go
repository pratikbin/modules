package string

import (
	"strings"
)

// TODO write testcase for empty and singular input
func JoinListStrings(listStrings ...string) string {
	return strings.Join(listStrings, idSeparator)
}

func SplitListString(listString string) []string {
	return strings.Split(listString, idSeparator)
}
