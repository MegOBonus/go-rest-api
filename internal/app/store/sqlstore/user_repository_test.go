package sqlstore_test

import (
	"github.com/MegOBonus/go-rest-api/internal/app/model"
	"github.com/MegOBonus/go-rest-api/internal/app/store"
	"github.com/MegOBonus/go-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	u := model.TesUser(t)
	s := sqlstore.New(db)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := model.TesUser(t).Email
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TesUser(t)
	s.User().Create(u)

	userAfterSelect, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, userAfterSelect)
}
