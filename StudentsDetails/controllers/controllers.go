package controllers

import (
	"example/StudentsDetails/models"
	"example/StudentsDetails/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudMarks(c *gin.Context) {
	marks, err := repositories.GetStudentMarks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": marks})
}
func GetStudMarksbyID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Student ID"})
		return
	}
	marks, err := repositories.GetMarksbyID(id)
	if err != nil {
		if err.Error() == "student not found" {
			fmt.Println(id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": marks})
}

func GetStudentDetails(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	students, err := repositories.GetStudentDetails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": students})
}

func GetStudentDetailsbyID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Student ID"})
		return
	}
	student, err := repositories.GetStudentDetailsbyID(id)
	if err != nil {
		if err.Error() == "student not found" {
			fmt.Println(id)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateStudent(c *gin.Context) {
	var student models.StudentDetails
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": student})
}

func UpdateStudent(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedStudent models.StudentDetails
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedStudent.ID = id
	if err := repositories.UpdateStudent(updatedStudent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedStudent})
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Student ID"})
	// 	return
	// }
	// var student models.StudentDetails
	// if err := c.ShouldBindJSON(&student); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
}

func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	if err := repositories.DeleteStudent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted successfully"})
}

// func Register(c *gin.Context) {
// 	var user models.LoginUser
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

// 	fmt.Println(hashedPassword)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 		return
// 	}
// 	user.Password = string(hashedPassword)

// error := repositories.RegisterUser(user)
// 	// query := "INSERT INTO users (username, password) VALUES (?, ?)"
// 	// _, err = DB.Dbconnection().Exec(query, user.Username, user.Password)
// 	if error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

// func Login(c *gin.Context) {
// 	var user models.LoginUser
// 	var loginVals models.LoginUser
// 	if err := c.ShouldBindJSON(&loginVals); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// err := repositories.LoginUser(user, loginVals)
// 	query := "SELECT id, username, password FROM users WHERE username = ?"
// 	row := config.DB.QueryRow(query, loginVals.Username)
// 	if err := row.Scan(&user.UID, &user.Username, &user.Password); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
// }
