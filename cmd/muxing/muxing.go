package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeader).Methods(http.MethodPost)

	//http.NewServeMux()
	//log.Fatalln(http.ListenAndServe(":8081", router))
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleName(w http.ResponseWriter, r *http.Request) {
	name := "mister X"
	if p, ok := mux.Vars(r)["PARAM"]; ok {
		name = p
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

}

func handleData(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "I got massage:\n"+string(str))

}

func handleHeader(w http.ResponseWriter, r *http.Request) {
	x := r.Header.Get("a")
	y := r.Header.Get("b")

	a, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Add("a+b", strconv.Itoa(a+b))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
