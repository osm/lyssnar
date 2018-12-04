package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

// app contains the internal data structure used by the application.
type app struct {
	db    *sql.DB
	dbURL string
	conf  *oauth2.Config
	port  string
}

// getEnv looks for the given key in the environment and logs a fatal
// error if the value can't be found or is empty.
func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("$%s must be set", key)
	}
	return val
}

// main is the entry point of the application.
func main() {
	port := getEnv("PORT")
	spotifyCallback := getEnv("SPOTIFY_CALLBACK")
	spotifyClientID := getEnv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := getEnv("SPOTIFY_CLIENT_SECRET")
	dbURL := getEnv("DATABASE_URL")

	a := &app{
		conf: &oauth2.Config{
			RedirectURL:  spotifyCallback,
			ClientID:     spotifyClientID,
			ClientSecret: spotifyClientSecret,
			Scopes: []string{
				"user-read-currently-playing",
			},
			Endpoint: spotify.Endpoint,
		},
		dbURL: dbURL,
		port:  port,
	}

	if err := a.initDB(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", a.route)
	http.ListenAndServe(":"+a.port, nil)
}
