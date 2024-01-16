package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var flagPort = flag.Int("port", 8080, "port to listen on")

type app struct{}

func (a *app) Run(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", *flagPort)
	fmt.Println("Listening on ", addr)

	if err := http.ListenAndServe(addr, a.router(ctx)); err != nil {
		return fmt.Errorf("failed to listen and serve HTTP connections: %w", err)
	}

	return nil
}

func (a *app) router(ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		if a == "" || b == "" {
			http.Error(w, "missing parameters", http.StatusBadRequest)
			return
		}

		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		j, err := strconv.ParseFloat(b, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		result := i + j

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]float64{"result": result}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		if a == "" || b == "" {
			http.Error(w, "missing parameters", http.StatusBadRequest)
			return
		}

		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		j, err := strconv.ParseFloat(b, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		result := i - j

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]float64{"result": result}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/mul", func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		if a == "" || b == "" {
			http.Error(w, "missing parameters", http.StatusBadRequest)
			return
		}

		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		j, err := strconv.ParseFloat(b, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		result := i * j

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]float64{"result": result}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/div", func(w http.ResponseWriter, r *http.Request) {
		a := r.URL.Query().Get("a")
		b := r.URL.Query().Get("b")

		if a == "" || b == "" {
			http.Error(w, "missing parameters", http.StatusBadRequest)
			return
		}

		i, err := strconv.ParseFloat(a, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		j, err := strconv.ParseFloat(b, 64)
		if err != nil {
			http.Error(w, "invalid parameter", http.StatusBadRequest)
			return
		}

		result := i / j

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]float64{"result": result}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	return mux
}

func main() {
	a := &app{}
	if err := a.Run(context.Background()); err != nil {
		panic(err)
	}
}
