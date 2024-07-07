package portfolio_dtos

type ParamsGetAllPortfolioOfUserDto struct {
	UserId   string
	Page     int
	Limit    int
	Search   string
	SortType string
	SortBy   string
}