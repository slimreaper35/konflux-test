package models

import (
	"time"

	"github.com/slimreaper35/konflux-test/database"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int64     `binding:"required"`
}

func (e *Event) Create() error {
	var query = `
	INSERT INTO events(name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?);
	`
	result, err := database.DB.Exec(query, e.Name, e.Description, e.Location, e.Datetime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func (e *Event) Update() error {
	var query = `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?;
	`
	_, err := database.DB.Exec(query, e.Name, e.Description, e.Location, e.Datetime, e.ID)
	return err
}

func (e *Event) Delete() error {
	var query = `
	DELETE FROM events
	WHERE id = ?;
	`
	_, err := database.DB.Exec(query, e.ID)
	return err
}

func GetEvents() ([]Event, error) {
	var query = `
	SELECT * FROM events;
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return []Event{}, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		events = append(events, event)
	}

	return events, nil
}

func GetEventBy(id int64) (*Event, error) {
	var query = `
	SELECT * FROM events
	WHERE id = ?;
	`
	var row = database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) RegisterUser(userID int64) error {
	var query = `
	INSERT INTO registrations(event_id, user_id)
	VALUES (?, ?);
	`
	_, err := database.DB.Exec(query, e.ID, userID)
	return err
}

func (e *Event) UnregisterUser(userID int64) error {
	var query = `
	DELETE FROM registrations
	WHERE event_id = ? AND user_id = ?;
	`
	_, err := database.DB.Exec(query, e.ID, userID)
	return err
}
