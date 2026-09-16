package app

type Application struct {
	PORT string
}


func NewApplication() *Application{ // point towards same memory address
	return &Application{
		PORT: ,
	}
}