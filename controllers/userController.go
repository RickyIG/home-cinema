package controllers

import (
	database "home-cinema/database"
	model "home-cinema/models"
	"home-cinema/repository"
	"net/http"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func PostRegisterUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi untuk memastikan bahwa username, email dan phonenumber-nya baru
	// dengan cara mencari ke database, jika ditemukan maka mengembalikan error
	username, email, phone_number, err := repository.ValidateUser(database.DbConnection, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist!"})
		return
	}

	if email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already Used!"})
		return
	}

	if phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already Used!"})
		return
	}

	if email && phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Phone number already Used!"})
		return
	}

	// validasi password dengan cara teori kebalikan, jika sesuai dengan regex maka error
	compiledPassRegex, err := regexp.Compile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error compiling regex!",
		})
		return
	}

	if compiledPassRegex.MatchString(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character!",
		})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	user.RoleID = 1

	err = repository.InsertRegister(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "User registered successfully!",
	})
}

func PostRegisterAdmin(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, phone_number, err := repository.ValidateUser(database.DbConnection, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist!"})
		return
	}

	if email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already Used!"})
		return
	}

	if phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already Used!"})
		return
	}

	if email && phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Phone number already Used!"})
		return
	}

	// validasi password teori kebalikan, jika sesuai dengan regex maka error
	compiledPassRegex, err := regexp.Compile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error compiling regex!",
		})
		return
	}

	if compiledPassRegex.MatchString(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character!",
		})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	user.RoleID = 2

	// Validate email
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email invalid!"})
		return
	}

	err = repository.InsertRegister(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "User registered successfully!",
	})
}

func PostRegisterFintech(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, phone_number, err := repository.ValidateUser(database.DbConnection, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist!"})
		return
	}

	if email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already Used!"})
		return
	}

	if phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already Used!"})
		return
	}

	if email && phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Phone number already Used!"})
		return
	}

	// validasi password teori kebalikan, jika sesuai dengan regex maka error
	compiledPassRegex, err := regexp.Compile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error compiling regex!",
		})
		return
	}

	if compiledPassRegex.MatchString(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password must be minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character!",
		})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	user.RoleID = 3

	// Validate email
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email invalid!"})
		return
	}

	err = repository.InsertRegister(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "User registered successfully!",
	})
}

func PostLogin(c *gin.Context) {
	var credential model.Login

	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUserByUsername(database.DbConnection, credential.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.UserID,
		"username":  user.Username,
		"role":      user.RoleID,
		"ExpiresAt": time.Now().Add(5 * time.Minute).Unix(),
	})

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}

// func GetUserProfile(c *gin.Context) {
// 	// Retrieve user ID from context
// 	userID, _ := c.Get("userID")

// 	user, err := repository.GetUserProfileByID(database.DbConnection, userID.(int))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

func GetUserProfile(c *gin.Context) {

	userID, isExist := c.Get("userID")
	if !isExist {
		c.JSON(http.StatusBadRequest, "User doesn't exist.")
		return
	}

	user, err := repository.GetUserProfileByID(database.DbConnection, userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := gin.H{
		"result": user,
	}
	c.JSON(http.StatusOK, result)
}

func UpdateUserBalance(c *gin.Context) {
	var balanceRequest model.BalanceRequest
	if err := c.ShouldBindJSON(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id_user")

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		panic(err)
	}

	err = repository.UpdateUserBalance(database.DbConnection, balanceRequest.Balance, userIDint)

	if err != nil {
		// panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Balance",
	})
}
