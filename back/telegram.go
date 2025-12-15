package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TgMessage struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendTelegram(cfg Config, data FormData) error {
	msg := TgMessage{
		ChatID: cfg.TelegramChatID,
		Text: fmt.Sprintf(
			"📨 *Новая заявка!*\n\n👤 Имя: %s\n📞 Телефон: %s\n💬 Комментарий: %s",
			data.Name, data.Phone, data.Comment,
		),
		ParseMode: "Markdown",
	}

	body, _ := json.Marshal(msg)
	resp, err := http.Post("https://api.telegram.org/bot"+cfg.TelegramBotToken+"/sendMessage",
		"application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("ошибка Telegram: %d", resp.StatusCode)
	}

	return nil
}
