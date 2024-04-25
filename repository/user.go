package repository

import (
	"database/sql"
	"errors"
	model "home-cinema/models"
	"time"
)

func InsertRegister(db *sql.DB, user model.User) (err error) {
	sql := "INSERT INTO users (username, password, first_name, last_name, email, phone_number, id_role) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	errs := db.QueryRow(sql, user.Username, user.Password, user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.RoleID)
	return errs.Err()
}

func ValidateUser(db *sql.DB, username, email, phone string) (isExistsUsername, isExistsEmail, isExistsPhoneNumber bool, err error) {

	var usernameCount, emailCount, phoneCount int

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&usernameCount)
	if err != nil {
		return false, false, false, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&emailCount)
	if err != nil {
		return false, false, false, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = $1", phone).Scan(&phoneCount)
	if err != nil {
		return false, false, false, err
	}

	isExistsUsername = usernameCount > 0
	isExistsEmail = emailCount > 0
	isExistsPhoneNumber = phoneCount > 0

	return isExistsUsername, isExistsEmail, isExistsPhoneNumber, nil
}

func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {
	query := "SELECT id_user, username, password ,id_role FROM users WHERE username = $1"
	user := &model.User{}
	err := db.QueryRow(query, username).Scan(&user.UserID, &user.Username, &user.Password, &user.RoleID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// untuk check
func GetUserById(db *sql.DB, id int) (exist bool, err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id_user = $1 AND id_role = 1", id).Scan(&count)
	if err != nil {
		return false, err
	}
	exist = count > 0
	return exist, nil
}

// func GetUserProfileByID(db *sql.DB, userID int) (*model.User, error) {
// 	sql := "SELECT * FROM users WHERE id_user = $1"

// 	rows, err := db.Query(sql, userID)
// 	if err != nil {

// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var user model.User

// 		err = rows.Scan(&user.UserID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.CreatedAt)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return &user, nil

// 	}
// 	return nil, err
// }

func GetUserProfileByID(db *sql.DB, userID int) (results []model.User, err error) {
	sql := "SELECT id_user, id_role, username, password, first_name, last_name, email, phone_number, balance, created_at FROM users WHERE id_user = $1"

	rows, err := db.Query(sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user = model.User{}

		err = rows.Scan(&user.UserID, &user.RoleID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Balance, &user.CreatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, user)
	}

	return
}

func UpdateUserBalance(db *sql.DB, balance int, userIDint int) (err error) {
	sql := "UPDATE users SET balance = balance + $1, updated_at = $2 WHERE id_user = $3"

	// formattedTime := film.JamTayang.Format("2006-01-02T00:00:00Z")

	errs := db.QueryRow(sql, balance, time.Now(), userIDint)

	return errs.Err()
}
