package main

// AlbumObjectSimplified contains the simplified album object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#album-object-simplified
type AlbumObjectSimplified struct {
	// The field is present when getting an artist’s albums. Possible
	// values are “album”, “single”, “compilation”, “appears_on”. Compare
	// to album_type this field represents relationship between the artist
	// and the album.
	AlbumGroup string `json:"album_group,omitempty"`

	// The type of the album: one of “album”, “single”, or “compilation”.
	AlbumType string `json:"album_type,omitempty"`

	// The artists of the album. Each artist object includes a link in
	// href to more detailed information about the artist.
	Artists []ArtistObjectSimplified

	// The markets in which the album is available: ISO 3166-1 alpha-2
	// country codes. Note that an album is considered available in a
	// market when at least 1 of its tracks is available in that market.
	AvailableMarkets []string `json:"available_markets"`

	// External URLs for this context.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the album.
	HREF string `json:"href"`

	// The [/documentation/web-api/#spotify-uris-and-ids) for the album.
	ID string `json:"id"`

	// The cover art for the album in various sizes, widest first.
	Images []ImageObject `json:"images"`

	// The name of the album. In case of an album takedown, the value may
	// be an empty string.
	Name string `json:"name"`

	// The date the album was first released, for example 1981. Depending
	// on the precision, it might be shown as 1981-12 or 1981-12-15.
	ReleaseDate string `json:"release_date"`

	// The precision with which release_date value is known: year , month
	// , or day.
	ReleaseDatePrecision string `json:"release_date_precision"`

	// Part of the response when Track Relinking is applied, the original
	// track is not available in the given market, and Spotify did not
	// have any tracks to relink it with. The track response will still
	// contain metadata for the original track, and a restrictions object
	// containing the reason why the track is not available:
	// "restrictions" : {"reason" : "market"}
	Restrictions map[string]string `json:"restrictions"`

	// The object type: “album”.
	Type string `json:"type"`

	// The Spotify URI for the album.
	URI string `json:"uri"`
}

// ArtistObjectSimplified contains the simplified artist object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#artist-object-simplified
type ArtistObjectSimplified struct {
	// External URLs for this context.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the
	// artist.
	HREF string `json:"href"`

	// The Spotify ID for the artist.
	ID string `json:"id"`

	// The name of the artist.
	Name string `json:"name"`

	// The object type: "artist".
	Type string `json:"type"`

	// The Spotify URI for the artist.
	URI string `json:"uri"`
}

// ContextObject contains the context object
// https://developer.spotify.com/documentation/web-api/reference/object-model/#context-object
type ContextObject struct {
	// The object type, e.g. “artist”, “playlist”, “album”.
	Type string `json:"type"`

	// A link to the Web API endpoint providing full details of the track.
	HREF string `json:"href"`

	// External URLs for this context.
	ExternalURLs map[string]string `json:"external_urls"`

	// The Spotify URI for the context.
	URI string `json:"uri"`
}

// CurrentlyPlayingObject contains a combination of previously defined
// objects.
type CurrentlyPlayingObject struct {
	// An Error Object. Can be null.
	Error *ErrorObject `json:"error,omitempty"`

	// A Context Object. Can be null.
	Context *ContextObject `json:"context"`

	// Unix Millisecond Timestamp when data was fetched.
	Timestamp int `json:"timestamp,omitempty"`

	// Progress into the currently playing track. Can be null.
	ProgressMS *int `json:"progress_ms,omitempty"`

	// If something is currently playing.
	IsPlaying bool `json:"is_playing,omitempty"`

	// The currently playing track. Can be null.
	Item *TrackObjectFull `json:"item,omitempty"`

	// The object type of the currently playing item. Can be one of track,
	// episode, ad or unknown.
	CurrentlyPlayingType string `json:"currently_playing_type,omitempty"`
}

// ErrorObject contains the error object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#error-object
type ErrorObject struct {
	// The HTTP status code that is also returned in the response header.
	// For further information, see Response Status Codes.
	Status int

	// A short description of the cause of the error.
	Message string
}

// FollowersObject contains the followers object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#followers-object
type FollowersObject struct {
	// A link to the Web API endpoint providing full details of the
	// followers; null if not available. Please note that this will always
	// be set to null, as the Web API does not support it at the moment.
	HREF *string

	// The total number of followers.
	Total int
}

// ImageObject contains the image object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#image-object
type ImageObject struct {
	// The image height in pixels. If unknown: null or not returned.
	Height int `json:"height"`

	// The source URL of the image.
	URL string `json:"url"`

	// The image width in pixels. If unknown: null or not returned.
	Width int `json:"width"`
}

// LinkedTrackObject contains the linked track object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#track-link
type LinkedTrackObject struct {
	// External URLs for this context.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the track.
	HREF string `json:"href"`

	// The Spotify ID for the track.
	ID string `json:"id"`

	// The object type: “track”.
	Type string `json:"type"`

	// The Spotify URI for the track.
	URI string `json:"uri"`
}

// TrackObjectFull contains the full track object.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#track-object-full
type TrackObjectFull struct {
	// The album on which the track appears. The album object includes a
	// link in href to full information about the album.
	Album AlbumObjectSimplified `json:"album"`

	// The artists who performed the track. Each artist object includes a
	// link in href to more detailed information about the artist.
	Artists []ArtistObjectSimplified `json:"artists"`

	// A list of the countries in which the track can be played,
	// identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`

	// The disc number (usually 1 unless the album consists of more than
	// one disc).
	DiscNumber int `json:"disc_number"`

	// The track length in milliseconds.
	DurationMS int `json:"duration_ms"`

	// Whether or not the track has explicit lyrics ( true = yes it does;
	// false = no it does not OR unknown).
	Explicit bool `json:"explicit"`

	// Known external IDs for the track.
	ExternalIDs map[string]string `json:"external_ids"`

	// External URLs for this context.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the track.
	HREF string `json:"href"`

	// The Spotify ID for the track.
	ID string `json:"id"`

	// Part of the response when Track Relinking is applied. If true , the
	// track is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`

	// Part of the response when Track Relinking is applied, and the
	// requested track has been replaced with different track. The track
	// in the linked_from object contains information about the originally
	// requested track.
	LinkedFrom *LinkedTrackObject `json:"linked_from,omitempty"`

	// Part of the response when Track Relinking is applied, the original
	// track is not available in the given market, and Spotify did not
	// have any tracks to relink it with. The track response will still
	// contain metadata for the original track, and a restrictions object
	// containing the reason why the track is not available:
	// "restrictions" : {"reason" : "market"}
	Restrictions map[string]string `json:"restrictions"`

	// The name of the track.
	Name string `json:"name"`

	// The popularity of the track. The value will be between 0 and 100,
	// with 100 being the most popular.  The popularity of a track is a
	// value between 0 and 100, with 100 being the most popular. The
	// popularity is calculated by algorithm and is based, in the most
	// part, on the total number of plays the track has had and how recent
	// those plays are.  Generally speaking, songs that are being played a
	// lot now will have a higher popularity than songs that were played a
	// lot in the past. Duplicate tracks (e.g. the same track from a
	// single and an album) are rated independently. Artist and album
	// popularity is derived mathematically from track popularity. Note
	// that the popularity value may lag actual popularity by a few days:
	// the value is not updated in real time.
	Popularity int `json:"popularity"`

	// A link to a 30 second preview (MP3 format) of the track. Can be
	// null.
	PreviewURL *string `json:"preview_url"`

	// The number of the track. If an album has several discs, the track
	// number is the number on the specified disc.
	TrackNumber int `json:"track_number"`

	// The object type: “track”.
	Type string `json:"type"`

	// The Spotify URI for the track.
	URI string `json:"uri"`

	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
}

// UserObject contains a partial user object as is fetched without any
// additinal user claims.
// https://developer.spotify.com/documentation/web-api/reference/object-model/#user-object-private
type UserObject struct {
	// An Error Object. Can be null.
	Error *ErrorObject `json:"error,omitempty"`

	// The name displayed on the user’s profile, null if not available.
	DisplayName *string `json:"display_name"`

	// Known external URLs for this user.
	ExternalURLs map[string]string `json:"external_urls"`

	// Information about the followers of the user.
	Followers FollowersObject `json:"followers"`

	// A link to the Web API endpoint for this user.
	HREF string `json:"href"`

	// The Spotify user ID for the user.
	ID string `json:"id"`

	// The user’s profile image.
	Images []ImageObject `json:"images"`

	// The object type: “user”
	Type string `json:"type"`

	// The Spotify URI for the user.
	URI string `json:"uri"`
}
