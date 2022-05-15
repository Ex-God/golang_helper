package utils 

import (
	"os"
	"io/ioutil"
  "log"
	"strings"
)

func SetEnv() {
	data, err := ioutil.ReadFile(".env")
	if err != nil { 		
		log.Fatal(err)
	}
	sdata := string(data)
	line := 0
	temp := strings.Split(sdata, "\n")
	for _, s := range temp {
		sline := strings.Split(s, "=")
		os.Setenv(sline[0], sline[1])
		line++
	}
}