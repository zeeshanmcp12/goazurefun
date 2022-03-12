package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "GetName handler function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	text := "Greeting handler function executed successfully. Pass a hello in the query string for a personalized response.\n"
	queryParam := r.URL.Query().Get("hello")
	response := r.RequestURI

	if queryParam != "" {
		text = fmt.Sprintf("Hello, %s. This HTTP triggered function (greeting) executed successfully.\n%v", queryParam, response)
	}
	fmt.Fprint(w, text)
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	jsonData := []Profile{
		{"Zeeshan", 32, "acloudtechie.com"},
		{"Abdullah", 30, "abc.com"},
	}

	finalJSON, _ := json.MarshalIndent(jsonData, "", "\t")

	fmt.Println(string(finalJSON))

	w.Write([]byte(finalJSON))

}

type Profile struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Website string `json:"website"`
}

func CheckNilErr(err error) {
	if err != nil {
		fmt.Printf("Something wrong %s", err)
	}
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/GetName", helloHandler)
	http.HandleFunc("/api/Greeting", greeting)
	http.HandleFunc("/api/GetProfile", userProfile)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
