package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"home-cinema/database"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: database.DbConnection}
}

// GetBalance retrieves the user's balance based on their ID.
func (repo *UserRepository) GetBalance(userID int) (int, error) {
	var balance int
	fmt.Println("cek userid", userID)

	row := repo.db.QueryRow(`SELECT balance FROM users WHERE id_user = $1`, userID)
	err := row.Scan(&balance)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return balance, nil
}

// UpdateBalance updates the user's balance with the provided new balance within a transaction.
func (repo *UserRepository) UpdateBalance(tx *sql.Tx, userID int, newBalance int) error {
	_, err := tx.Exec(`UPDATE users SET balance = $1 WHERE id_user = $2`, newBalance, userID)
	if err != nil {
		return err
	}

	return nil
}
