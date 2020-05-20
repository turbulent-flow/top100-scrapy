package conversion

import (
	"strconv"
	"strings"
)

func ToSingleString(params []int) string {
	paramStrings := make([]string, 0)
	for _, param := range params {
		paramString := strconv.Itoa(param)
		paramStrings = append(paramStrings, paramString)
	}
	return strings.Join(paramStrings, ",")
}
