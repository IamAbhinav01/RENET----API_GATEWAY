package app

import (
	"log"
	"net/http"
	"renet/config/db"
	"renet/config/env"
	"renet/router"
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

func (app *Application)Run() error{
	addr := app.PORT

	//add : if not exsists

	if(!strings.HasPrefix(addr,":")){
		addr = ":"+addr
	}
	db,_ := db.InitDB()
	log.Println(db)
	server:=http.Server{
		Addr: addr,
		Handler: router.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on PORT",addr)

	return server.ListenAndServe()
}