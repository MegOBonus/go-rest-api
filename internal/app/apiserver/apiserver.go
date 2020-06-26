package apiserver

import (
	"database/sql"
	"github.com/MegOBonus/go-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)

	if err != nil {
		return err
	}

	defer db.Close()

	s := sqlstore.New(db)
	srv := newServer(s)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
