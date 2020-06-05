package store_test

import (
	"github.com/MegOBonus/go-rest-api/internal/app/model"
	"github.com/MegOBonus/go-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TesUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := model.TesUser(t).Email
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TesUser(t)
	s.User().Create(u)

	userAfterSelect, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, userAfterSelect)
}
