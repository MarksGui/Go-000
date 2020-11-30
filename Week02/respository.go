package main

import (
	"database/sql"

	"github.com/pkg/errors"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (p *UserRepository) GetInfo() (*UserEntity, error) {
	// exec query SQL
	err := sql.ErrNoRows
	if err != nil {
		return nil, errors.Wrap(err, "query user failed")
	}
	return &UserEntity{}, nil
}
