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

func (c *Course) Create(name string, description string, categoryID string) (Course, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)",
		id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.DB.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var id string
		var name string
		var description string
		var categoryID string
		err := rows.Scan(&id, &name, &description, &categoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.DB.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var id string
		var name string
		var description string
		var categoryID string
		err := rows.Scan(&id, &name, &description, &categoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}
