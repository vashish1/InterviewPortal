package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func SendResponse(w http.ResponseWriter, data interface{}, code int) {
	b, _ := json.Marshal(data)
	w.WriteHeader(code)
	w.Write(b)
	return
}

func generateID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	return rand.Intn(max-min+1) + min
}

// func RequestParamCheck(){

// }
