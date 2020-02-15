package service

import (
	"AY1st/model"
	"AY1st/repository"
)

// PostsInterface is
type PostsInterface interface {
	Create(userID uint64, bookAction *model.PostInput) (*model.PostInput, error)
	GetOne(id uint64) (*model.Post, error)
}

// Posts is
type Posts struct {
	UsersRepo                   repository.UsersInterface
	BooksRepo                   repository.BooksInterface
	ActionsRepo                 repository.ActionsInterface
	UsersBooksRegistrationsRepo repository.UsersBooksRegistrationsInterface
}

// NewPosts is
func NewPosts(UsersRepo repository.UsersInterface, BooksRepo repository.BooksInterface, UsersBooksRegistrationsRepo repository.UsersBooksRegistrationsInterface, ActionsRepo repository.ActionsInterface) *Posts {
	b := Posts{
		UsersRepo:                   UsersRepo,
		BooksRepo:                   BooksRepo,
		UsersBooksRegistrationsRepo: UsersBooksRegistrationsRepo,
		ActionsRepo:                 ActionsRepo,
	}
	return &b
}

// Create はPost新規登録
func (p *Posts) Create(userID uint64, input *model.PostInput) (*model.PostInput, error) {

	//1. booksテーブルに既に存在するか否かを見る
	bo, err := p.BooksRepo.GetByRakutenID(input.RakutenID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
	}
	//1.1. 存在しなければ保存する。
	book := &model.Book{}
	if bo == nil {
		book, err = p.BooksRepo.Create(&input.BookBody)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
		}
	}
	if bo != nil {
		book = bo
	}

	//2. booksテーブルとusersテーブルの結びつけをusers_books_registrationsに保存する。
	ubr, err := p.UsersBooksRegistrationsRepo.Create(userID, book.ID)
	if err != nil {
		return nil, model.NewError(model.ErrorCannotCreate, "Cannot create")
	}

	//3. users_books_registrationsテーブルのIDを使用して、actionをactionsテーブルに保存する。
	for _, v := range input.ActionInputBody {
		_, err := p.ActionsRepo.Create(ubr.ID, v.Content)
		if err != nil {
			return nil, model.NewError(model.ErrorCannotCreate, "Cannot create")
		}
	}

	//4. フロントにレスポンスを返す。
	return input, nil
}

// GetOne は一意な投稿を取得
func (p *Posts) GetOne(id uint64) (*model.Post, error) {

	// ここでDBに対して何度かアクセスする
	ubr, err := p.UsersBooksRegistrationsRepo.GetOne(id)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "ubr not found")
	}

	user, err := p.UsersRepo.GetOne(ubr.UserID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "user not found")
	}

	book, err := p.BooksRepo.GetOne(ubr.BookID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
	}

	actions, err := p.ActionsRepo.Get(ubr.ID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "actions not found")
	}

	// いい感じのモデルに上記で得られた値を代入する
	post := &model.Post{
		UserBookRegistration: *ubr,
		DisplayName:          user.DisplayName,
		AvartarURL:           user.AvartarURL,
		BookBody:             book.BookBody,
		ActionBody:           actions,
	}

	return post, nil
}
