package model

import "testing"

func TesUser(t *testing.T) *User {
	return &User{
		Email:    "test@test.com",
		Password: "password11",
	}
}
