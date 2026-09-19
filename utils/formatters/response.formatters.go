package formatters

import (
	"net/http"
	"renet/utils/converter"
)

func SuccessResponse(w http.ResponseWriter,status int,data any) error{
	response := map[string]any{}
	response["status"] = status
	response["data"] = data

	return converter.ConvertToJSON(w,response)
}
func ErrorResponse(w http.ResponseWriter,status int,err error,data any) error{
	response := map[string]any{}
	response["status"] = status
	response["data"] = data

	if err != nil{
		response["error"] = err.Error()
	}else{
		response["error"] = nil
	}
	return converter.ConvertToJSON(w,response)
}