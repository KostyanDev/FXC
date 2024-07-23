package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"
)

type storage struct {
	logger *logrus.Logger
	ext    *sqlx.DB
}

func New(logger *logrus.Logger, ext *sqlx.DB) *storage {
	return &storage{
		logger: logger,
		ext:    ext,
	}
}
