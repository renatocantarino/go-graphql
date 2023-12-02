package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{DB: db}
}

func (c *Course) Create(name, description string) (Course, error) {
	id := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO courses (id, name ,description)  VALUES ($1,$2,$3)",
		id, name, description)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil

}
