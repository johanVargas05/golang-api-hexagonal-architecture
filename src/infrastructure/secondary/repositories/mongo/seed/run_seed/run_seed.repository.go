package run_seed_repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/pkg"
	models_mongo "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/mongo/models"
)

type RunSeedRepository struct {
	db pkg.MongoConnectionPort
}

func New(db pkg.MongoConnectionPort) *RunSeedRepository {
	return &RunSeedRepository{
		db,
	}
}

func (r *RunSeedRepository) Execute(portfolios[]*entities.Portfolio) error{

	ctx := context.Background()

	collection := r.db.Connection().Collection("clients-portfolios")

	var documents []interface{}

	for _, portfolio := range portfolios {
		createdDate:= portfolio.CreateAt()
		taxes:= []models_mongo.Tax{}

		for _, tax := range portfolio.Taxes() {
			taxes = append(taxes, models_mongo.Tax{
				TaxType: tax.TypeTax(),
				TaxId: tax.ID(),
				Rate: tax.Rate(),
			})
		}
		
		id, err := primitive.ObjectIDFromHex(portfolio.ID())
		if err !=nil {
			return seed_errors.ErrSeedNotExecuted
		}
		documents = append(documents, models_mongo.Portfolio{
			ID: id,
			Channel: portfolio.Channel(),
			Country: portfolio.Country(),
			CreatedDate: &createdDate,
			CustomerCode: portfolio.CustomerID(),
			Route: portfolio.Route(),
			Sku: portfolio.SKU(),
			Title: portfolio.Title(),
			CategoryId: portfolio.CategoryID(),
			Category: portfolio.Category(),
			Brand: portfolio.Brand(),
			Classification: portfolio.Classification(),
			UnitsPerBox: portfolio.UnitsPerBox(),
			MinOrderUnits: portfolio.MinOrderUnits(),
			PackageDescription: portfolio.PackageDescription(),
			PackageUnitDescription: portfolio.PackageUnitDescription(),
			QuantityMaxRedeem: portfolio.QuantityMaxRedeem(),
			RedeemUnit: portfolio.RedeemUnit(),
			OrderReasonRedeem: portfolio.OrderReasonRedeem(),
			SkuRedeem: portfolio.SKURedeem(),
			Price: models_mongo.Price{
				FullPrice: portfolio.FullPrice(),
				Taxes: taxes,
			},
			Points: portfolio.Points(),
		})
		}

	result,err:=collection.InsertMany(ctx, documents)
	if err!=nil{
		return seed_errors.ErrSeedNotExecuted
	}

	if result.InsertedIDs == nil {
		return seed_errors.ErrSeedNotExecuted
	}

	searchIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "branch", Value: "text"},
			{Key: "category", Value: "text"},
			{Key: "sku", Value: "text"},
			{Key: "classification", Value: "text"},
		},
		Options: options.Index().SetName("portfolio_search_index"),
	}
	_,err=collection.Indexes().CreateOne(ctx, searchIndex)
	
	if err!=nil{
		return errors.New("error creating index search")
	}

	_,err=collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key:"customerCode", Value:1},
			{Key:"title", Value: 1},
			{Key:"sku", Value: 1},
			{Key:"branch", Value: 1},
			{Key:"createdDate", Value: 1},
			{Key:"points", Value: 1},
			{Key:"price.fullPrice",Value:1},
			{Key:"minOrderUnits", Value: 1},
		},
		Options: options.Index().SetName("portfolio_order_index"),
	})

	if err!=nil{
		return errors.New("error creating index order")
	}


	return nil
}


