package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	config "example/StudentsDetails/Config"

	"example/StudentsDetails/models"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/auth/google/callback",
		ClientID:     "507015603058-e86p59t5irp55eq4f9btgbflg5n4gjfb.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

type GoogleUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Credentials struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

const jwtSecret = "secret_key_aishu"

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if creds.Password != creds.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	totpKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: creds.Username,
		SecretSize:  32,
	})

	fmt.Println("Key:", totpKey.Secret())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP secret"})
		return
	}

	// encodedSecret := base32.StdEncoding.EncodeToString([]byte(totpKey.Secret()))

	// fmt.Println("Encoded Secret:", encodedSecret)

	// _, err = config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, string(hashedPassword))
	_, err = config.DB.Exec("INSERT INTO users (username, password,totp_secret) VALUES (?, ?, ?)",
		creds.Username, string(hashedPassword), totpKey.Secret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user into database"})
		return
	}

	fmt.Println("URL:", totpKey.URL())

	// Generate QR Code
	// qrCode, _ := qrcode.Encode(totpKey.URL(), qrcode.Medium, 256)
	// qrBase64 := base64.StdEncoding.EncodeToString(qrCode)

	// c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	c.JSON(http.StatusOK, gin.H{
		"message":     "User registered successfully",
		"qr_code":     totpKey.URL(),
		"totp_secret": totpKey.Secret(),
	})

}

// Verify OTP and enable 2FA for new user
func VerifyTOTP(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		TOTP     string `json:"totp"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Get the stored TOTP secret from DB
	var storedSecret string
	err := config.DB.QueryRow("SELECT totp_secret FROM users WHERE username = ?", req.Username).Scan(&storedSecret)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	fmt.Println(" RequestedTOTP:", req.TOTP)
	fmt.Println("Stored TOTP:", storedSecret)

	// Decode Base32 Secret (Google Authenticator format)
	storedSecret = strings.TrimRight(storedSecret, "=") // Fix padding
	storedSecret = strings.ToUpper(storedSecret)        // Ensure uppercase
	fmt.Println("DecodedSecret", storedSecret)
	// Verify OTP
	valid := totp.Validate(req.TOTP, storedSecret)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "2FA enabled successfully"})
}

func Login(c *gin.Context) {
	var loginDetails models.LoginUser

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var storedHashedPassword, storedTOTP string
	result := config.DB.QueryRow("SELECT password,totp_secret FROM users WHERE username = ?", loginDetails.Username)
	err := result.Scan(&storedHashedPassword, &storedTOTP)
	fmt.Println("User Query Result:", result)
	// err := result.Scan(&storedHashedPassword, &storedTOTPSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(loginDetails.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Verify OTP
	valid := totp.Validate(loginDetails.TOTP, storedTOTP)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	//JWT token implementation
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256,
	// 	jwt.MapClaims{
	// 		"username": loginDetails.Username,
	// 		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	// 	})
	// tokenString, err := token.SignedString(jwtSecret)

	// if err != nil {
	// 	fmt.Println("Error occur while implementing JWT ")
	// }

	session := sessions.Default(c)
	// session.Clear() //Clear any existing session datas
	// session.Set("user_name", loginDetails.Username)
	session.Set("user_id", loginDetails.UID)
	session.Set("user_name", loginDetails.Username)

	session.Save()

	fmt.Println("loginDetails.Username:", loginDetails.Username)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}

func JwtLogin(c *gin.Context) {
	var loginDetails models.LoginUser

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	const LoginUsername = "nn"
	var storedHashedPassword, storedTOTP string
	result := config.DB.QueryRow("SELECT password,totp_secret FROM users WHERE username =?", LoginUsername)
	err := result.Scan(&storedHashedPassword, &storedTOTP)
	fmt.Println("User Query Result:", result)
	// err := result.Scan(&storedHashedPassword, &storedTOTPSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
		return
	}

	// if bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(loginDetails.Password)) != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	// Verify OTP
	// valid := totp.Validate(loginDetails.TOTP, storedTOTP)
	// if !valid {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	// 	return
	// }

	expirationTime := time.Now().Add(time.Hour * 1)
	claims := &Claims{
		Username: loginDetails.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
			c.Abort()
			return
		}
		fmt.Println("TokenString:", tokenString)
		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		fmt.Println("Token:", token)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store username from token in context
		c.Set("username", "nn")
		c.Next()
	}
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	fmt.Println("session userId:", userID)
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

func LogoutUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func HandleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c *gin.Context) {
	// Get authorization code from Google
	code := c.Query("code")

	// Exchange authorization code for an access token
	token, err := googleOauthConfig.Exchange(context.Background(), code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}
	// Retrieve user info from Google
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		log.Println("Failed to get user info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	fmt.Println("Response from callback:", resp)
	defer resp.Body.Close()

	// Decode the response
	var googleUser struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		log.Println("Failed to parse user info:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	// Store user in database if they don't exist
	var userID int
	err = config.DB.QueryRow("SELECT id FROM users WHERE username = ?", googleUser.Email).Scan(&userID)
	if err == sql.ErrNoRows {
		_, err = config.DB.Exec("INSERT INTO users (username) VALUES (?)", googleUser.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
			return
		}
	}

	// Set session
	session := sessions.Default(c)
	// session.Clear()                      //Clear the existing session data if any
	session.Set("auth_method", "google") // Store auth method
	// session.Set("email", googleUser.Email)
	session.Set("user_id", googleUser.ID)
	session.Save()

	// Redirect to frontend students page
	c.Redirect(http.StatusFound, "http://localhost:8080/students")
	// c.JSON(http.StatusOK, gin.H{"message": "Login successful", "redirectUrl": "http://localhost:8080/students"})

}
