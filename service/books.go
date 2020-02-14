package service

import (
	"AY1st/model"
	"AY1st/repository"
)

// BooksInterface is
type BooksInterface interface {
	Create(userID uint64, bookAction *model.BookActionInput) (*model.BookActionInput, error)
}

// Books is
type Books struct {
	BooksRepo                   repository.BooksInterface
	ActionsRepo                 repository.ActionsInterface
	UsersBooksRegistrationsRepo repository.UsersBooksRegistrationsInterface
}

// NewBooks is
func NewBooks(BooksRepo repository.BooksInterface, UsersBooksRegistrationsRepo repository.UsersBooksRegistrationsInterface, ActionsRepo repository.ActionsInterface) *Books {
	b := Books{
		BooksRepo:                   BooksRepo,
		UsersBooksRegistrationsRepo: UsersBooksRegistrationsRepo,
		ActionsRepo:                 ActionsRepo,
	}
	return &b
}

// Create はPost新規登録
func (b *Books) Create(userID uint64, input *model.BookActionInput) (*model.BookActionInput, error) {

	//1. booksテーブルに既に存在するか否かを見る
	bo, err := b.BooksRepo.GetByRakutenID(input.RakutenID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
	}
	//1.1. 存在しなければ保存する。
	book := &model.Book{}
	if bo == nil {
		book, err = b.BooksRepo.Create(&input.BookBody)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
		}
	}
	if bo != nil {
		book = bo
	}

	//2. booksテーブルとusersテーブルの結びつけをusers_books_registrationsに保存する。
	ubr, err := b.UsersBooksRegistrationsRepo.Create(userID, book.ID)
	if err != nil {
		return nil, model.NewError(model.ErrorCannotCreate, "Cannot create")
	}

	//3. users_books_registrationsテーブルのIDを使用して、actionをactionsテーブルに保存する。
	for _, v := range input.ActionInputBody {
		_, err := b.ActionsRepo.Create(ubr.ID, v.Content)
		if err != nil {
			return nil, model.NewError(model.ErrorCannotCreate, "Cannot create")
		}
	}

	//4. フロントにレスポンスを返す。
	return input, nil
}
