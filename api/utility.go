package api

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, code int) {
	b, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(b)
	return
}

// func RequestParamCheck(){

// }