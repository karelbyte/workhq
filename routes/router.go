package routes

import (
	_ "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	// router.GET("/api/users", controllers.UserIndex)
	// router.POST("/api/users", controllers.UserStore)
	// router.POST("/api/users/:id", controllers.UserUpdate)
	// router.GET("/api/users/:id", controllers.UserShow)

	return router
}
