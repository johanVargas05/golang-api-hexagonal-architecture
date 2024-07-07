package pagination_dtos

type PaginationResponseDto struct {
	CurrentPage     int
	PageSize        int
	HasNextPage     bool
	HasPreviousPage bool
	TotalItems			int
}
