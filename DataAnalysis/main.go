package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*
* Benutzt Mux um einen HttpRouter zu erstellen, der Anfragen mit einer Liste von verfügbaren Routen vergleicht
*
 */

// Gibt Testweise HelloWorld aus
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

/*
* StrictSlash defines the trailing slash behavior for new routes. The initial value is false.
* When true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa. In other words, your application will always see the path as specified in the route.
* When false, if the route path is "/path", accessing "/path/" will not match this route and vice versa.
 */

// Kümmert sich um eingehende Anfragen
func handleIncomingRequests() {
	// Erstellt neuen HttpRouter mit StrictSlash
	// Dies erlaubt dem Programm den Pfad der Route zu sehen
	router := mux.NewRouter().StrictSlash(true)
	// Der Standardpfad ruft die func helloWorld auf
	router.HandleFunc("/", helloWorld).Methods("GET")

	// Startet einen HttpServer mit Adresse und Handler bzw Router
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Starting ORM...")

	handleIncomingRequests()
}
