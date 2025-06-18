package files_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jameswhoughton/files/services/controller/files"
	"github.com/jameswhoughton/files/services/controller/mongoStore"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestReturnsExpectedErrorIfIdDoesNotExist(t *testing.T) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalln(err)
	}

	service := files.NewService(mongoStore.NewFileRepository(client))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	fm, err := service.GetFileMeta(ctx, uuid.New())

	if !errors.Is(err, files.ErrFileMetaNotFound) {
		t.Errorf("expecting error `%s`, got `%s`", files.ErrFileMetaNotFound, err)
	}

	if fm.Id != uuid.Nil {
		t.Errorf("Expected blank FileMeta, got %+v", fm)
	}
}
