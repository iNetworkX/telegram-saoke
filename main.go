package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7787538544:AAHu2SUr_v-NxZyUvuF2pRlNTZPDyWCB48s")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Xin nhập Mã số giao dịch")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				maSoGD := strings.ReplaceAll(update.Message.Text, " ", "")
				ketQua := Kiem_tra_ket_qua(maSoGD)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, ketQua)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}

func Kiem_tra_ket_qua(maSo string) string {
	var ketQua string
	content, err := os.ReadFile("saoke_modified.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	saoKe := strings.Split(string(content), "|")
	for i := 0; i < len(saoKe); i++ {
		if strings.Contains(saoKe[i], maSo) {
			saoKe[i] = strings.Replace(saoKe[i], "@", "\n", 3)
			ketQua = saoKe[i]
		}
	}
	return ketQua
}
