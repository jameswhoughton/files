package files

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Chunk struct {
	Id       uuid.UUID
	Hash     string
	Order    int
	Location []uuid.UUID
}

type FileMeta struct {
	Id         uuid.UUID
	Name       string
	Hash       string
	UploadedAt time.Time
	ExpiresAt  time.Time
	Chunks     []Chunk
}

var ErrFileMetaNotFound = errors.New("Meta information for file not found")

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (FileMeta, error)
	Store(ctx context.Context, fileMeta FileMeta) error
}
