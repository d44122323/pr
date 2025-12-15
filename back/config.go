package main

type Config struct {
	TelegramBotToken string
	TelegramChatID   string
	EmailUser        string
	EmailPassword    string
}

func LoadConfig() Config {
	return Config{
		TelegramBotToken: "7938411427:AAEU3sCS-8ccGxnoVNABdhTPgflLBBieShk",
		TelegramChatID:   "7296280181",
		EmailUser:        "your@gmail.com",
		EmailPassword:    "your_app_password",
	}
}
