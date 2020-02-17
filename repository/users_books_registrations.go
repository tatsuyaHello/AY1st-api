package repository

import (
	"AY1st/model"
	"AY1st/util"
	"AY1st/util/ptr"
	"fmt"
	"strings"

	"github.com/go-xorm/xorm"
)

// UsersBooksRegistrationsInterface is
type UsersBooksRegistrationsInterface interface {
	Create(userID, bookID uint64) (*model.UserBookRegistration, error)
	GetOne(id uint64) (*model.UserBookRegistration, error)
	GetAll() ([]*model.Post, error)
	Delete(id uint64) error
	Update(id uint64) error
	GetPostOfUser(userID uint64) ([]*model.PostOfUser, error)
}

// UsersBooksRegistrations is health check (debug)
type UsersBooksRegistrations struct {
	engine xorm.EngineInterface
}

// NewUsersBooksRegistrations initializes UsersBooksRegistrations
func NewUsersBooksRegistrations(engine xorm.EngineInterface) *UsersBooksRegistrations {
	u := UsersBooksRegistrations{
		engine: engine,
	}
	return &u
}

// Create はユーザ本の登録の新規
func (ubr *UsersBooksRegistrations) Create(userID, bookID uint64) (*model.UserBookRegistration, error) {
	logger := util.GetLogger()

	ubrs := &model.UserBookRegistration{}
	ubrs.Common.UnsetDefaltCols()
	ubrs.UserID = userID
	ubrs.BookID = bookID
	ubrs.IsActionCompleted = ptr.False()
	session := ubr.engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return nil, err
	}

	affected, err := session.InsertOne(ubrs)
	if err != nil {
		session.Rollback()
		logger.Error(err)
		return nil, err
	}
	if affected != 1 {
		session.Rollback()
		logger.Error(err)
		return nil, fmt.Errorf("no user book registration was created")
	}
	err = session.Commit()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	res := &model.UserBookRegistration{}
	ok, err := ubr.engine.Where("user_id = ? AND book_id = ?", userID, bookID).Get(res)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get book")
	}
	if !ok {
		return nil, fmt.Errorf("can not get book")
	}

	return res, nil
}

// GetOne は一意な投稿の結びつきを取得
func (ubr *UsersBooksRegistrations) GetOne(id uint64) (*model.UserBookRegistration, error) {
	ubrs := &model.UserBookRegistration{}
	ok, err := ubr.engine.ID(id).Get(ubrs)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get ubr")
	}
	if !ok {
		return nil, fmt.Errorf("can not get ubr")
	}
	return ubrs, nil
}

// GetAll は全ての投稿を取得
func (ubr *UsersBooksRegistrations) GetAll() ([]*model.Post, error) {

	s := ubr.engine.NewSession()

	posts := []*model.Post{}

	s = s.Select(strings.Join([]string{
		"users_books_registrations.*",
		"users.display_name",
		"users.avatar_url",
		"books.rakuten_url",
		"books.rakuten_review",
		"books.isbn",
		"books.title",
		"books.price",
		"books.author",
		"books.book_img_url",
	}, " ,"))

	s.Join("LEFT", "users", "users.id = users_books_registrations.user_id")
	s.Join("LEFT", "books", "books.id = users_books_registrations.book_id")

	err := s.Find(&posts)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get posts")
	}
	return posts, nil
}

// Delete は投稿の削除
func (ubr *UsersBooksRegistrations) Delete(id uint64) error {
	session := ubr.engine.NewSession()
	defer session.Close()
	err := session.Begin()

	ubrs := &model.UserBookRegistration{}
	affected, err := ubr.engine.ID(id).Delete(ubrs)
	if err != nil {
		session.Rollback()
		return err
	}
	if affected != 1 {
		session.Rollback()
		return fmt.Errorf("could not delete users_books_registration-id = %v", id)
	}

	err = session.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Update は投稿の更新
func (ubr *UsersBooksRegistrations) Update(id uint64) error {

	ubrs := &model.UserBookRegistration{}

	ubrs.IsActionCompleted = ptr.True()
	now := util.GetTimeNow()
	ubrs.Common.UpdatedAt = &now
	_, err := ubr.engine.ID(id).Update(ubrs)
	if err != nil {
		return err
	}
	return nil
}

// GetPostOfUser は全ての投稿を取得
func (ubr *UsersBooksRegistrations) GetPostOfUser(userID uint64) ([]*model.PostOfUser, error) {

	s := ubr.engine.NewSession()

	posts := []*model.PostOfUser{}

	s = s.Table("users_books_registrations")

	s = s.Where("user_id = ?", userID)

	s = s.Select(strings.Join([]string{
		"users_books_registrations.*",
		"books.rakuten_url",
		"books.rakuten_review",
		"books.isbn",
		"books.title",
		"books.price",
		"books.author",
		"books.book_img_url",
	}, " ,"))

	s.Join("LEFT", "books", "books.id = users_books_registrations.book_id")

	err := s.Find(&posts)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get posts")
	}
	return posts, nil
}
