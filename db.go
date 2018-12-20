package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/osm/migrator"
	"github.com/osm/migrator/repository"
)

// initDB initializes the database and runs any migrations that might
// not have been executed before.
func (a *app) initDB() error {
	var err error
	if a.db, err = sql.Open("postgres", a.dbURL); err != nil {
		return fmt.Errorf("can't initialize database connection: %v", err)
	}

	return migrator.ToLatest(a.db, getDatabaseRepository())
}

// getDatabaseRepository returns a new repository with all the migrations
// that exists for the database.
func getDatabaseRepository() repository.Source {
	return repository.FromMemory(map[int]string{
		1: "CREATE TABLE migration (version TEXT NOT NULL PRIMARY KEY);",
		2: "CREATE TABLE credential (id text NOT NULL PRIMARY KEY, access_token text NOT NULL, refresh_token text NOT NULL, created_at timestamp with time zone NOT NULL, updated_at timestamp with time zone);",
	})
}

// getTokens fetches the access and refresh tokens for the given user id.
func (a *app) getTokens(id string) (string, string) {
	var at, rt string
	a.db.QueryRow("SELECT access_token, refresh_token FROM credential WHERE id = $1", id).Scan(&at, &rt)
	return at, rt
}

// updateAccessToken updates the access token for the given user id.
func (a *app) updateAccessToken(id, at string) error {
	stmt, err := a.db.Prepare("UPDATE credential SET access_token = $1, updated_at = now() WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(at, id)
	if err != nil {
		return err
	}

	return nil
}

// storeTokens stores the access and refresh tokens for the given user id.
// It will remove the existing entry, if any, before it stores the new
// values.
func (a *app) storeTokens(id, at, rt string) error {
	a.db.Exec("DELETE FROM credential WHERE id = $1", id)

	stmt, err := a.db.Prepare("INSERT INTO credential VALUES($1, $2, $3, now(), null)")
	if err != nil {
		fmt.Println("prep" + err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, at, rt)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// deleteUser removes the credential for the given user id.
func (a *app) deleteUser(id string) {
	a.db.Exec("DELETE FROM credential WHERE id = $1", id)
}
