package handler

import (
	// "errors"
	// "ToDoList/internal/models"
	// "fmt"
	// "net/http"

	"ToDoList/internal/service"
	// "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// // tut old
// type Handler struct {
// 	message string
// }

// func (h Handler) Get() string {
// 	return h.message
// }

// func (h Handler) HandleSubmit(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	u := new(model.User)
// 	u.Set(r.FormValue("username"), r.FormValue("password"))
// 	fmt.Println(u.Username, u.Pwd)

// 	return
// }

// // tut new
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		// fmt.Println("AaA")
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
