package godotenv

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

func EnvParser(line string) {
	r, err := regexp.Compile(`^(.+)=(.+)$`)
	if err != nil {
		fmt.Println(err)
	}
	result := r.FindStringSubmatch(line)
	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
	fmt.Println(result)
}
