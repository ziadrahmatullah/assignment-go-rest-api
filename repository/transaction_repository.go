package repository

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	FindListTransaction(context.Context, dto.ListTransactionsReq) (*dto.TransactionPaginationRes, error)
	TopUpTransaction(context.Context, model.Transaction) (*model.Transaction, error)
	TransferTransaction(context.Context, model.Transaction) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

var sortBy = map[string]string{
	"date":   "created_at",
	"amount": "amount",
	"to":     "receiver",
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

func (tr *transactionRepository) FindListTransaction(ctx context.Context, req dto.ListTransactionsReq) (*dto.TransactionPaginationRes, error) {
	raw := "SELECT * FROM transactions "
	searchSql := tr.SearchTransaction(req.Search)
	filterSql, err := tr.FilterTransaction(req.FilterStart, req.FilterEnd, searchSql)
	if err != nil {
		return nil, err
	}
	sortSql, err := tr.SortByTransaction(req.SortBy, req.SortType)
	if err != nil {
		return nil, err
	}
	page, limit, paginationSql, err := tr.PaginationTransaction(req.PaginationLimit, req.PaginationPage)
	if err != nil {
		return nil, err
	}
	var transactions []model.Transaction
	var totalData int
	rawCount := "SELECT count(*) FROM transactions " + searchSql + filterSql
	raw += searchSql + filterSql + sortSql + paginationSql
	err = tr.db.WithContext(ctx).Model(&model.Transaction{}).Raw(rawCount).Scan(&totalData).Error
	if err != nil {
		return nil, apperror.ErrFindListTransactionQuery
	}
	totalPage := math.Ceil(float64(totalData) / float64(limit))
	if page > int(totalPage) {
		return nil, apperror.ErrPageNotFound
	}
	err = tr.db.WithContext(ctx).Table("transactions").Raw(raw).Scan(&transactions).Error
	if err != nil {
		return nil, apperror.ErrFindListTransactionQuery
	}
	res := &dto.TransactionPaginationRes{
		Data:      transactions,
		TotalData: totalData,
		TotalPage: int(totalPage),
		Page:      page,
	}
	return res, nil
}

func (tr *transactionRepository) SearchTransaction(word *string) (sql string) {
	if word == nil || *word == "" {
		return ""
	}
	raw := "%" + *word + "%"
	sql = fmt.Sprintf("WHERE description ILIKE '%s' ", raw)
	return
}

func (tr *transactionRepository) FilterTransaction(start, end *string, prevSql string) (sql string, err error) {
	if (start == nil || *start == "") && (end == nil || *end == "") {
		return "", nil
	}
	if (start != nil || *start != "") && (end == nil || *end == "") {
		return "", apperror.ErrInvalidFilterFormat
	}
	if (start == nil || *start == "") && (end != nil || *end != "") {
		return "", apperror.ErrInvalidFilterFormat
	}
	_, err = time.Parse("2006-01-02", *start)
	if err != nil {
		return "", apperror.ErrWrongStartDateFormat
	}
	_, err = time.Parse("2006-01-02", *end)
	if err != nil {
		return "", apperror.ErrWrongEndDateFormat
	}
	if prevSql == "" {
		sql = fmt.Sprintf("WHERE created_at BETWEEN '%s' AND '%s' ", *start, *end)
	} else {
		sql = fmt.Sprintf("AND created_at BETWEEN '%s' AND '%s' ", *start, *end)
	}
	return sql, nil
}

func (tr *transactionRepository) SortByTransaction(sortByWord, sort *string) (sql string, err error) {
	if (sortByWord == nil || *sortByWord == "") && (*sort == "" || sort == nil){
		return "", nil
	}
	if (sortByWord == nil || *sortByWord == "") && (*sort != "" || sort != nil){
		return "", apperror.ErrInvalidSortFormat
	}
	valSortBy, ok := sortBy[*sortByWord]
	if !ok {
		return "", apperror.ErrSortByTransactionQuery
	}
	var valSortType string
	if *sort == "" || sort == nil {
		valSortType = sortType["desc"]
	} else {
		valSortType, ok = sortType[*sort]
		if !ok {
			return "", apperror.ErrSortTypeTrasacntionQueqry
		}
	}
	sql = fmt.Sprintf("ORDER BY %s %s ", valSortBy, valSortType)
	return sql, nil
}

func (tr *transactionRepository) PaginationTransaction(limit, page *string) (pageRest, limitRes int, sql string, err error) {
	intLimit := 10
	intPage := 1
	if limit != nil && *limit != "" {
		intLimit, err = strconv.Atoi(*limit)
		if err != nil {
			return intPage, intLimit, "", apperror.ErrInvalidPagination
		}
	}
	if page != nil && *page != "" {
		intPage, err = strconv.Atoi(*page)
		if err != nil {
			return intPage, intLimit, "", apperror.ErrInvalidPagination
		}
	}
	sql = fmt.Sprintf("LIMIT %d OFFSET %d", intLimit, (intPage-1)*intLimit)
	return intPage, intLimit, sql, nil
}

func (tr *transactionRepository) TopUpTransaction(ctx context.Context, req model.Transaction) (transaction *model.Transaction, err error) {
	tx := tr.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Table("wallets").
		Where("id = ?", req.WalletId).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("balance", gorm.Expr("balance + ?", req.Amount))
	tx.Table("transactions").Create(&req)
	if req.Amount.GreaterThan(model.AmountReward) {
		tx.Table("attempts").
			Where("wallet_id = ?", req.WalletId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Update("remaining_attempt", gorm.Expr("remaining_attempt + ?", 1))
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	return &req, nil
}

func (tr *transactionRepository) TransferTransaction(ctx context.Context, req model.Transaction) (transaction *model.Transaction, err error) {
	tx := tr.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Table("wallets").
		Where("wallet_number = ?", req.Sender).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("balance", gorm.Expr("balance - ?", req.Amount))
	var receiver model.Wallet
	tx.Table("wallets").
		Where("wallet_number = ?", req.Receiver).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("balance", gorm.Expr("balance + ?", req.Amount)).Scan(&receiver)
	tx.Table("transactions").Create(&req)
	receiverHistory := &model.Transaction{
		WalletId:        receiver.ID,
		TransactionType: req.TransactionType,
		Sender:          req.Sender,
		Receiver:        req.Receiver,
		Amount:          req.Amount,
		Description:     req.Description,
	}
	tx.Table("transactions").Create(&receiverHistory)
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	return &req, nil
}
