package get_all_portfolios_of_suer_controller

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	portfolio_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/portfolio"
	portfolio_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/portfolio"
)
type bodyParams struct {
	Search  string `json:"search"`
	SortType string `json:"sort_type"`
	SortBy string `json:"sort_by"`
	Page int `json:"current_page"`
	Limit int `json:"page_size"`
}

type responseDataBody struct {
	ID string `json:"id"`
	Channel string `json:"channel"`
	Country string `json:"country"`
	CreateAt string `json:"create_at"`
	CustomerID string `json:"customer_code"`
	Route string `json:"route"`
	SKU string `json:"sku"`
	Title string `json:"title"`
	CategoryID string `json:"category_id"`
	Category string `json:"category"`
	Brand string `json:"brand"`
	Classification string `json:"classification"`
	UnitsPerBox int `json:"units_per_box"`
	MinOrderUnits float64 `json:"min_order_units"`
	PackageDescription string `json:"package_description"`
	PackageUnitDescription string `json:"package_unit_description"`
	QuantityMaxRedeem int `json:"quantity_max_redeem"`
	RedeemUnit string `json:"redeem_unit"`
	OrderReasonRedeem int `json:"order_reason_redeem"`
	SKURedeem bool `json:"sku_redeem"`
	Price float64 `json:"price"`
	Points 		int `json:"points"`
	Taxes	[]map[string]interface{} `json:"taxes"`
}
type GetAllPortfoliosOfUserController struct {
	useCase portfolio_ports.AllPortFolioOfUserQueryUseCasePort
}

func New(useCase portfolio_ports.AllPortFolioOfUserQueryUseCasePort) *GetAllPortfoliosOfUserController {
	return &GetAllPortfoliosOfUserController{
		useCase,
	}
}

func (controller *GetAllPortfoliosOfUserController) Execute(ctx echo.Context) error {
	body:=&bodyParams{}
	consumerId:=ctx.Param("consumerId")
	ctx.Bind(body)

	params:=&portfolio_dtos.ParamsGetAllPortfolioOfUserDto{
		Search: body.Search,
		SortType: body.SortType,
		SortBy: body.SortBy,
		Page: body.Page,
		Limit: body.Limit,
		UserId: consumerId,
	}

	data,pagination,err:=controller.useCase.Execute(params)

	if err!=nil{
		if errors.Is(err, portfolio_errors.ErrNotFoundPortfoliosOfUser){
			return ctx.JSON(404, map[string]interface{}{
				"code": "404",
				"message": "Not found",
				"error": err.Error(),
			})
		}
		return ctx.JSON(500, map[string]interface{}{
			"code": "500",
			"message": "Internal server error",
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, map[string]interface{}{
		"code": "200",
		"message": "Success",
		"data": mapperEntityToResponse(data),
		"pagination": map[string]interface{}{
			"current_page": pagination.CurrentPage,
			"page_size": pagination.PageSize,
			"hast_next": pagination.HasNextPage,
			"has_previous": pagination.HasPreviousPage,
			"total_items": pagination.TotalItems,
		},
	})

}

func mapperEntityToResponse(entities []*entities.Portfolio) []*responseDataBody {
	taxes := make([]map[string]interface{}, len(entities[0].Taxes()))
	for i, tax := range entities[0].Taxes() {
		taxes[i] = map[string]interface{}{
			"tax_type": tax.TypeTax(),
			"percentage": tax.Rate(),
			"tax_id": tax.ID(),
		}
	
	}

	var response []*responseDataBody = make([]*responseDataBody, len(entities)) 
	for i, entity := range entities {
		response[i] = &responseDataBody{
			ID: entity.ID(),
			Channel: entity.Channel(),
			Country: entity.Country(),
			CreateAt: fmt.Sprintf("%v", entity.CreateAt()),
			CustomerID: entity.CustomerID(),
			Route: entity.Route(),
			SKU: entity.SKU(),
			Title: entity.Title(),
			CategoryID: entity.CategoryID(),
			Category: entity.Category(),
			Brand: entity.Brand(),
			Classification: entity.Classification(),
			UnitsPerBox: entity.UnitsPerBox(),
			MinOrderUnits: entity.MinOrderUnits(),
			PackageDescription: entity.PackageDescription(),
			PackageUnitDescription: entity.PackageUnitDescription(),
			QuantityMaxRedeem: entity.QuantityMaxRedeem(),
			RedeemUnit: entity.RedeemUnit(),
			OrderReasonRedeem: entity.OrderReasonRedeem(),
			SKURedeem: entity.SKURedeem(),
			Price: entity.Price(),
			Points: entity.Points(),
			Taxes: taxes,
		}
	}
	return response
}