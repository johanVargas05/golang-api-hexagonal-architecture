package all_portfolios_of_user_service

import (
	"errors"

	pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"
	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	portfolio_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/portfolio"
	pagination_service "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/services/pagination"
	validate_object "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/validate_objects"
)

type AllPortFolioOfUserService struct {
	allPortFolioOfUserQueryRepository portfolio_ports.AllPortFolioOfUserQueryRepositoryPort

}

func New(allPortFolioOfUserQueryRepository portfolio_ports.AllPortFolioOfUserQueryRepositoryPort,
	) *AllPortFolioOfUserService {
	return &AllPortFolioOfUserService{
		allPortFolioOfUserQueryRepository,
	}
}

func (service *AllPortFolioOfUserService) Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio,*pagination_dtos.PaginationResponseDto, error) {
	search:=validate_object.NewStringValueObject(params.Search, "search").IsOptional().MaxLength(100).MinLength(3).TransformLowerCase()
	sortType:=validate_object.NewStringValueObject(params.SortType, "sortType").IsOptional().MaxLength(4).MinLength(3).TransformLowerCase()
	sortBy:=validate_object.NewStringValueObject(params.SortBy, "sortBy").IsOptional().MaxLength(16).MinLength(5).TransformSnakeCase()
	page:=validate_object.NewNumberValueObject(params.Page, "page").IsDifferentZero().IsPositive()
	pageSize:=validate_object.NewNumberValueObject(params.Limit, "pageSize").IsDifferentZero().IsPositive()
	UserId:=validate_object.NewStringValueObject(params.UserId, "userId").IsID()

	err:=search.Validate()

	if err!=nil{
		return nil, nil, err
	}

	err=sortType.Validate()

	if err!=nil{
		return nil,nil, err
	}

	sortTypeValue:=sortType.Value()

	if sortTypeValue==""{
		sortTypeValue="desc"
	}

	if sortTypeValue!="asc" && sortTypeValue!="desc"{
		return nil,nil, errors.New("sortType must be asc or desc")
	}

	err=sortBy.Validate()

	if err!=nil{
		return nil,nil, err
	}

	sortByValue:=sortBy.Value()

	if sortByValue==""{
		sortByValue="created_at"
	}

	err=page.Validate()

	if err!=nil{

		return nil,nil, err
	}

	err=pageSize.Validate()

	if err!=nil{
		return nil,nil, err
	}

	err=UserId.Validate()

	if err!=nil{
		return nil,nil, err
	}

	params.Search=search.Value()
	params.SortType=sortTypeValue
	params.SortBy=sortByValue
	params.Page=page.Value()
	params.Limit=pageSize.Value()
	params.UserId=UserId.Value()
	
	result,totalItems,err:=service.allPortFolioOfUserQueryRepository.Execute(params)

	if err!=nil{
		return nil,nil, err
	}

	paginationResponse:=pagination_service.New(params.Page,params.Limit).Execute(totalItems)

	return result,paginationResponse,nil

}