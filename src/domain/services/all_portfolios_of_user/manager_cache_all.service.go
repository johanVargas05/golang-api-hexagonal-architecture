package all_portfolios_of_user_service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"
	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	manager_cache_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/manager_cache"
)
type taxCacheParams struct {
	ID string
	TypeTax string
	Rate int
}

type portfolioCacheParams struct {
	ID                  string
	Channel             string
	Country             string
	CreateAt            *time.Time
	CustomerID          string
	Route               string
	SKU                 string
	Title               string
	CategoryID          string
	Category            string
	Brand               string
	Classification      string
	UnitsPerBox         string
	MinOrderUnits       string
	PackageDescription  string
	PackageUnitDescription string
	QuantityMaxRedeem   int
	RedeemUnit          string
	OrderReasonRedeem   int
	SKURedeem           bool
	Price               float64
	Points              int
	Taxes               []*taxCacheParams
}
type dataCacheStruct struct {
	Data []*portfolioCacheParams 
	Pagination *pagination_dtos.PaginationResponseDto 
}

type managerCacheAllPortfolioOfUserService struct {
	managerCacheService manager_cache_ports.ManagerCacheServicePort
}

func NewManagerCache(managerCacheService manager_cache_ports.ManagerCacheServicePort) *managerCacheAllPortfolioOfUserService {
	return &managerCacheAllPortfolioOfUserService{
		managerCacheService,
	}
}

func (service *managerCacheAllPortfolioOfUserService) Get(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, *pagination_dtos.PaginationResponseDto) {

	key := generatePaths(params)

	dataCache, err := service.managerCacheService.GetData( key, &dataCacheStruct{})

	if err != nil {
		return nil, nil
	}

	data := dataCache.(*dataCacheStruct)

	dataEntities := make([]*entities.Portfolio, len(data.Data))

	for i, item := range data.Data {
		taxes:= make([]*entities.Tax, len(item.Taxes))
		for j, tax := range item.Taxes {
			taxes[j] = entities.InitTax(tax.ID, tax.TypeTax, tax.Rate)
		}
		itemData:= &entities.PortfolioEntityParams{
			ID: item.ID,
			Channel: item.Channel,
			Country: item.Country,
			CreateAt: item.CreateAt,
			CustomerID: item.CustomerID,
			Route: item.Route,
			SKU: item.SKU,
			Title: item.Title,
			CategoryID: item.CategoryID,
			Category: item.Category,
			Brand: item.Brand,
			Classification: item.Classification,
			UnitsPerBox: item.UnitsPerBox,
			MinOrderUnits: item.MinOrderUnits,
			PackageDescription: item.PackageDescription,
			PackageUnitDescription: item.PackageUnitDescription,
			QuantityMaxRedeem: item.QuantityMaxRedeem,
			RedeemUnit: item.RedeemUnit,
			OrderReasonRedeem: item.OrderReasonRedeem,
			SKURedeem: item.SKURedeem,
			Price: item.Price,
			Points: item.Points,
			Taxes: taxes,
		}
		dataEntities[i] = entities.InitPortfolio(itemData)
	}
	
	return dataEntities, data.Pagination
	
}

func (service *managerCacheAllPortfolioOfUserService) Set(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto, data []*entities.Portfolio, pagination *pagination_dtos.PaginationResponseDto) {

key:= generatePaths(params)

	dataCache := make([]*portfolioCacheParams, len(data))

	for i, item := range data {
		createAt := item.CreateAt()
		taxesCache := make([]*taxCacheParams, len(item.Taxes()))
		for j, tax := range item.Taxes() {
			taxesCache[j] = &taxCacheParams{
				ID: tax.ID(),
				TypeTax: tax.TypeTax(),
				Rate: tax.RateRaw(),
			}
		}
		dataCache[i] = &portfolioCacheParams{
			ID: item.ID(),
			Channel: item.Channel(),
			Country: item.Country(),
			CreateAt: &createAt,
			CustomerID: item.CustomerID(),
			Route: item.Route(),
			SKU: item.SKU(),
			Title: item.Title(),
			CategoryID: item.CategoryID(),
			Category: item.Category(),
			Brand: item.Brand(),
			Classification: item.Classification(),
			UnitsPerBox: fmt.Sprintf("%d", item.UnitsPerBox()),
			MinOrderUnits: fmt.Sprintf("%f", item.MinOrderUnits()),
			PackageDescription: item.PackageDescription(),
			PackageUnitDescription: item.PackageUnitDescription(),
			QuantityMaxRedeem: item.QuantityMaxRedeem(),
			RedeemUnit: item.RedeemUnit(),
			OrderReasonRedeem: item.OrderReasonRedeem(),
			SKURedeem: item.SKURedeem(),
			Price: item.FullPrice(),
			Points: item.Points(),
			Taxes: taxesCache,
		}
	}

	dataCacheStruct := &dataCacheStruct{
		Data: dataCache,
		Pagination: pagination,
	}

	service.managerCacheService.SetData(key, dataCacheStruct, time.Hour*24)
}

func generatePaths(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%v", params)))
	return fmt.Sprintf("customer/%s/currentPage/%d/pageSize/%d/hash/%s/all-portfolios", params.UserId, params.Page, params.Limit, hex.EncodeToString(hash[:]))
}
