package handler

import (
	// "errors"
	"fmt"
	"net/http"
	"ToDoList/internal/models"
)

type Handler struct {
	message string
}

func (h Handler) Get() string {
	return h.message
}

func (h Handler) HandleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	u := new(model.User)
	u.Set(r.FormValue("username"), r.FormValue("password"))
	fmt.Println(u.Username, u.Pwd)

	// Теперь у вас есть значение "message" из формы, которое можно обработать или сохранить в базе данных
	// Вы можете использовать message в вашем коде для обработки данных
	return
}
