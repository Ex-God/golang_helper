package utils

import (
	"io/ioutil"
  "log"
	"strings"
)

func SearchFromFile(query string, file string) ([]string) {
	data, err := ioutil.ReadFile(file)
	if err != nil { 		
		log.Fatal(err)
	}
	sdata := string(data)
	line := 0
	temp := strings.Split(sdata, "\n")
	answer := make([]string, 0)
	for _, s := range temp {
		if(strings.Contains(s, query)) {
			answer = append(answer, s)
		}
		line++
	}
	if (len(answer) > 10) {
		errText := make([]string, 0)
		errText = append(errText, "Too much")
		return errText
	}
	if (len(answer) < 1) {
		errText := make([]string, 0)
		errText = append(errText, "Not found")
		return errText
	}
	return answer
}