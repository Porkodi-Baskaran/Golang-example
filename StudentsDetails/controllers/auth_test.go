package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	r := gin.Default()

	//  Add session middleware (Fix for "Key does not exist" error)
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/api/login", Login)
	r.GET("api/students", AuthRequired, GetStudentDetails)
	r.GET("/api/students/:id", GetStudentDetailsbyID)
	r.POST("api/students/", CreateStudent)
	r.PUT("api/students/:id", UpdateStudent)
	r.DELETE("api/students/:id", DeleteStudent)
	return r
}

// Test user registration (Check if TOTP secret is generated)
func TestRegister(t *testing.T) {
	// Initialize a new Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate a request payload
	requestBody := `{"username":"testuser", "password":"Test@1234", "confirmPassword":"Test@1234"}`
	c.Request = httptest.NewRequest("POST", "/api/register", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the Register function
	Register(c)

	// Check if response status is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the response contains a QR code URL
	responseBody := w.Body.String()
	assert.Contains(t, responseBody, "qr_code")
}

// Test verifying a valid OTP
func TestVerifyTOTP(t *testing.T) {

	//  Generate an OTP using the exact same secret
	validOTP, _ := totp.GenerateCode("RYQS4O43ESNWHDK5U5CMIEJA26UQ42QQFYMWOFZ7CICGGXZFZWSQ", time.Now())
	fmt.Println("Generated OTP for Test:", validOTP)

	//  Create a test request
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Send request with correct OTP
	requestBody := `{"username":"testuser", "totp":"` + validOTP + `"}`
	c.Request = httptest.NewRequest("POST", "/api/verify-2fa", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call VerifyTOTP function
	VerifyTOTP(c)

	//  Check if response is 200 OK
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 but got HTTP %d", w.Code)

	//  Check success message
	assert.Contains(t, w.Body.String(), "2FA enabled successfully", "Response should contain success message")
}

// Test login with correct username, password, and OTP
func TestLogin(t *testing.T) {

	// Generate a valid OTP
	validOTP, _ := totp.GenerateCode("RYQS4O43ESNWHDK5U5CMIEJA26UQ42QQFYMWOFZ7CICGGXZFZWSQ", time.Now())
	router := setupTestRouter()

	// Initialize a new Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate a request payload with correct credentials
	requestBody := `{"username":"testuser", "password":"Test@1234", "totp":"` + validOTP + `"}`
	c.Request = httptest.NewRequest("POST", "/api/login", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, c.Request)
	// // Call the Login function
	// Login(c)

	// Check if response status is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response message
	assert.Contains(t, w.Body.String(), "Login successfully")
}

// Test login with incorrect OTP
func TestLoginIncorrectOTP(t *testing.T) {

	router := setupTestRouter()
	// Initialize a new Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate a request payload with correct password but wrong OTP
	requestBody := `{"username":"testuser", "password":"Test@1234", "totp":"123456"}`
	c.Request = httptest.NewRequest("POST", "/api/login", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the Login function
	// Login(c)
	router.ServeHTTP(w, c.Request)

	// Check if response status is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check response message
	assert.Contains(t, w.Body.String(), "Invalid OTP")
}

// Test login with incorrect password
func TestLoginIncorrectPassword(t *testing.T) {

	// Generate a valid OTP
	validOTP, _ := totp.GenerateCode("RYQS4O43ESNWHDK5U5CMIEJA26UQ42QQFYMWOFZ7CICGGXZFZWSQ", time.Now())

	// Insert a test user into the database
	// _, _ = config.DB.Exec("INSERT INTO users (username, password, totp_secret) VALUES (?, ?, ?)",
	// 	"testuser", string(hashedPassword), key.Secret())
	router := setupTestRouter()
	// Initialize a new Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate a request payload with incorrect password but correct OTP
	requestBody := `{"username":"testuser", "password":"WrongPass@123", "totp":"` + validOTP + `"}`
	c.Request = httptest.NewRequest("POST", "/api/login", strings.NewReader(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the Login function
	// Login(c)
	router.ServeHTTP(w, c.Request)

	// Check if response status is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Check response message
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}

func TestGetStudentDetails(t *testing.T) {
	router := setupTestRouter()
	// // Initialize a new Gin context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)

	// _, _ = Config.DB.Exec("INSERT INTO student (Name, Class, Address) VALUES (?, ?, ?)", "John Doe", 10, "123 Main St")

	// // Simulate a request payload with correct password but wrong OTP
	// requestBody := `{"id":1, "Name":"Tom", "Class":"9", "Address":"Chennai"}`
	req := httptest.NewRequest("GET", "/api/students", nil)
	// strings.NewReader(requestBody))
	// c.Request.Header.Set("Content-Type", "application/json")
	c, _ := gin.CreateTestContext(w)
	store := cookie.NewStore([]byte("secret"))
	session := sessions.Sessions("mysession", store)
	c.Request = req
	session(c)
	s := sessions.Default(c)
	s.Set("user_id", 1) // Simulating a logged-in user
	s.Save()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 but got %d", w.Code)
	assert.Contains(t, w.Body.String(), "data", "Response should contain student data")
}

func TestGetStudentDetailsbyID(t *testing.T) {
	router := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/students/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 but got %d", w.Code)
	assert.Contains(t, w.Body.String(), "data", "Response should contain student name")

}

func TestCreateStudent(t *testing.T) {
	router := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/students/",
		strings.NewReader(`{"Name":"Jane Doe", "Class":"12", "Address":"456 Oak St"}`))
	fmt.Println(req)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected HTTP 201 but got %d", w.Code)
	assert.Contains(t, w.Body.String(), "data", "Response should confirm student creation")
}

func TestUpdateStudent(t *testing.T) {
	router := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/api/students/24",
		strings.NewReader(`{"Name":"Jane Doe1", "Class":"5", "Address":"123 Oak St"}`))
	fmt.Println(req)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 but got %d", w.Code)
	assert.Contains(t, w.Body.String(), "data", "Response should confirm student updated")
}

func TestDeleteStudent(t *testing.T) {
	router := setupTestRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/api/students/24", nil)
	fmt.Println(req)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 but got %d", w.Code)
	assert.Contains(t, w.Body.String(), "Student deleted successfully", "Response should confirm student deletion")
}
