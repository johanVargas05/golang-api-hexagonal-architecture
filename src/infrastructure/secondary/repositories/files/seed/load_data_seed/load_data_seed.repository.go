package load_data_seed_repository

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
)

type Tax struct {
	TaxType string `json:"taxType"`
	TaxId   string `json:"taxId"`
	Rate    int `json:"rate"`
}

type Price struct {
	FullPrice int   `json:"fullPrice"`
	Taxes     []Tax `json:"taxes"`
}

type Portfolio struct {
	ID                  map[string]string `json:"_id"`
	Channel             string            `json:"channel"`
	Country             string            `json:"country"`
	CreatedDate         map[string]string `json:"createdDate"`
	CustomerCode        string            `json:"customerCode"`
	Route               string            `json:"route"`
	Sku                 string            `json:"sku"`
	Title               string            `json:"title"`
	CategoryId          string            `json:"categoryId"`
	Category            string            `json:"category"`
	Brand               string            `json:"brand"`
	Classification      string            `json:"classification"`
	UnitsPerBox         string            `json:"unitsPerBox"`
	MinOrderUnits       string            `json:"minOrderUnits"`
	PackageDescription  string            `json:"packageDescription"`
	PackageUnitDescription string        `json:"packageUnitDescription"`
	QuantityMaxRedeem   int               `json:"quantityMaxRedeem"`
	RedeemUnit          string            `json:"redeemUnit"`
	OrderReasonRedeem   string            `json:"orderReasonRedeem"`
	SkuRedeem           bool              `json:"skuRedeem"`
	Price               Price             `json:"price"`
	Points              int               `json:"points"`
}

type LoadDataSeedRepository struct{}

func New() *LoadDataSeedRepository {
	return &LoadDataSeedRepository{}
}

func (l *LoadDataSeedRepository) Execute() ([]*entities.Portfolio, error) {
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
			filePath = "src/infrastructure/secondary/repositories/files/seed/data/portfolios.clients-portfolios.json"
	}
	fileJson,err:=os.Open(filePath)
	
	if err!=nil{
		return nil,err
	}

	defer fileJson.Close()

	content,err:=io.ReadAll(fileJson)

	if err!=nil{
		return nil,err
	}

	var portfolios []*Portfolio

	err = json.Unmarshal(content, &portfolios)

	if err!=nil{
		return nil,err
	}

	var portfoliosEntities []*entities.Portfolio

	for _,portfolio:=range portfolios{
		createAt,err:=time.Parse("2006-01-02T15:04:05.000Z07:00",portfolio.CreatedDate["$date"])
		
		if err!=nil{
			return nil,err
		}

		orderReasonRedeem,err:=strconv.Atoi(portfolio.OrderReasonRedeem)

		if err!=nil{
			return nil,err
		}

		var taxes []*entities.Tax
		for _,tax:=range portfolio.Price.Taxes{
			taxes=append(taxes,entities.InitTax(tax.TaxId,tax.TaxType,tax.Rate))
		}
		
		portfolioEntity:=entities.InitPortfolio(&entities.PortfolioEntityParams{
			ID:                  portfolio.ID["$oid"],
			Channel:             portfolio.Channel,
			Country:             portfolio.Country,
			CreateAt: 					&createAt,
			CustomerID:        portfolio.CustomerCode,
			Route:               portfolio.Route,
			SKU:                 portfolio.Sku,
			Title:               portfolio.Title,
			CategoryID:          portfolio.CategoryId,
			Category:            portfolio.Category,
			Brand:               portfolio.Brand,
			Classification:      portfolio.Classification,
			UnitsPerBox:         portfolio.UnitsPerBox,
			MinOrderUnits:       portfolio.MinOrderUnits,
			PackageDescription:  portfolio.PackageDescription,
			PackageUnitDescription: portfolio.PackageUnitDescription,
			QuantityMaxRedeem:   portfolio.QuantityMaxRedeem,
			RedeemUnit:          portfolio.RedeemUnit,
			OrderReasonRedeem:   orderReasonRedeem,
			SKURedeem:           portfolio.SkuRedeem,
			Price:               float64(portfolio.Price.FullPrice),
			Points:              portfolio.Points,
			Taxes: 						 taxes,
		})

		portfoliosEntities=append(portfoliosEntities,portfolioEntity)
	}

	return portfoliosEntities,nil
}
