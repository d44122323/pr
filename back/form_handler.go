package main

import (
	"fmt"
	"net/http"
)

func FormHandler(cfg Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data := FormData{
			Name:    r.FormValue("name"),
			Phone:   r.FormValue("phone"),
			Comment: r.FormValue("comment"),
		}

		if err := ValidateForm(data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := SendTelegram(cfg, data); err != nil {
			fmt.Println("Ошибка отправки в Telegram:", err)
		}

		SendEmail(cfg, data)

		http.Redirect(w, r, "/success.html", http.StatusSeeOther)
	}
}
