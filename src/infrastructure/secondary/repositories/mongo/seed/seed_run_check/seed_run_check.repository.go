package seed_run_check_repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/pkg"
)

type SeedRunCheckRepository struct {
	db pkg.MongoConnectionPort
}

func New(db pkg.MongoConnectionPort) *SeedRunCheckRepository {
	return &SeedRunCheckRepository{
		db,
	}
}

func (r *SeedRunCheckRepository) Execute() error {
	collection := r.db.Connection().Collection("clients-portfolios")

	var result bson.M
	err := collection.FindOne(context.Background(), bson.M{}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil
	}

	if err != nil {
		return err
	}

	return seed_errors.ErrSeedAlreadyExecuted

}
