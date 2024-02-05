package main

import (
	"context"
	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type gis struct {
	caption string
	url     string
	texturl string
}
type gisqq struct {
	caption string
	url     string
}
type kbutt struct {
	caption        string
	KeyboardButton tgbotapi.ReplyKeyboardMarkup
}
type kbuttirl struct {
	caption        string
	KeyboardButton tgbotapi.InlineKeyboardMarkup
}
type helloy struct {
	caption string
}

func (bot *Bot) sendScreenshot(chatId int64, MetioAndAnime gis) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := MetioAndAnime.url
	var imageBuf []byte
	if err := chromedp.Run(
		ctx,
		ScreenshotTasks(url, &imageBuf),
	); err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FileBytes{
		Name:  "pict",
		Bytes: imageBuf,
	})
	msg.Caption = MetioAndAnime.caption
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(MetioAndAnime.texturl, MetioAndAnime.url),
		),
	)

	bot.Send(msg)
}
func (bot *Bot) sendMetiooqq(From int64, Gismetqq gisqq) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	url := Gismetqq.url
	var imageBuf []byte
	if err := chromedp.Run(
		ctx,
		ScreenshotTasks(url, &imageBuf),
	); err != nil {
		log.Fatal(err)
	}
	msg := tgbotapi.NewPhoto(From, tgbotapi.FileBytes{
		Name:  "pict",
		Bytes: imageBuf,
	})
	msg.Caption = Gismetqq.caption
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Посмотреть вы можете тут 👉", Gismetqq.url),
		),
	)

	bot.Send(msg)
}
func (bot *Bot) sendKeyboardButtn(From int64, KeyBut kbutt) {
	msg := tgbotapi.NewMessage(From, KeyBut.caption)
	msg.ReplyMarkup = KeyBut.KeyboardButton
	bot.Send(msg)
}
func (bot *Bot) sendKeyboardButtnIrl(From int64, KeyBytIrl kbuttirl) {
	msg := tgbotapi.NewMessage(From, KeyBytIrl.caption)
	msg.ReplyMarkup = KeyBytIrl.KeyboardButton
	bot.Send(msg)
}
func (bot *Bot) sendHelloy(chatId int64, helo helloy) {
	msg := tgbotapi.NewMessage(chatId, helo.caption)
	bot.Send(msg)
}

func ScreenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ScreenshotScale(url, 100, imageBuf),
	}
}
