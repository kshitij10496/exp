package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := chi.NewRouter()
	router.Get("/", handleIndex())
	router.Get("/contact/1/edit", handleContactEdit())
	router.Put("/contact/1", handleContactPut())

	// serve the static folder with JS.
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := ":8080"
	logger.Info("serving HTTP traffic", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("serving HTTP traffic: %v", err)
	}
}

func handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := `
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="static/js/htmx.min.js"></script>
</head>
<body>
    <main>
        <div>
            Hello World!
        </div>

        <div hx-target="this" hx-swap="outerHTML">
            <div><label>First Name</label>: Joe</div>
            <div><label>Last Name</label>: Blow</div>
            <div><label>Email</label>: joe@blow.com</div>
            <button hx-get="/contact/1/edit" class="btn btn-primary">
            Click To Edit
            </button>
        </div>
    </main>
</body>
</html>
		`
		w.Write([]byte(content))
	}
}

func handleContactEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := `<form hx-put="/contact/1" hx-target="this" hx-swap="outerHTML">
		<div>
		  <label>First Name</label>
		  <input type="text" name="firstName" value="Joe">
		</div>
		<div class="form-group">
		  <label>Last Name</label>
		  <input type="text" name="lastName" value="Blow">
		</div>
		<div class="form-group">
		  <label>Email Address</label>
		  <input type="email" name="email" value="joe@blow.com">
		</div>
		<button class="btn">Submit</button>
		<button class="btn" hx-get="/contact/1">Cancel</button>
	  </form>`
		w.Write([]byte(content))
	}
}

func handleContactPut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		for k, v := range r.Form {
			fmt.Println(k, v)
		}
		content := `<div hx-target="this" hx-swap="outerHTML">
<div><label>First Name</label>: Joe</div>
<div><label>Last Name</label>: Blow</div>
<div><label>Email</label>: joe@blow.com</div>
<button hx-get="/contact/1/edit" class="btn btn-primary">
Click To Edit
</button>
</div>`
		w.Write([]byte(content))
	}
}
