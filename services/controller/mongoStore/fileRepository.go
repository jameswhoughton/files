package mongoStore

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jameswhoughton/files/services/controller/files"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewFileRepository(client *mongo.Client) files.Repository {
	return &FileRepository{
		client: client,
	}
}

type FileRepository struct {
	client *mongo.Client
}

func (r FileRepository) Get(ctx context.Context, id uuid.UUID) (files.FileMeta, error) {
	var result files.FileMeta

	collection := r.client.Database("files").Collection("file_meta")

	err := collection.FindOne(ctx, bson.D{{"_id", id.String()}}).Decode(&result)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return files.FileMeta{}, files.ErrFileMetaNotFound
	}

	return result, nil
}

func (r *FileRepository) Store(ctx context.Context, fileMeta files.FileMeta) error {
	_, err := r.client.Database("files").Collection("file_meta").InsertOne(ctx, fileMeta)

	if err != nil {
		return fmt.Errorf("failed to store meta information for file: %w", err)
	}

	return nil
}
