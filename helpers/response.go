package helpers

import (
	"encoding/json"
	"net/http"
)

type ResposeWithData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "OK!"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = &ResposeWithData{
			Code:    code,
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    code,
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)

}
