package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// getUserObject fetches the user profile from the Spotify API.
func (a *app) getUserObject(at, rt string) (*UserObject, error) {
	// Allow the code to be retried once, we do this because the access token
	// might have expired. When this is the case we'll use the refresh token
	// to acquire a new access token and try again.
	retry := true

	// Construct the oauth2 token.
	t := &oauth2.Token{AccessToken: at, RefreshToken: rt}
start:
	cli := a.conf.Client(oauth2.NoContext, t)
	res, err := cli.Get("https://api.spotify.com/v1/me")
	if err != nil {
		log.Printf("failed to get https://api.spotify.com/v1/me, error: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed reading res body: %s", err.Error())
		return nil, err
	}

	uo := &UserObject{}
	if err := json.Unmarshal(d, uo); err != nil {
		log.Printf("failed to unmarshal json in getID: %s, data: %s", err.Error(), string(d))
		return nil, err
	}

	// The token has probably expired, invalidate the access token
	// and try again.
	if uo.Error != nil && uo.Error.Status == 401 && retry {
		retry = false
		t.AccessToken = ""
		goto start
	}

	return uo, nil
}

// getCurrentlyPlaying fetches the currently playing song for the user that
// the token belongs to from the Spotify API.
func (a *app) getCurrentlyPlayingObject(id, at, rt string) (*CurrentlyPlayingObject, error) {
	// Allow the code to be retried once, we do this because the access token
	// might have expired. When this is the case we'll use the refresh token
	// to acquire a new access token and try again.
	retry := true

	// Construct the oauth2 token.
	t := &oauth2.Token{AccessToken: at, RefreshToken: rt}
start:
	cli := a.conf.Client(oauth2.NoContext, t)
	res, err := cli.Get("https://api.spotify.com/v1/me/player/currently-playing")
	if err != nil {
		log.Printf("failed to get https://api.spotify.com/v1/me/player/currently-playing, error: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	// The API returns a No Content status code if the user isn't playing
	// anything. If that's the case we'll just return nil.
	if res.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed reading res body: %s", err.Error())
		return nil, err
	}

	cpo := &CurrentlyPlayingObject{}
	if err := json.Unmarshal(d, cpo); err != nil {
		log.Printf("failed to unmarshal json in getCurrentlyPlaying: %s, %s", err.Error(), string(d))
		return nil, err
	}

	// The token has probably expired, invalidate the access token
	// and try again.
	if cpo.Error != nil && cpo.Error.Status == 401 && retry {
		retry = false
		t.AccessToken = ""
		goto start
	}

	// Check whether the token used in the request is the same as the
	// token passed to the method. If it changed we'll update the
	// database with the used (new) token.
	used, _ := cli.Transport.(*oauth2.Transport).Source.Token()
	if used.AccessToken != at {
		a.updateAccessToken(id, used.AccessToken)
	}

	return cpo, nil
}
