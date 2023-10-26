package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Загружаем HTML-файл в объект шаблона
	tmpl, err := template.ParseFiles("web/mainpage.html")
	if err != nil {
		fmt.Println("Ошибка загрузки шаблона:", err)
		log.Fatal(err)
	}

	// Обработчик для корневого URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Выводим HTML-код файла index.html
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Запуск сервера на порту 8080
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		log.Fatal(err)
	}
}