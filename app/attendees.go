package app

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(attendee)
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

	resp := u.Message(false, "")
	existing_data := models.GetAttendee(ID)
	if existing_data.AttendedAfterparty == true || existing_data.AttendedEvent == true {
		resp = u.Message(false, "duplicate entry")
		data := existing_data
		resp["data"] = data
	} else {
		resp = u.Message(true, "success")
		data := attendee.Update(ID)
		resp["data"] = data
	}
	u.Respond(w, resp)
}
