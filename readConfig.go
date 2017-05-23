package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Config stores the data from json twitterConfig.json file
type Config struct {
	TelegramBotToken string `json:"telegramBotToken"`
}

var config Config

func readConfig() {
	file, e := ioutil.ReadFile("config.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	fmt.Println("config.json read comlete, Telegram bot ready")

}
