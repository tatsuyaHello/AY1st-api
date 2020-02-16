package service

import (
	"AY1st/model"
	"AY1st/repository"
)

// BooksInterface is
type BooksInterface interface {
	GetOne(id uint64) (*model.Book, error)
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
