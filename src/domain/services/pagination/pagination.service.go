package pagination_service

import pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"


type PaginationService struct {
	page		int
	pageSize	int
}

func New(page,pageSize int) *PaginationService {
	return &PaginationService{
		page,
		pageSize,
	}
}

func (service *PaginationService) Execute(totalItems int) *pagination_dtos.PaginationResponseDto {
	response := pagination_dtos.PaginationResponseDto{}
	response.CurrentPage = service.page
	response.PageSize = service.pageSize
	response.HasNextPage = totalItems > (service.page * service.pageSize)
	response.HasPreviousPage = service.page > 1
	response.TotalItems = totalItems

	return &response
}

func (service *PaginationService) GetOffset() int {
	limit := service.GetLimit()
	if service.page <= 0 {
		service.page = 1
	}
	return (service.page - 1) * limit
}

func (service *PaginationService) GetLimit() int {
	if service.pageSize <= 10 {
		service.pageSize = 10
	}
	return service.pageSize
}
