package models

import "time"

// ? Declar the Event
type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

// ? Made the variable called events it is the type of Event
var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

// ? this function is used to return all the events
func GetAllEvents() []Event {
	return events
}
