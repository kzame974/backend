package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := loginRequest.Username
	password := loginRequest.Password

	db, err := sql.Open("postgres", "postgres://user:password@localhost/database?sslmode=disable")
	if err != nil {
		logrus.WithError(err).Error("Error connecting to database")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT id, username, email, password FROM users WHERE username=$1", username)
	user := User{}
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		logrus.WithError(err).Error("Error querying database")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if user.Password != password {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func main() {
	router := gin.Default()
	router.POST("/login", Login)
	router.Run()
}
