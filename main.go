package main

import (
	"log"
	"renet/app"
)

func main(){
	application := app.NewApplication()
	err := application.Run()
	if err != nil{
		log.Println("Error occured while starting the application")
	}
}