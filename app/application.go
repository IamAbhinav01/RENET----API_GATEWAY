package app

import (
	"log"
	"net/http"
	"renet/config/env"
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

	server:=http.Server{
		Addr: addr,
		Handler: http.DefaultServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on PORT",addr)

	return server.ListenAndServe()
}