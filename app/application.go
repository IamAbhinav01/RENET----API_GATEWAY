package app

import (
	"log"
	"net/http"
	"renet/DB/repositories"
	"renet/config/db"
	"renet/config/env"
	"renet/controllers"
	"renet/router"
	"renet/services"
	"strings"
	"time"
)

type Application struct {
	PORT string
}

func NewApplication() *Application { // point towards same memory address
	return &Application{
		PORT: env.GetString("PORT"),
	}
}

func (app *Application) Run() error {
	addr := app.PORT

	//add : if not exsists

	if !strings.HasPrefix(addr, ":") {
		addr = ":" + addr
	}
	database, err := db.InitDB()
	if err != nil {
		return err
	}

	repository := repositories.NewUserRespository(database)
	service:=services.NewUserService(repository)
	controller:=controllers.NewAuthController(service)
	appRouter:=router.Router(controller)

	server := http.Server{
		Addr:         addr,
		Handler:      appRouter,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on PORT", addr)

	return server.ListenAndServe()
}
