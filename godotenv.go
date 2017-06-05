package godotenv

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func Godotenv(path string) map[string]string {
	lines := readFile(path)
	envs := dotenvParser(lines)
	return envs
}

func readFile(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func dotenvParser(lines []string) map[string]string {
	var dotenvmap = make(map[string]string, 5)
	for _, line := range lines {
		line = cleanLine(line)
		if line != "" {
			key, value, err := parseLine(line)
			if !err {
				dotenvmap[key] = value
			}
		}
	}
	return dotenvmap
}

func cleanLine(line string) string {
	line = strings.Split(line, "#")[0]
	line = strings.Trim(line, " ")
	return line
}

func parseLine(line string) (string, string, bool) {
	r, err := regexp.Compile(`^(\w+)=(.+)$`)
	if err != nil {
		fmt.Println(err)
	}
	result := r.FindStringSubmatch(line)
	if len(result) > 0 {
		return result[1], result[2], false
	}
	return "", "", true
}
