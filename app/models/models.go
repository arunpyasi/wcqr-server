package models

import (
	"fmt"

	u "github.com/openarun/wcqr-server/app/utils"
	"github.com/jinzhu/gorm"
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
	if err := GetDB().Model(&attendee).Where("id=?", id).First(attendee).Error; err != nil {
		fmt.Println(err)
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
	}
	return attendee
}

func GetAttendees() []*Attendee {
	attendees := []*Attendee{}
	if err := GetDB().Find(&attendees).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
	}
	return attendees

}
func (attendee *Attendee) Update(id int) *Attendee {
	if err := GetDB().Model(&attendee).Where("id=?", id).Updates(attendee).Error; err != nil {
		fmt.Println(err)
	}
	attendee = GetAttendee(id)
	return attendee
}
