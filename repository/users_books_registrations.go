package repository

import (
	"AY1st/model"
	"AY1st/util"
	"fmt"

	"github.com/go-xorm/xorm"
)

// UsersBooksRegistrationsInterface is
type UsersBooksRegistrationsInterface interface {
	Create(userID, bookID uint64) (*model.UserBookRegistration, error)
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
	ubrs.IsActionCompleted = 0
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
