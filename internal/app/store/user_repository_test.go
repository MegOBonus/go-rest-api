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

	u, err := s.User().Create(&model.User{Email: "test@test.test"})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	println("dburl", databaseURL)
	defer teardown("users")

	email := "test@test.test"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	userAfterInsert, err := s.User().Create(&model.User{Email: email})
	assert.NoError(t, err)
	assert.NotNil(t, userAfterInsert)

	userAfterSelect, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, userAfterSelect)
	assert.Equal(t, userAfterSelect, userAfterInsert)
}
