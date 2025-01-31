package main

import (
	// "example/http"

	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

// // Define the interface
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// // Circle type that implements the Shape interface
// type Circle struct {
// 	radius float64
// }

// // Rectangle type that implements the Shape interface
// type Rectangle struct {
// 	length, width float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.length + r.width)
// }

// // Function to determine the type of shape and calculate area
//
//	func calculateArea(Shape interface{}) {
//		switch s := Shape.(type) {
//		case Circle:
//			fmt.Printf("Circle area: %.2f\n", s.Area())
//		case Rectangle:
//			fmt.Printf("Rectangle area: %.2f\n", s.Area())
//		default:
//			fmt.Println("Unknown shape")
//		}
//	}
const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
		ctx.Value(keyServerAddr),
		hasFirst, first,
		hasSecond, second, body)

	io.WriteString(w, "This is my website!\n")
}
func postHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("context:", ctx)
	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	myName := r.PostFormValue("myName")

	// if myName == "" {
	// 	myName = "HTTP"
	// }
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
}

// func headers(w http.ResponseWriter, req *http.Request) {
// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Println("Sending response")
// 			fmt.Println(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	// mux.HandleFunc("GET /hello", getHello)
	mux.HandleFunc("POST /hello", postHello)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}
	// serverTwo := &http.Server{
	// 	Addr:    ":7034",
	// 	Handler: mux,
	// 	BaseContext: func(l net.Listener) context.Context {
	// 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
	// 		return ctx
	// 	},
	// }

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
	// go func() {
	// 	err := serverTwo.ListenAndServe()
	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Printf("server two closed\n")
	// 	} else if err != nil {
	// 		fmt.Printf("error listening for server two: %s\n", err)
	// 	}
	// 	cancelCtx()
	// }()

	<-ctx.Done()
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	// http.ListenAndServe(":8080", nil)

	// http.GinAPIServer()
	// http.HttpClient()
	// http.Httpserverfunc()
	// while select  list  map  interface  struct  list()struct
	// how to create simple http server without using gin also Http client
	// Http handler without gin
	// jwt  cookie https server
	//
	// basic.BasicFn()
	// basic.Gochannel()
	// basic.Goselectfun()
	// basic.GoMapfun()
	// basic.Arrayfunc()
	// basic.Slicefunc()
	// basic.Switchcasefunc()
	// basic.Structfunc()
	// basic.Listfunc()
	// basic.Liststructfunc()
	// basic.Interfacefunc()

	// c := Circle{radius: 5}

	// r := Rectangle{length: 4, width: 3}
	// calculateArea(c)
	// calculateArea(r)

}
