package all_port_folio_of_user_repository

import (
	"context"
	"fmt"

	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	portfolio_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/pkg"
	models_mongo "github.com/johanVargas05/golang-api-hexagonal-architecture/src/infrastructure/secondary/repositories/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
)

var fieldsMongo = map[string]string{
	"min_order_units": "minOrderUnits",
	"create_at":      "createdDate",
	"title":           "title",
	"brand":           "brand",
	"price":           "price.fullPrice",
	"points":          "points",
}

var typeSort = map[string]int{
	"asc":  1,
	"desc": -1,
}

type AllPortfolioOfUserRepository struct {
	db pkg.MongoConnectionPort
}

func New(db pkg.MongoConnectionPort) *AllPortfolioOfUserRepository {
	return &AllPortfolioOfUserRepository{
		db,
	}
}

func (repository *AllPortfolioOfUserRepository) Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio,int, error) {
	collection:=repository.db.Connection().Collection("clients-portfolios")
	
	aggregation,totalItemsQuery:=buildQueries(params)

	totalItems, err := collection.CountDocuments(context.Background(), totalItemsQuery)
	
	if err != nil {
		return nil, 0, err
	}

	cursor, err := collection.Aggregate(context.Background(), aggregation)

	if err != nil {
		return nil, 0, err

	}

	var portfolios []*models_mongo.Portfolio
	
	if err = cursor.All(context.Background(), &portfolios); err != nil {
		return nil, 0, err
	}

	if len(portfolios) == 0 {
		return nil, 0, portfolio_errors.ErrNotFoundPortfoliosOfUser
	}

	result:=mapperModelToEntity(portfolios)

	return result,int(totalItems), nil

}

func buildQueries(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]bson.D,bson.D) {
	aggregation:=[]bson.D{}	
	totalItemsQuery:=bson.D{}
	if params.Search != "" {
		aggregation = append(aggregation, bson.D{
			{
				Key: "$match", Value: bson.D{
					{Key: "customerCode",Value: params.UserId},
					{Key: "$text", Value: 
						bson.D{
							{Key: "$search", Value: params.Search},
						},
					}	,
				},
			},
		})
		totalItemsQuery = bson.D{
			{
				Key: "customerCode", Value: params.UserId,
			},
			{
				Key: "$text", Value: bson.D{
					{Key: "$search", Value: params.Search},
				},
			},
		}
	} else {
		aggregation = append(aggregation, bson.D{{Key: "$match", Value: bson.D{{Key: "customerCode", Value: params.UserId}}}})
		totalItemsQuery = bson.D{{Key: "customerCode", Value: params.UserId}}
	}
	aggregation = append(aggregation, bson.D{{Key: "$sort", Value: bson.D{{Key: fieldsMongo[params.SortBy], Value: typeSort[params.SortType]}}}})

	aggregation = append(aggregation, bson.D{{Key: "$skip", Value: (params.Page -1) * params.Limit}})
	aggregation = append(aggregation, bson.D{{Key: "$limit", Value: params.Limit}})

	return aggregation,totalItemsQuery
}

func mapperModelToEntity(portfolios []*models_mongo.Portfolio) []*entities.Portfolio {
	var portfoliosEntities []*entities.Portfolio = make([]*entities.Portfolio, len(portfolios))

	for i, portfolio := range portfolios {

		taxes := make([]*entities.Tax, len(portfolio.Price.Taxes))

		for j, tax := range portfolio.Price.Taxes {
			taxes[j] = entities.InitTax(tax.TaxId, tax.TaxType, int(tax.Rate))
		}

		portfoliosEntities[i] = entities.InitPortfolio(&entities.PortfolioEntityParams{
			ID:                 portfolio.ID.Hex(),
			Channel:            portfolio.Channel,
			Country:            portfolio.Country,
			CreateAt:        portfolio.CreatedDate,
			CustomerID: 			 portfolio.CustomerCode,
			Route:              portfolio.Route,
			SKU: 							portfolio.Sku,
			Title: 						portfolio.Title,
			CategoryID: 				portfolio.CategoryId,
			Category: 					portfolio.Category,
			Brand: 						portfolio.Brand,
			Classification: 	portfolio.Classification,
			UnitsPerBox: 			fmt.Sprintf("%d", portfolio.UnitsPerBox),
			MinOrderUnits: 		fmt.Sprintf("%.2f", portfolio.MinOrderUnits),
			PackageDescription: portfolio.PackageDescription,
			PackageUnitDescription: portfolio.PackageUnitDescription,
			QuantityMaxRedeem: portfolio.QuantityMaxRedeem,
			RedeemUnit: portfolio.RedeemUnit,
			OrderReasonRedeem: portfolio.OrderReasonRedeem,
			SKURedeem: portfolio.SkuRedeem,
			Price: portfolio.Price.FullPrice,
			Points: portfolio.Points,
			Taxes: taxes,
		})
			
	}
	
	return portfoliosEntities
}
	