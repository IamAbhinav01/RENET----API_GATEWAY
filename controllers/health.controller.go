package controllers

import (
	"net/http"
	"renet/utils/formatters"
)


func HealthController(w http.ResponseWriter,r *http.Request){
	
	formatters.SuccessResponse(w,http.StatusAccepted,"Health route is working")
}