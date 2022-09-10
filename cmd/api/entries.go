//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// createEntryHandler for the "POST /v1/entries" endpoit
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry...")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the entry id
	fmt.Fprintf(w, "show the details for entries %d\n", id)
}
