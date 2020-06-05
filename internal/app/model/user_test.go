package model_test

import (
	"github.com/MegOBonus/go-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *model.User {
				return model.TesUser(t)
			},
			isValid: true,
		},
		{
			name: "With encrypted password",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Password = ""
				u.EncryptedPassword = "eohfweoiifkpwqrjiow"
				return u
			},
			isValid: true,
		},
		{
			name: "Empty email",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid email",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "Empty password",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Too small password",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Password = "1234"
				return u
			},
			isValid: false,
		},
		{
			name: "Too big password",
			u: func() *model.User {
				u := model.TesUser(t)
				u.Password = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TesUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
