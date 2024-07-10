package internal

import "strings"

func EncodeURIComponent(str string) string {
	result := strings.Replace(str, "+", "%20", -1)
	result = strings.Replace(result, "%21", "!", -1)
	result = strings.Replace(result, "%27", "'", -1)
	result = strings.Replace(result, "%28", "(", -1)
	result = strings.Replace(result, "%29", ")", -1)
	result = strings.Replace(result, "%2A", "*", -1)
	return result
}
