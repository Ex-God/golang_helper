package main

import (
	"os"
	"golang_helper/utils"
	"golang_helper/telegram"
)

func main() {
	utils.SetEnv()
	botToken := os.Getenv("GH_BOT_TOKEN")
	telegram.StartBot(botToken)
}