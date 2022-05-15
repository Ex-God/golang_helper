package telegram

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
	"bytes"
	"fmt"
)

func GetUpdates(botUrl string, offset int) ([]Update, error){
	response, err := http.Get(botUrl + "/getUpdates" + "?offset=" + fmt.Sprint(offset))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Response
	err = json.Unmarshal(blob, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.Result, nil
}

func SendMessage(botUrl string, message map[string]interface{}) {
	blob, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(blob))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}