package telegram

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"golang_helper/utils"
)

func StartBot(botToken string) {
	botUrl := "https://api.telegram.org/bot" + botToken
	offset := 0
	fmt.Println("Bot is started...")
	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			AnswerBot(botUrl, update)
			offset = update.Update_Id + 1
		}
	}
}

func AnswerBot(botUrl string, update Update) {
	// fmt.Printf("%+v", update)
	if checkAdmin(update, os.Getenv("GH_ADMIN_ID")) {
		if(update.Message.Text == "/start") {
			message := NewMessage(update.Message.From.Id, "Ready to search")
			SendMessage(botUrl, message)
			return
		}
		file := os.Getenv("GH_SEARCH_FILE")
		query := update.Message.Text
		text := utils.SearchFromFile(query, file)
		for _, r := range text {
			message := NewMessage(update.Message.From.Id, r)
			SendMessage(botUrl, message)
		}
	}
}

func NewMessage(chatId int, text string) (map[string]interface{}) {
	message := map[string]interface{}{
		"chat_id": chatId,
		"text": text,
	}
	return message
}

func NewKeyboard(buttons [][]map[string]interface{}) (map[string]interface{}) {
	keyboard := map[string]interface{}{
		"inline_keyboard": buttons,
	}
	return keyboard
}

func checkAdmin(update Update, adminId string) (bool) {
	if strconv.Itoa(update.Message.From.Id) == adminId {
		return true
	}
	return false
}