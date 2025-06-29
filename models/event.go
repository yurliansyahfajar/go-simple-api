package models

import (
	"fmt"
	"time"

	"github.com/yurliansyahfajar/go-simple-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.UserID)

	if err != nil {
		return fmt.Errorf("error exec insert: %w", err)
	}

	eventId, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("error call last insert ID: %w", err)
	}

	e.ID = eventId
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error get all events from DB: %w", err)
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scan events rows from DB: %w", err)
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)
	if err != nil {
		return nil, fmt.Errorf("error scan events rows from DB: %w", err)
	}

	return &event, nil
}

func (e *Event) UpdateEvent() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Datetime, e.ID)

	return err
}

func (e *Event) DeleteEvent() error {
	query := `DELETE from events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err

}

func (e *Event) Register(userId int64) error {
	query := `INSTER INTO registration(user_id, event_id) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registration WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("prepare statement error: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
