package repository

import (
	"fmt"

	"AY1st/model"
	"AY1st/util"
	"AY1st/util/ptr"

	"github.com/go-xorm/xorm"
)

// ActionsInterface is health check (debug)
type ActionsInterface interface {
	Create(ubrID uint64, content string) (*model.Action, error)
	Get(id uint64) ([]*model.Action, error)
	Update(input *model.ActionUpdateInput) (*model.Action, error)
	GetOne(id uint64) (*model.Action, error)
}

// Actions is health check (debug)
type Actions struct {
	engine xorm.EngineInterface
}

// NewActions initializes Actions
func NewActions(engine xorm.EngineInterface) *Actions {
	u := Actions{
		engine: engine,
	}
	return &u
}

// Create は本の行動の新規登録
func (a *Actions) Create(ubrID uint64, content string) (*model.Action, error) {
	logger := util.GetLogger()

	action := &model.Action{}
	action.Common.UnsetDefaltCols()
	action.IsFinished = ptr.False()
	action.UserBookRegistrationID = ubrID
	action.Content = ptr.String(content)
	session := a.engine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return nil, err
	}

	affected, err := session.InsertOne(action)
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

	return action, nil
}

// Get はアクションを取得
func (a *Actions) Get(ubrID uint64) ([]*model.Action, error) {
	actions := []*model.Action{}
	// actions := new(model.ActionBody)
	// actions := make([]*model.ActionBody, 0)
	err := a.engine.Where("user_book_registration_id = ?", ubrID).Find(&actions)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get actions")
	}
	return actions, nil
}

// Update は投稿の更新
func (a *Actions) Update(input *model.ActionUpdateInput) (*model.Action, error) {

	value := &model.Action{}
	value.Common.ID = input.ID
	value.IsFinished = input.IsFinished
	value.Content = input.Content
	now := util.GetTimeNow()
	value.UpdatedAt = &now

	_, err := a.engine.ID(input.ID).Update(value)
	if err != nil {
		return nil, err
	}

	res, err := a.GetOne(input.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetOne は一意なアクションを取得
func (a *Actions) GetOne(id uint64) (*model.Action, error) {
	action := &model.Action{}
	ok, err := a.engine.ID(id).Get(action)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get action")
	}
	if !ok {
		return nil, fmt.Errorf("can not get action")
	}
	return action, nil
}
