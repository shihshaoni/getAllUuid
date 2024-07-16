package service

import (
	"context"
	"getAllUuid/model"
	"getAllUuid/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func FetchAllUUIDs(client *mongo.Client, ctx context.Context) ([]model.UUID, error) {
	return repository.GetAllUUIDs(client, ctx)
}
