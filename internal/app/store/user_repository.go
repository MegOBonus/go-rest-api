package store

import "github.com/MegOBonus/go-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users where email = $1",
		email,
	).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		return nil, err
	}

	return user, nil
}
