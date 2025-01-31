package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttpClient() {

	// c := http.Client{Timeout: time.Duration(1) * time.Second}
	// resp, err := c.Get("https://www.google.com")
	// if err != nil {
	// 	fmt.Printf("Error %s", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// fmt.Printf("Body : %s", body)

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Printf("Request Failed: %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}
	// Log the request body
	bodyString := string(body)
	log.Print(bodyString)
}
