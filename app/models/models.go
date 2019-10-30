package models

import (
	"fmt"

	u "github.com/arunpyasi/wcqr-server/app/utils"
)

type Attendee struct {
	ID                 uint   `json:"id" gorm:"primary_key"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	Email              string `json:"email"`
	AttendedEvent      bool   `json:"attended_event"`
	AttendedAfterparty bool   `json:"attended_afterparty"`
}

func (attendee *Attendee) Create() map[string]interface{} {

	GetDB().Create(attendee)
	resp := u.Message(true, "success")
	resp["data"] = attendee
	return resp
}

func GetAttendee(id int) *Attendee {
	attendee := &Attendee{}
	GetDB().Model(&attendee).Where("id=?", id).First(attendee)
	return attendee
}

func GetAttendees() []*Attendee {
	attendees := []*Attendee{}
	GetDB().Find(&attendees)
	fmt.Println(attendees)
	return attendees

}
func (attendee *Attendee) Update(id int) *Attendee {
	GetDB().Model(&attendee).Where("id=?", id).Updates(attendee)
	fmt.Println(attendee)
	attendee = GetAttendee(id)
	return attendee
}
