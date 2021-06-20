package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/InterviewPortal/pkg/database"
	"github.com/vashish1/InterviewPortal/pkg/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	var res models.Response
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	if err != nil {
		res = models.Response{
			Success: false,
			Error:   err.Error(),
		}
		SendResponse(w, res, http.StatusBadRequest)
		return
	}
	ok, err := database.InsertUserData(data)
	if ok {
		res.Success = true
		SendResponse(w, res, http.StatusAccepted)
		return
	}
	res.Success = false
	res.Error = err.Error()
	SendResponse(w, res, http.StatusInternalServerError)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	ok, usersList := database.GetUserData()
	if ok {
		res.Success = true
		res.Data = usersList
		SendResponse(w, res, http.StatusOK)
		return
	}
	res.Success = false
	res.Error = "Error while fetching user list"
	SendResponse(w, res, http.StatusInternalServerError)
	return
}
