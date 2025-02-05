package main

// import (
// 	"crypto/tls"
// 	"example/https"
// 	"log"
// 	"net/http"
// )

// var (
// 	CertFilePath = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/server-cert.pem"
// 	KeyFilePath  = "C:/Users/dell/OneDrive/Desktop/go/src/TLS certificate/server-key.pem"
// )

// func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
// 	w.Write([]byte("Hello,World!\n"))
// }
// func main() {
// 	// basic.CookieExample()

// 	// basic.SessionCookiefunc()
// 	https.HttpsClient("https://localhost:8080")

// 	// simple https server

// 	// load tls certificates
// 	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
// 	if err != nil {
// 		log.Fatalf("Error loading certificate and key file: %v", err)
// 	}

// 	tlsConfig := &tls.Config{
// 		Certificates: []tls.Certificate{serverTLSCert},
// 	}
// 	server := http.Server{
// 		Addr:      ":4443",
// 		Handler:   http.HandlerFunc(httpRequestHandler),
// 		TLSConfig: tlsConfig,
// 	}
// 	defer server.Close()
// 	log.Fatal(server.ListenAndServeTLS("", ""))
// }

// // "example/http"

// // // Define the interface
// // type Shape interface {
// // 	Area() float64
// // 	Perimeter() float64
// // }

// // // Circle type that implements the Shape interface
// // type Circle struct {
// // 	radius float64
// // }

// // // Rectangle type that implements the Shape interface
// // type Rectangle struct {
// // 	length, width float64
// // }

// // func (c Circle) Area() float64 {
// // 	return math.Pi * c.radius * c.radius
// // }

// // func (c Circle) Perimeter() float64 {
// // 	return 2 * math.Pi * c.radius
// // }

// // func (r Rectangle) Area() float64 {
// // 	return r.length * r.width
// // }

// // func (r Rectangle) Perimeter() float64 {
// // 	return 2 * (r.length + r.width)
// // }

// // // Function to determine the type of shape and calculate area
// //
// //	func calculateArea(Shape interface{}) {
// //		switch s := Shape.(type) {
// //		case Circle:
// //			fmt.Printf("Circle area: %.2f\n", s.Area())
// //		case Rectangle:
// //			fmt.Printf("Rectangle area: %.2f\n", s.Area())
// //		default:
// //			fmt.Println("Unknown shape")
// //		}
// //	}
// //----http server and client code
// // const keyServerAddr = "serverAddr"

// // func getRoot(w http.ResponseWriter, r *http.Request) {
// // 	ctx := r.Context()

// // 	hasFirst := r.URL.Query().Has("first")
// // 	first := r.URL.Query().Get("first")
// // 	hasSecond := r.URL.Query().Has("second")
// // 	second := r.URL.Query().Get("second")

// // 	body, err := ioutil.ReadAll(r.Body)
// // 	if err != nil {
// // 		fmt.Printf("could not read body: %s\n", err)
// // 	}

// // 	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
// // 		ctx.Value(keyServerAddr),
// // 		hasFirst, first,
// // 		hasSecond, second, body)

// // 	io.WriteString(w, "This is my website!\n")
// // }
// // func postHello(w http.ResponseWriter, r *http.Request) {
// // 	ctx := r.Context()
// // 	fmt.Println("context:", ctx)
// // 	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
// // 	myName := r.PostFormValue("myName")

// // if myName == "" {
// // 	myName = "HTTP"
// // }
// // 	if myName == "" {
// // 		w.Header().Set("x-missing-field", "myName")
// // 		w.WriteHeader(http.StatusBadRequest)
// // 		return

// // 	}
// // 	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
// // }

// // func headers(w http.ResponseWriter, req *http.Request) {
// // 	for name, headers := range req.Header {
// // 		for _, h := range headers {
// // 			fmt.Println("Sending response")
// // 			fmt.Println(w, "%v: %v\n", name, h)
// // 		}
// // 	}
// // }

// // JWT token code start
// // var secretKey = []byte("secret-key")

// // func main() {
// // 	router := mux.NewRouter()

// // 	router.HandleFunc("/login", LoginHandler).Methods("POST")
// // 	router.HandleFunc("/protected", ProtectedHandler).Methods("GET")

// // 	fmt.Println("Starting the server")
// // 	err := http.ListenAndServe("localhost:8080", router)
// // 	if err != nil {
// // 		fmt.Println("Could not start the server", err)
// // 	}
// // }

// // func createToken(username string) (string, error) {
// // 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// // 		jwt.MapClaims{
// // 			"username": username,
// // 			"exp":      time.Now().Add(time.Hour * 24).Unix(),
// // 		})
// // 	tokenString, err := token.SignedString(secretKey)
// // 	if err != nil {
// // 		return "", nil
// // 	}
// // 	return tokenString, nil
// // }

// // func verifyToken(tokenString string) error {
// // 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// // 		return secretKey, nil
// // 	})
// // 	if err != nil {
// // 		return err
// // 	}
// // 	if !token.Valid {
// // 		return fmt.Errorf("invalid token")
// // 	}
// // 	return nil
// // }

// // type User struct {
// // 	Username string `json:"username"`
// // 	Password string `json:"password"`
// // }

// // // POST
// // func LoginHandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	var u User
// // 	json.NewDecoder(r.Body).Decode(&u)
// // 	fmt.Printf("The user request value %v", u)

// // 	if u.Username == "Aishu" && u.Password == "123456" {
// // 		tokenString, err := createToken(u.Username)
// // 		if err != nil {
// // 			w.WriteHeader(http.StatusInternalServerError)
// // 			fmt.Println("no username found")
// // 		}
// // 		w.WriteHeader(http.StatusOK)
// // 		fmt.Fprint(w, tokenString)
// // 		return
// // 	} else {
// // 		w.WriteHeader(http.StatusUnauthorized)
// // 		fmt.Fprint(w, "Invalid credentials")
// // 	}
// // }

// // // GET
// // func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	tokenString := r.Header.Get("Authorization")
// // 	if tokenString == "" {
// // 		w.WriteHeader(http.StatusUnauthorized)
// // 		fmt.Fprint(w, "Missing authorization header")
// // 		return
// // 	}
// // 	tokenString = tokenString[len("Bearer "):]
// // 	// fmt.Println(tokenString)

// // 	err := verifyToken(tokenString)
// // 	if err != nil {
// // 		w.WriteHeader(http.StatusUnauthorized)
// // 		fmt.Fprint(w, "Invalid token")
// // 		return
// // 	}

// // 	fmt.Fprint(w, "Welcome to the the protected area")

// // }

// // JWT token code ends

// //Http server and client code with get and post
// // mux := http.NewServeMux()
// // mux.HandleFunc("/", getRoot)
// // // mux.HandleFunc("GET /hello", getHello)
// // mux.HandleFunc("POST /hello", postHello)

// // ctx := context.Background()
// // server := &http.Server{
// // 	Addr:    ":8080",
// // 	Handler: mux,
// // 	BaseContext: func(l net.Listener) context.Context {
// // 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
// // 		return ctx
// // 	},
// // }
// // serverTwo := &http.Server{
// // 	Addr:    ":7034",
// // 	Handler: mux,
// // 	BaseContext: func(l net.Listener) context.Context {
// // 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
// // 		return ctx
// // 	},
// // }

// // err := server.ListenAndServe()
// // if errors.Is(err, http.ErrServerClosed) {
// // 	fmt.Printf("server closed\n")
// // } else if err != nil {
// // 	fmt.Printf("error listening for server: %s\n", err)
// // }
// // go func() {
// // 	err := serverTwo.ListenAndServe()
// // 	if errors.Is(err, http.ErrServerClosed) {
// // 		fmt.Printf("server two closed\n")
// // 	} else if err != nil {
// // 		fmt.Printf("error listening for server two: %s\n", err)
// // 	}
// // 	cancelCtx()
// // }()

// // <-ctx.Done()
// // http.HandleFunc("/hello", hello)
// // http.HandleFunc("/headers", headers)

// // http.ListenAndServe(":8080", nil)

// // http.GinAPIServer()
// // http.HttpClient()
// // http.Httpserverfunc()
// // while select  list  map  interface  struct  list()struct
// // how to create simple http server without using gin also Http client
// // Http handler without gin
// // jwt  cookie https server
// //
// // basic.BasicFn()
// // basic.Gochannel()
// // basic.Goselectfun()
// // basic.GoMapfun()
// // basic.Arrayfunc()
// // basic.Slicefunc()
// // basic.Switchcasefunc()
// // basic.Structfunc()
// // basic.Listfunc()
// // basic.Liststructfunc()
// // basic.Interfacefunc()

// // c := Circle{radius: 5}

// // r := Rectangle{length: 4, width: 3}
// // calculateArea(c)
// // calculateArea(r)
