//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createEntryHandler for the "POST /v1/entries" endpoit
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new entry...")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Display the entry id
	fmt.Fprintf(w, "show the details for school %d\n", id)
}
