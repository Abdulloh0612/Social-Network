package storage

import (
	"comment-service/storage/postgres"
	"comment-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.CommentStorageI
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.postRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewCommentRepo(db)}
}
