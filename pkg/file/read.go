package file

import (
	"io/ioutil"
	"strings"
)

func Read(filePath string) (content string, err error) {
	raw, err := ioutil.ReadFile(filePath)
	content = strings.TrimSpace(string(raw))
	return content, err
}
