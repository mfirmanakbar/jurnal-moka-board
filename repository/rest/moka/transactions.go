package moka

import "github.com/mfirmanakbar/jurnal-moka-board/domain/mokapos/transactions"

type RestTransactionsRepository interface {
	GetLatestTransactions(transactions.Request)
}

type TransactionsRepository struct{}

func NewRestTransactionsRepository() RestTransactionsRepository {
	return &TransactionsRepository{}
}

type GetLatestTransactionsParams struct {
	OutletId int64 `json:"outlet_id"`
}

func (t TransactionsRepository) GetLatestTransactions(request transactions.Request) {
	panic("implement me")
}
