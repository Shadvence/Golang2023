package main

import (
	"Assignments/internal/data"
	"Assignments/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string   `json:"title"`
		Year   int32    `json:"year"`
		Genres []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	book := &data.Book{
		Title:  input.Title,
		Year:   input.Year,
		Genres: input.Genres,
	}

	v := validator.New()

	if data.ValidateBook(v, book); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "All the Light We Cannot See",
		Author:    "Anthony Doerr",
		Genres:    []string{"novel"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
