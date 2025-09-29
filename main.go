package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/letv1nnn/SysProg-TgBot/botui"
	"github.com/letv1nnn/SysProg-TgBot/sqlite"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Note: .env file not found, relying on environment variables")
	}

	token := os.Getenv("TOKEN")

	if token == "" {
		log.Panic("TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	store, err := sqlite.New("./sqlite/labs.db")
	if err != nil {
		log.Fatal(err)
	}

	labFlag := false

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if labFlag {
				taskName := update.Message.Text

				task, err := store.Get(taskName)
				if err != nil {
					log.Fatal(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Lab task by the given name does not exist")
					bot.Send(msg)
				}

				if task != nil {
					expl := tgbotapi.NewMessage(update.Message.Chat.ID, "EXPLANATION\n"+task.Expl)

					escapedCode := botui.EscapeMarkdownV2(taskName) + "\n```c\n" + botui.EscapeMarkdownV2(task.Code) + "\n```"
					code := tgbotapi.NewMessage(update.Message.Chat.ID, escapedCode)
					code.ParseMode = "MarkdownV2"

					if _, err := bot.Send(code); err != nil {
						log.Printf("failed to send code: %v", err)
					}
					if _, err := bot.Send(expl); err != nil {
						log.Printf("failed to send explanation: %v", err)
					}
				}

				labFlag = false
				continue
			}

			botui.SendReplyKeyboard(bot, update.Message.Chat.ID)

			switch update.Message.Text {
			case "CSC1038 topics covered":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "CSC1038 topics covered")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			case "Labs":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter lab task file name:")
				bot.Send(msg)
				labFlag = true
			case "Resources":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, botui.Resources)
				bot.Send(msg)
			case "About":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, botui.About)
				bot.Send(msg)
			}
		}
	}
}
