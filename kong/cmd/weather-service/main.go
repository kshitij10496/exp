package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var flagPort = flag.Int("port", 8080, "port to listen on")

type app struct {
	client *http.Client
}

func New(client *http.Client) *app {
	return &app{
		client: client,
	}
}

func (a *app) Run(ctx context.Context) error {
	addr := fmt.Sprintf(":%d", *flagPort)
	fmt.Println("Listening on ", addr)

	if err := http.ListenAndServe(addr, a.router(ctx)); err != nil {
		return fmt.Errorf("failed to listen and serve HTTP connections: %w", err)
	}

	return nil
}

func (a *app) fetchWeather(ctx context.Context, location string) (string, error) {
	uri := fmt.Sprintf("https://wttr.in/%s", location)

	u, err := url.Parse(uri)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}
	qparams := u.Query()
	qparams.Add("format", "%t\n")
	u.RawQuery = qparams.Encode()

	uf := u.String()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uf, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	if err := resp.Body.Close(); err != nil {
		return "", fmt.Errorf("failed to close response body: %w", err)
	}

	return string(b), nil
}

func (a *app) router(ctx context.Context) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		location := r.URL.Query().Get("location")
		if location == "" {
			location = "berlin"
		}
		result, err := a.fetchWeather(ctx, location)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if result == "" {
			http.Error(w, "failed to fetch weather", http.StatusInternalServerError)
			return
		}

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{"result": result}); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	})
	return mux
}

func main() {
	a := New(http.DefaultClient)
	if err := a.Run(context.Background()); err != nil {
		panic(err)
	}
}
