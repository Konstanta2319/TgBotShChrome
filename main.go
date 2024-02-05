package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const Token = ""

type Bot struct {
	*tgbotapi.BotAPI
}

func main() {
	tgbot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	tgbot.Debug = true
	bot := &Bot{
		BotAPI: tgbot,
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatId := update.Message.Chat.ID

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			text := strings.ToLower(update.Message.Text)

			gis, ok := MetioAndAnime[text]
			if ok {
				bot.sendScreenshot(chatId, gis)
				continue
			}
			helloy, ok := helo[text]
			if ok {
				bot.sendHelloy(chatId, helloy)
				continue
			}
			kbuttirl, ok := KeyBytIrl[text]
			if ok {
				bot.sendKeyboardButtnIrl(chatId, kbuttirl)
				continue
			}
			kbutt, ok := KeyBut[text]
			if ok {
				bot.sendKeyboardButtn(chatId, kbutt)
				continue
			}

			//if text == "как дела?" {
			//	msg := tgbotapi.NewMessage(chatId, answers2[rand.Intn(len(answers2))])
			//	bot.Send(msg)
			//} else {
			//	msg := tgbotapi.NewMessage(chatId, "я не понимю, простите😭")
			//	bot.Send(msg)
			//}

		}
		if update.CallbackQuery != nil {
			From := update.CallbackQuery.From.ID
			text := strings.ToLower(update.CallbackQuery.Data)
			gisqq, ok := MetioData[text]
			if ok {
				bot.sendMetiooqq(From, gisqq)
				continue
			}
			kbutt, ok := KeyBut[text]
			if ok {
				bot.sendKeyboardButtn(From, kbutt)
				continue
			}
		}
	}
}
