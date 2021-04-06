package main

import (
	"fmt"
	"net/http"
)

/*
* - ResponseWriter Interface wird vom Router benutzt um eine Response zu erstellen
* - Request ist die eingehende Anfrage, die der Server empfängt
 */

func AllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Customer Endpoint")
}
