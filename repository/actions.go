package repository

import (
	"fmt"

	"AY1st/model"
	"AY1st/util"

	"github.com/go-xorm/xorm"
)

// ActionsInterface is health check (debug)
type ActionsInterface interface {
	Create(ubrID uint64, content string) (*model.Action, error)
	Get(id uint64) ([]*model.ActionBody, error)
	Update(input *model.Action) (*model.Action, error)
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
	action.IsFinished = false
	action.UserBookRegistrationID = ubrID
	action.Content = content
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
func (a *Actions) Get(id uint64) ([]*model.ActionBody, error) {
	actions := []*model.ActionBody{}
	// actions := new(model.ActionBody)
	// actions := make([]*model.ActionBody, 0)
	err := a.engine.ID(id).Find(&actions)
	if err != nil {
		util.GetLogger().Error(err)
		return nil, fmt.Errorf("can not get actions")
	}
	return actions, nil
}

// Update は投稿の更新
func (a *Actions) Update(input *model.Action) (*model.Action, error) {

	action := &model.Action{}

	action.ActionBody = input.ActionBody
	action.ID = input.ID
	now := util.GetTimeNow()
	action.Common.UpdatedAt = &now
	_, err := a.engine.ID(action.ID).Update(action)
	if err != nil {
		return nil, err
	}
	return action, nil
}
