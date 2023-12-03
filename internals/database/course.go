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
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{DB: db}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO courses (id, name ,description ,category_id )  VALUES ($1,$2,$3,$4)",
		id, name, description, categoryID)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil

}
