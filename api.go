package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorAPI defines a struct that is returned on API errors.
// It will use the same structure as the Spotify API to make error handling
// easy for clients.
type ErrorAPI struct {
	Error ErrorObject `json:"error"`
}

// newErrorAPI returns a JSON encoded ErrorAPI object.
func newErrorAPI(status int, message string) string {
	e := &ErrorAPI{}
	e.Error.Status = status
	e.Error.Message = message

	j, _ := json.Marshal(e)
	return string(j)
}

// currentlyPlayingAPI returns the song that the given user id is currently
// playing.
func (a *app) currentlyPlayingAPI(w http.ResponseWriter, r *http.Request, id string) {
	// Get the access and refresh tokens from the database for
	// the given user.
	// If the tokens are empty we'll know that the user hasn't authorized
	// his/her account.
	at, rt := a.getTokens(id)
	if at == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, newErrorAPI(http.StatusNotFound, "not found"))
		return
	}

	// Get the currently playing object for the requested user id.
	cpo, err := a.getCurrentlyPlayingObject(id, at, rt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, newErrorAPI(http.StatusInternalServerError, "internal server error"))
		return
	}

	// This case means that the user isn't currently playing anything.
	if cpo == nil && err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, newErrorAPI(http.StatusOK, "user is not playing anything"))
		return
	}

	j, _ := json.Marshal(cpo)
	fmt.Fprintf(w, string(j))
}

// currentlyPlayingShortAPI returns a formatted text with the currently
// playing song for the given user.
func (a *app) currentlyPlayingShortAPI(w http.ResponseWriter, r *http.Request, id string) {
	// Get the access and refresh tokens from the database for
	// the given user.
	// If the tokens are empty we'll know that the user hasn't authorized
	// his/her account.
	at, rt := a.getTokens(id)
	if at == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, newErrorAPI(http.StatusNotFound, "not found"))
		return
	}

	// Get the currently playing object for the requested user id.
	cpo, err := a.getCurrentlyPlayingObject(id, at, rt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, newErrorAPI(http.StatusInternalServerError, "internal server error"))
		return
	}

	// This case means that the user isn't currently playing anything.
	if (cpo == nil && err == nil) || (!cpo.IsPlaying) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, newErrorAPI(http.StatusOK, "user is not playing anything"))
		return
	}

	// Format the data.
	var artists string
	for _, a := range cpo.Item.Artists {
		if len(artists) == 0 {
			artists = a.Name
		} else {
			artists = artists + ", " + a.Name
		}
	}

	url := cpo.Item.ExternalURLs["spotify"]
	uri := cpo.Item.ID
	track := cpo.Item.Name
	message := fmt.Sprintf("%s - %s @ %s / spotify:track:%s", artists, track, url, uri)

	// Wrap it in a map that we can JSON encode and return it.
	out := map[string]string{"playing": message}
	j, _ := json.Marshal(out)
	fmt.Fprintf(w, string(j))
}
