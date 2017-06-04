package godotenv

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var GlobalEnv = make(map[string]string, 5)

func ReadFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func EnvParser(lines []string) {
	for _, line := range lines {
		err, key, value := parseLine(line)
		if !err {
			GlobalEnv[key] = value
		}
	}
}

func parseLine(line string) (bool, string, string) {
	r, err := regexp.Compile(`^(\w+)=(.+)$`)
	if err != nil {
		fmt.Println(err)
	}
	result := r.FindStringSubmatch(line)
	fmt.Println(result)
	if len(result) > 0 {
		return false, result[1], result[2]
	}
	return true, "", ""
}
