package store

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestStore(t *testing.T, dbUrl string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseURL = dbUrl
	s := New(config)
	if err := s.Open(); err != nil {
		log.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ","))); err != nil {
				log.Fatal(err)
			}
		}

		s.Close()
	}
}
