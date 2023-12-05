package repository

import (
	"context"
	"fmt"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindListTransaction(context.Context, string) ([]dto.TransactionsRes, error)
	SearchTransaction(context.Context, string) string
	FilterTransaction(context.Context, string, string) (string, error)
	SortByTransaction(context.Context, string, string) (string, error)
	PaginationTransaction(context.Context, int, int) string
}

type transactionRepository struct {
	db *gorm.DB
}

var sortBy = map[string]string{
	"date":   "created_at",
	"amount": "amount",
	"to":     "recipient_id",
}

var sortType = map[string]string{
	"desc": "DESC",
	"asc":  "ASC",
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (tr *transactionRepository) FindListTransaction(ctx context.Context, sql string) (transactions []dto.TransactionsRes, err error) {
	raw := "SELECT * FROM transactions"
	raw += sql
	err = tr.db.WithContext(ctx).Raw(raw).Scan(&transactions).Error
	if err != nil {
		return nil, apperror.ErrFindListTransactionQuery
	}
	return transactions, nil
}

func (tr *transactionRepository) SearchTransaction(ctx context.Context, word string) (sql string) {
	raw := "%" + word + "%"
	sql = fmt.Sprintf("WHERE description ILIKE %s", raw)
	return
}

func (tr *transactionRepository) FilterTransaction(ctx context.Context, start string, end string) (sql string, err error) {
	_, err = time.Parse("2006-01-02", start)
	if err != nil {
		return "", apperror.ErrWrongStartDateFormat
	}
	_, err = time.Parse("2006-01-02", end)
	if err != nil {
		return "", apperror.ErrWrongEndDateFormat
	}
	sql = fmt.Sprintf("AND created_at BETWEEN %s AND %s", start, end)
	return sql, nil
}

func (tr *transactionRepository) SortByTransaction(ctx context.Context, sortByWord string, sort string) (sql string, err error) {
	valSortBy, ok1 := sortBy[sortByWord]
	if !ok1 {
		return "", apperror.ErrSortByTransactionQuery
	}
	valSortType, ok2 := sortType[sort]
	if !ok2 {
		return "", apperror.ErrSortTypeTrasacntionQueqry
	}
	sql = fmt.Sprintf("ORDER BY %s %s", valSortBy, valSortType)
	return sql, nil
}

func (tr *transactionRepository) PaginationTransaction(ctx context.Context, limit int, page int) (sql string) {
	sql = fmt.Sprintf("LIMIT %d OFFSET %d", limit, (page-1)*limit)
	return
}
