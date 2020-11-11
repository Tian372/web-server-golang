package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tian372/serverTest/services"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "zt45"
	password = "zt45duke"
	dbname   = "web_dev"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	//testing code
	fmt.Fprintln(w, "<body><h1>Test Home Page</h1><h2>H2</h2><h3>H3</h3> <h4>H4</h4> <h5>H5</h5> <h6>H6</h6></body>")
}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	user := services.NewUser(psqlInfo)
	router := mux.NewRouter()

	router.HandleFunc("/", testHandle).Methods("GET")
	router.HandleFunc("/data", user.Create).Methods("GET")

	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}

}
