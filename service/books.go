package service

import (
	"AY1st/model"
	"AY1st/repository"
)

// BooksInterface is
type BooksInterface interface {
	GetOne(id uint64) (*model.Book, error)
	Create(input *model.BookBody) (*model.Book, error)
}

// Books is
type Books struct {
	BooksRepo repository.BooksInterface
}

// NewBooks is
func NewBooks(BooksRepo repository.BooksInterface) *Books {
	b := Books{
		BooksRepo: BooksRepo,
	}
	return &b
}

// GetOne は一意な本を取得
func (b *Books) GetOne(id uint64) (*model.Book, error) {

	bo, err := b.BooksRepo.GetOne(id)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
	}

	return bo, nil
}

// Create はPost新規登録
func (b *Books) Create(input *model.BookBody) (*model.Book, error) {

	//1. booksテーブルに既に存在するか否かを見る
	bo, err := b.BooksRepo.GetByIsbn(input.Isbn)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
	}
	//1.1. 存在しなければ保存する。
	book := &model.Book{}
	if bo == nil {
		book, err = b.BooksRepo.Create(input)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
		}
	}
	if bo != nil {
		book = bo
	}

	//4. フロントにレスポンスを返す。
	return book, nil
}
