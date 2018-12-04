package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// Collection of templates and data.
var (
	dCss              string
	dFavicon          string
	dFavicon16        string
	dFavicon32        string
	tAuthorized       = template.Must(template.ParseFiles("./ui/authorized.html"))
	tCurrentlyPlaying = template.Must(template.ParseFiles("./ui/currently-playing.html"))
	tError            = template.Must(template.ParseFiles("./ui/error.html"))
	tLanding          = template.Must(template.ParseFiles("./ui/landing.html"))
)

// loadStaticFile reads the contents of the given file, if it can't find the
// find it'll log a fatal error.
func loadStaticFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("can't find %s", file)
	}
	return string(data)
}

func init() {
	// Fetch the contents of the style sheets on init so that we don't
	// have to do it on each request.
	dCss = loadStaticFile("./ui/lyssnar.css")

	// Load the favicons.
	dFavicon = loadStaticFile("./ui/favicon.ico")
	dFavicon16 = loadStaticFile("./ui/favicon-16x16.png")
	dFavicon32 = loadStaticFile("./ui/favicon-32x32.png")
}

// errorNotFound displays the 404 page.
func (a *app) errorNotFound(w http.ResponseWriter, r *http.Request) {
	tError.Execute(w, map[string]string{"header": ":-(", "message": "The requested page doesn't exist."})
}

// landing renders the landing page of the site.
func (a *app) landing(w http.ResponseWriter, r *http.Request) {
	tLanding.Execute(w, nil)
}

// callback handles the response from the authorization page at Spotify.
func (a *app) callback(w http.ResponseWriter, r *http.Request) {
	// Make sure we didn't get an error back from Spotify.
	if r.FormValue("error") != "" {
		tError.Execute(w, map[string]string{"header": ":-(", "message": "An error occured, try again later."})
		return
	}

	// Exchange the code for a token.
	t, err := a.conf.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		tError.Execute(w, map[string]string{"header": ":-(", "message": "An error occured, try again later."})
		return
	}

	// Get the user fro
	u, err := a.getUserObject(t.AccessToken, t.RefreshToken)
	if err != nil {
		tError.Execute(w, map[string]string{"header": ":-(", "message": "An error occured, try again later."})
		return
	}

	// Store the access and refresh tokens in our database.
	a.storeTokens(u.ID, t.AccessToken, t.RefreshToken)

	// Render the output.
	tAuthorized.Execute(w, map[string]string{"id": u.ID})
}

// currentlyPlaying displays what the requested user currently is playing.
func (a *app) currentlyPlaying(w http.ResponseWriter, r *http.Request, id string) {
	// Get the access and refresh tokens from the database for
	// the given user.
	// If the tokens are empty we'll know that the user hasn't authorized
	// his/her account.
	at, rt := a.getTokens(id)
	if at == "" {
		tError.Execute(w, map[string]string{"header": ":-(", "message": "The account is not authorized on lyssnar.com yet"})
		return
	}

	// Get the currently playing object for the requested user id.
	cpo, err := a.getCurrentlyPlayingObject(id, at, rt)
	if err != nil {
		tError.Execute(w, map[string]string{"header": ":-(", "message": "An error occured, try again later."})
		return
	}

	// This case means that the user isn't currently playing anything.
	if cpo == nil && err == nil {
		tError.Execute(w, map[string]string{"header": "Not active", "message": fmt.Sprintf("%s is not using Spotify right now", id)})
		return
	}

	// Extract the correct image URL.
	imageURL := ""
	for _, i := range cpo.Item.Album.Images {
		if i.Height > 200 && i.Height < 500 {
			imageURL = i.URL
		}
	}

	// Concatenate all artists into one string.
	artists := ""
	for _, a := range cpo.Item.Artists {
		if artists == "" {
			artists = a.Name
		} else {
			artists = fmt.Sprintf("%s, %s", artists, a.Name)
		}
	}

	// Render the user template.
	tCurrentlyPlaying.Execute(w, map[string]string{
		"id":     id,
		"artist": artists,
		"track":  cpo.Item.Name,
		"url":    cpo.Item.ExternalURLs["spotify"],
		"image":  imageURL,
	})
}
