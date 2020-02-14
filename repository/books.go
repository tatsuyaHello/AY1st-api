package repository

import (
	"fmt"

	"AY1st/model"
	"AY1st/util"

	"github.com/go-xorm/xorm"
)

// BooksInterface is health check (debug)
type BooksInterface interface {
	Create(input *model.BookBody) (*model.Book, error)
	GetByRakutenID(rakutenID string) (*model.Book, error)
}

// Books is health check (debug)
type Books struct {
	engine xorm.EngineInterface
}

// NewBooks initializes Books
func NewBooks(engine xorm.EngineInterface) *Books {
	u := Books{
		engine: engine,
	}
	return &u
}

// Create は本の新規登録
func (b *Books) Create(input *model.BookBody) (*model.Book, error) {
	logger := util.GetLogger()

	book := &model.Book{}
	book.Common.UnsetDefaltCols()
	book.BookBody = *input
	session := b.engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return nil, err
	}

	affected, err := session.InsertOne(book)
	if err != nil {
		session.Rollback()
		logger.Error(err)
		return nil, err
	}
	if affected != 1 {
		session.Rollback()
		logger.Error(err)
		return nil, fmt.Errorf("no book was created")
	}
	err = session.Commit()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	outputBook, err := b.GetByRakutenID(input.RakutenID)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get book")
	}

	return outputBook, nil
}

// GetByRakutenID は一意なユーザーを取得
func (b *Books) GetByRakutenID(id string) (*model.Book, error) {
	book := &model.Book{}
	ok, err := b.engine.Where("rakuten_id = ?", id).Get(book)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get book")
	}
	if !ok {
		return nil, nil
	}
	return book, nil
}
