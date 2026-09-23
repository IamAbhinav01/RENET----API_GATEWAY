package middlewares

import (
	"context"
	"net/http"
	dtos "renet/DTOs"
	"renet/utils/converter"
	"renet/utils/formatters"
	"renet/utils/validators"
)

var PayloadContextKet string = "payload"

func SignUpRequestValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dtos.SignupRequestDTO
		err := converter.DecodeFromJSON(r,&payload)
		if err != nil{
			formatters.ErrorResponse(w,http.StatusBadRequest,err,"Error occured while decoding the json")
			return 
		}

		validationErr := validators.Validate.Struct(payload)

		if validationErr != nil {
				formatters.ErrorResponse(w,http.StatusBadRequest,validationErr,"Invalid Request Payload")
				return 
		}

		reqContext := r.Context()
		ctx := context.WithValue(reqContext,PayloadContextKet,payload)

		next.ServeHTTP(w,r.WithContext(ctx))

	})
}