package file

import "io/ioutil"

func Write(filePath string, content string) (err error) {
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	return err
}
