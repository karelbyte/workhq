package routes

import (
	landlordcontrollers "elpuertodigital/workhq/controllers/landlord"
	_ "fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RoutesLandlord() *httprouter.Router {

	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/landlord/tenants", landlordcontrollers.TenantIndex)
	router.POST("/landlord/tenants", landlordcontrollers.TenantStore)
	router.POST("/landlord/tenants/:id", landlordcontrollers.TenantUpdate)
	router.GET("/landlord/tenants/:id", landlordcontrollers.TenantShow)

	return router
}
