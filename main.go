package main

import (
	"elpuertodigital/workhq/routes"
	"elpuertodigital/workhq/setup"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main ()  {
	// load enviroment
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error: file .env not found.")
	}
	
	// db logic init, check migrate and seed
	setup.Migrate()

		// api routes
	routes := routes.Routes()
	http.ListenAndServe(":8080", routes)

	fmt.Println("---------------------")
	fmt.Println("System online!")
}