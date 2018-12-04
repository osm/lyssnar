# Lyssnar

Provides a small web site and API that displays what you're currently
listening to on Spotify.

This is the code that is running on https://lyssnar.com

## Requirements

* PostgreSQL server
* Spotify application

You'll need to create your own Spotify application if you want to play around
with the code, you can do this for free at https://developer.spotify.com/dashboard/applications 

## Development

This is the setup I use when I develop, it assumes that you have a Postgres
server running on your computer and that you have acquired client id and
secrets for your own Spotify app.

```sh
$ cat .env
PORT=8080
DATABASE_URL=postgres://@localhost:5432/lyssnar?sslmode=disable
SPOTIFY_CLIENT_ID=<client id>
SPOTIFY_CLIENT_SECRET=<client secret>
SPOTIFY_CALLBACK=http://localhost:8080/callback
$ export $(cat .env | xargs)
$ make
$ ./lyssnar
$ open http://localhost:8080
```
