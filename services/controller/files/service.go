package files

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

type Service struct {
	repo Repository
}

func (s *Service) GetFileMeta(ctx context.Context, id uuid.UUID) (FileMeta, error) {
	fm, err := s.repo.Get(ctx, id)

	if err != nil {
		if err == ErrFileMetaNotFound || time.Now().After(fm.ExpiresAt) {
			return FileMeta{}, fmt.Errorf("%w id=%s", err, id.String())
		}

		return FileMeta{}, fmt.Errorf("Failed to fetch meta information for file id=%s: %w", id.String(), err)
	}

	return fm, nil
}

func (s *Service) Store(ctx context.Context, fileMeta FileMeta) error {
	return nil
}
