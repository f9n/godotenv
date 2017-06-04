package godotenv

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
