package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)



func init(){
	_  = godotenv.Load()
}

func Load() error {
	return godotenv.Load()
}


func GetString(key string) string{
	value , isPresent := os.LookupEnv(key)

	if(!isPresent){
		log.Fatalln("Error occurred while checked the key in the env ")
		return ""
	}

	return value
}

func GetInt(key string) int{
	value , isPresent := os.LookupEnv(key)

	if(!isPresent){
		log.Fatalln("Error occurred while checked the key in the env ")
		return 0
	}

	valueInt, err := strconv.Atoi(value)

	if err != nil{
		log.Fatal("Error occured while converting string to int",err)
		return 0
	}
	return valueInt
}