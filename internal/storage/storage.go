package storage

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Storage struct {
	logger *logrus.Logger
	ext    *sql.DB
}

func New(logger *logrus.Logger, ext *sql.DB) *Storage {
	return &Storage{
		logger: logger,
		ext:    ext,
	}
}
