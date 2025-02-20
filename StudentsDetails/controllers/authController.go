package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	config "example/StudentsDetails/Config"

	"example/StudentsDetails/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// var (
// 	googleOauthConfig = &oauth2.Config{
// 		RedirectURL:  "http://localhost:8080/api/auth/google/callback",
// 		ClientID:     "507015603058-e86p59t5irp55eq4f9btgbflg5n4gjfb.apps.googleusercontent.com",
// 		ClientSecret: "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",
// 		Endpoint:     google.Endpoint,
// 	}
// )

type GoogleUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", creds.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user into database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var loginDetails models.LoginUser

	fmt.Println(&loginDetails)

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var storedHashedPassword string
	result := config.DB.QueryRow("SELECT password FROM users WHERE username = ?", loginDetails.Username)
	err := result.Scan(&storedHashedPassword)
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

	session := sessions.Default(c)
	session.Set("user_name", loginDetails.Username)
	session.Save()

	fmt.Println("loginDetails.Username:", loginDetails.Username)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("user_name")
	fmt.Println(userName)
	if userName == nil {
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

// func generateStateOauthCookie(c *gin.Context) string {
// 	var expiration = time.Now().Add(365 * 24 * time.Hour)
// 	state := "random-state-value" // Ideally, generate a unique random value
// 	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
// 	http.SetCookie(c.Writer, &cookie)
// 	return state
// }

// func HandleGoogleLogin(c *gin.Context) {
// 	state := generateStateOauthCookie(c)
// 	url := googleOauthConfig.AuthCodeURL(state)
// 	c.Redirect(http.StatusTemporaryRedirect, url)
// }

// func HandleGoogleCallback(c *gin.Context) {
// 	code := c.Query("code")
// 	fmt.Println("code:", code)
// 	token, err := googleOauthConfig.Exchange(context.Background(), code)
// 	fmt.Println("Token:", token)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
// 		return
// 	}

// 	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info"})
// 		return
// 	}

// 	var googleUser GoogleUser
// 	if err := json.Unmarshal(body, &googleUser); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
// 		return
// 	}

// 	session := sessions.Default(c)
// 	session.Set("user_name", googleUser.ID)
// 	session.Save()

// 	c.JSON(http.StatusOK, gin.H{"message": "User authenticated with Google", "user_id": googleUser.ID})
// }

// var creds Credentials
// if err := c.BindJSON(&creds); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 	return
// }
// var storedHashedPassword string
// err := config.DB.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username).Scan(&storedHashedPassword)
// if err != nil {
// 	if err == sql.ErrNoRows {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user"})
// 	return
// }

// if bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(creds.Password)) != nil {
// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 	return
// }

// expirationTime := time.Now().Add(5 * time.Minute)
// claims := &Claims{
// 	Username: creds.Username,
// 	StandardClaims: jwt.StandardClaims{
// 		ExpiresAt: expirationTime.Unix(),
// 	},
// }

// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// tokenString, err := token.SignedString(jwtKey)
// fmt.Println("Token", token)
// fmt.Println("Tokenstring:", tokenString)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 	return
// }

// http.SetCookie(c.Writer, &http.Cookie{
// 	Name:    "token",
// 	Value:   tokenString,
// 	Expires: expirationTime,
// })

// c.JSON(http.StatusOK, gin.H{"message": "Login successfully"})
