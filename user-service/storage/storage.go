package storage

import (
	"user-service/storage/postgres"
	"user-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewUserRepo(db)}
}
