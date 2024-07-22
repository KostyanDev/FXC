package service

import (
	"context"

	"app/internal/storage"

	"github.com/sirupsen/logrus"
)

type Service struct {
	ctx     context.Context
	log     *logrus.Logger
	storage storage.Storage
}

func New(ctx context.Context, log *logrus.Logger, storage storage.Storage) *Service {
	return &Service{
		ctx:     ctx,
		log:     log,
		storage: storage,
	}
}
