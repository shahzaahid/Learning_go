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
// var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (name, description, dateTime, user_id)
    VALUES($1, $2, $3, $4)`

	stmt, err := db.DB.Prepare(query)
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
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE ID = $1"

	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
        UPDATE events 
        SET name = $1 , description = $2, dateTime = $3
        WHERE id = $4
    `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.DateTime, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event Event) DeleteEventById() error {
	query := `DELETE FROM events WHERE ID = $1`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err

	//he return err statement will execute regardless of whether there is an error or not. If there is no error (err is nil), the function will return nil. If there is an error, it will be returned from the function.

}
