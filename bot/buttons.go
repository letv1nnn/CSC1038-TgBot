package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func sendReplyKeyboard(bot *tgbotapi.BotAPI, chatID int64) {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("CSC1038 topics covered"),
			tgbotapi.NewKeyboardButton("Labs"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Resources"),
			tgbotapi.NewKeyboardButton("About"),
		),
	)
	keyboard.ResizeKeyboard = true
	msg := tgbotapi.NewMessage(chatID, "Choose an option:")
	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}
