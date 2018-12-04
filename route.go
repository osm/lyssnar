package main

import (
	"fmt"
	"net/http"
	"regexp"
)

// Collection of regular expressions that the incoming requests are matched
// against.
var (
	rAuthorize           = regexp.MustCompile(`^/authorize$`)
	rCallback            = regexp.MustCompile(`^/callback$`)
	rCss                 = regexp.MustCompile(`^/lyssnar.css$`)
	rFavicon             = regexp.MustCompile(`^/favicon.ico$`)
	rFavicon16           = regexp.MustCompile(`^/favicon-16x16.png$`)
	rFavicon32           = regexp.MustCompile(`^/favicon-32x32.png$`)
	rLanding             = regexp.MustCompile(`^/$`)
	rCurrentlyPlaying    = regexp.MustCompile(`^/~([a-zA-Z-]+)$`)
	rCurrentlyPlayingAPI = regexp.MustCompile(`^/v1/user/([a-zA-Z-]+)/currently-playing$`)
)

// route handles all http requests and routes them to the appropriate
// handler.
func (a *app) route(w http.ResponseWriter, r *http.Request) {
	if m := rCss.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		fmt.Fprintf(w, dCss)
	} else if m := rFavicon.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "image/x-icon")
		fmt.Fprintf(w, dFavicon)
	} else if m := rFavicon16.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "image/png")
		fmt.Fprintf(w, dFavicon16)
	} else if m := rFavicon32.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "image/png")
		fmt.Fprintf(w, dFavicon32)
	} else if m := rLanding.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		a.landing(w, r)
	} else if m := rCurrentlyPlaying.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		a.currentlyPlaying(w, r, m[1])
	} else if m := rCurrentlyPlayingAPI.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		a.currentlyPlayingAPI(w, r, m[1])
	} else if m := rAuthorize.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		http.Redirect(w, r, a.conf.AuthCodeURL(newUUID()), http.StatusTemporaryRedirect)
	} else if m := rCallback.FindStringSubmatch(r.URL.Path); len(m) > 0 {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		a.callback(w, r)
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		a.errorNotFound(w, r)
	}
}
