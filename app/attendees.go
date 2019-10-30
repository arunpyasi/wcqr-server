package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/arunpyasi/wcqr-server/app/models"
	u "github.com/arunpyasi/wcqr-server/app/utils"
)

func getAttendees(w http.ResponseWriter, r *http.Request) {
	data := models.GetAttendees()

	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func addNewAttendee(w http.ResponseWriter, r *http.Request) {
	attendee := &models.Attendee{}
	err := json.NewDecoder(r.Body).Decode(attendee)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := attendee.Create()
	u.Respond(w, resp)
}

func getAttendee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	data := models.GetAttendee(ID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func updateAttendee(w http.ResponseWriter, r *http.Request) {
	attendee := &models.Attendee{}
	err := json.NewDecoder(r.Body).Decode(attendee)
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := attendee.Update(ID)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
