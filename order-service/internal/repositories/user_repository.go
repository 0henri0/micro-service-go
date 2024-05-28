package repositories

import (
	"database/sql"
	"micro-service-go/internal/models"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func (u userRepository) GetUsers() ([]models.User, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}
