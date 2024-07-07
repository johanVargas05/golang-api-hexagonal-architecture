package portfolio_errors

import "errors"

var (
	ErrNotFoundPortfoliosOfUser = errors.New("portfolios of user not found")
)