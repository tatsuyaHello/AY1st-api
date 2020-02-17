package service

import (
	"AY1st/model"
	"AY1st/repository"
	"AY1st/util/ptr"
)

// PostsInterface is
type PostsInterface interface {
	Create(userID uint64, bookAction *model.PostInput) (*model.PostInput, error)
	GetOne(id uint64) (*model.Post, error)
	GetAll() ([]*model.Post, error)
	Delete(id uint64) error
	Update(input []*model.ActionUpdateInput) ([]*model.Action, error)
	GetPostOfUser(userID uint64) ([]*model.PostOfUser, error)
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
	bo, err := p.BooksRepo.GetByIsbn(input.Isbn)
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
		AvatarURL:            user.AvatarURL,
		BookBody:             book.BookBody,
		Action:               actions,
	}

	return post, nil
}

// GetAll は全ての投稿を取得
func (p *Posts) GetAll() ([]*model.Post, error) {

	// ここでDBに対して何度かアクセスする
	posts, err := p.UsersBooksRegistrationsRepo.GetAll()
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "ubr not found")
	}

	for i, v := range posts {
		actions, err := p.ActionsRepo.Get(v.ID)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "action not found")
		}
		var acts []*model.Action
		for _, action := range actions {
			acts = append(acts, action)
		}
		posts[i].Action = acts
	}

	return posts, nil
}

// Delete は投稿削除
func (p *Posts) Delete(id uint64) error {
	_, err := p.UsersBooksRegistrationsRepo.GetOne(id)
	if err != nil {
		return model.NewError(model.ErrorResourceNotFound, "users_books_registration not found")
	}

	// 投稿の削除
	err = p.UsersBooksRegistrationsRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// Update は投稿を更新
func (p *Posts) Update(input []*model.ActionUpdateInput) ([]*model.Action, error) {

	var result []*model.Action
	var res *model.Action
	var err error
	// Actionの更新を行う。
	for _, v := range input {
		res, err = p.ActionsRepo.Update(v)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	// 今回更新するusers_book_registration_idに関するis_finishedが全てtrueになっていれば、is_action_completedをtrueにしてuserのtotal_priceを増加させる。
	actions, err := p.ActionsRepo.Get(res.UserBookRegistrationID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "actions not found")
	}
	isCompleted := true
	for _, v := range actions {
		if *v.IsFinished == *ptr.False() {
			isCompleted = *ptr.False()
		}
	}
	if isCompleted == *ptr.True() {
		// users_books_registrationsのis_action_completedをtrueにする
		err := p.UsersBooksRegistrationsRepo.Update(res.UserBookRegistrationID)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "users_books_registration cannot update")
		}
		// users_books_registrationsから情報を取得する
		ubrs, err := p.UsersBooksRegistrationsRepo.GetOne(res.UserBookRegistrationID)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "users_books_registration not found")
		}
		// booksから本の情報を取得する
		books, err := p.BooksRepo.GetOne(ubrs.BookID)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "book not found")
		}
		// usersのtotal_priceを更新する
		_, err = p.UsersRepo.UpdatePrice(ubrs.UserID, books.Price)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "user price cannot update")
		}
	}

	return result, nil
}

// GetPostOfUser は全ての投稿を取得
func (p *Posts) GetPostOfUser(userID uint64) ([]*model.PostOfUser, error) {

	// ここでDBに対して何度かアクセスする
	posts, err := p.UsersBooksRegistrationsRepo.GetPostOfUser(userID)
	if err != nil {
		return nil, model.NewError(model.ErrorResourceNotFound, "ubr not found")
	}

	for i, v := range posts {
		actions, err := p.ActionsRepo.Get(v.ID)
		if err != nil {
			return nil, model.NewError(model.ErrorResourceNotFound, "action not found")
		}
		var acts []*model.Action
		for _, action := range actions {
			acts = append(acts, action)
		}
		posts[i].Action = acts
	}

	return posts, nil
}
