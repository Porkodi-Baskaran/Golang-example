package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	config "example/StudentsDetails/Config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Dbconnection() // Ensure we use a test DB
}

func setupTestRouter() *gin.Engine {
	r := gin.Default()

	//  Add session middleware (Fix for "Key does not exist" error)
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/api/login", Login)
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
