package database

import (
	"github.com/yama-koo/go-api-example/domain"
)

// UserRepository struct
type UserRepository struct {
	SQLHandler
}

// Store func
func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

// FindByID func
func (repo *UserRepository) FindByID(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}

	var id int
	var firstName string
	var lastName string

	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}

	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName

	return
}

// FindAll func
func (repo *UserRepository) FindAll() (user domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last+name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	var users domain.Users
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}
