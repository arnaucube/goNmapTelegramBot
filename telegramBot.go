package main

import (
	"fmt"
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func telegramBot() {
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		fmt.Println("new ip to scan: " + update.Message.Text)
		var result string
		ok := checkIp(update.Message.Text)
		if update.Message.Text == "/start" {
			result = "welcome to nmap telegram bot, written in Go lang"
		} else {

			if ok {
				result = nmap(update.Message.Text)
			} else {
				result = "ip not valid"
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, result)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
