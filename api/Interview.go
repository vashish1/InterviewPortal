package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/vashish1/InterviewPortal/pkg/database"
	"github.com/vashish1/InterviewPortal/pkg/models"
)

type input struct {
	Id           int           `json:"id,omitempty"`
	StartTime    string        `json:"start_time,omitempty"`
	EndTime      string        `json:"end_time,omitempty"`
	Participants []models.User `json:"participants,omitempty"`
}

func AddInterview(w http.ResponseWriter, r *http.Request) {
	var data input
	var res models.Response
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	if err != nil {
		res.Success = false
		res.Error = err.Error()
		SendResponse(w, res, http.StatusBadRequest)
		return
	}
	if len(data.Participants) < 2 {
		res.Success = false
		res.Error = "There must be at least 2 participants"
		SendResponse(w, res, http.StatusNotAcceptable)
		return
	}

	start, err := time.Parse("2006-01-02T15:04:00Z", data.StartTime)
	fmt.Println(err)
	end, err := time.Parse("2006-01-02T15:04:00Z", data.EndTime)
	fmt.Println(err)
	if start.Before(time.Now()) && end.After(time.Now()) {
		res.Success = false
		res.Error = "Interview cannot be scheduled for past time"
		SendResponse(w, res, http.StatusBadRequest)
		return
	}
	for _, user := range data.Participants {
		ok := database.CheckAvailability(user, start, end)
		if !ok {
			res.Success = false
			res.Error = user.Name + "is not availabe on required time slot"
			SendResponse(w, res, http.StatusNotAcceptable)
			return
		}
	}
	Interviewdata := models.Interview{
		StartTime:    start,
		EndTime:      end,
		Participants: data.Participants,
	}
	if ok, err := database.InsertInterviewDetails(Interviewdata); !ok {
		res.Success = false
		res.Error = err.Error()
		SendResponse(w, res, http.StatusInternalServerError)
		return
	}
	res.Success = true
	SendResponse(w, res, http.StatusOK)
	return
}

func EditInterview(w http.ResponseWriter, r *http.Request) {
	var data input
	var res models.Response
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	if err != nil {
		res.Success = false
		res.Error = err.Error()
		SendResponse(w, res, http.StatusBadRequest)
		return
	}
	start, err := time.Parse("2006-01-02T15:04:00Z", data.StartTime)
	fmt.Println(err)
	end, err := time.Parse("2006-01-02T15:04:00Z", data.EndTime)
	Interviewdata := models.Interview{
		StartTime:    start,
		EndTime:      end,
		Participants: data.Participants,
	}
	ok, err := database.UpdateInterview(Interviewdata)
	if ok {
		res.Success = true
		SendResponse(w, res, http.StatusOK)
		return
	}
	res.Success = false
	res.Error = err.Error()
	SendResponse(w, res, http.StatusBadRequest)
	return
}

func GetInterviewList(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	ok, List := database.GetInterviews()
	if ok {
		res.Success = true
		res.Data = List
		SendResponse(w, res, http.StatusOK)
		return
	}
	res.Success = false
	res.Error = "Error while fetching list"
	SendResponse(w, res, http.StatusInternalServerError)
	return
}
