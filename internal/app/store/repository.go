package store

import "github.com/MegOBonus/go-rest-api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}