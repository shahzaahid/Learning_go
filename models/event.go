package models

import (
	"time"

	"example.com/test/db"
)

// ? Declar the Event
type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	// Location    string `binding:"required"`
	DateTime time.Time
	UserId   int
}

// ? Made the variable called events it is the type of Event
var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (name, description, dateTime, user_id)
	VALUES($1, $2, $3, $4)`

	stmt, err := db.DB.Prepare(query)
	// events = append(events, e)
	if err != nil {

		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.UserId)
	if err != nil {

		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return err

}

// ? this function is used to return all the events
// func GetEventByName(name string) (Event, error) {
// 	query := "SELECT * FROM events WHERE name = $1 LIMIT 1"
// 	row := db.DB.QueryRow(query, name)

// 	var event Event
// 	err := row.Scan(&event.Name, &event.Description, &event.DateTime, &event.UserId)

// 	if err != nil {
// 		return Event{}, err
// 	}

// 	return event, nil
// }

func GetAllEvents() []Event {
	return events
}
