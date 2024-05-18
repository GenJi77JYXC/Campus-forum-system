package repository

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"errors"

	"gorm.io/gorm"
)

type userTokenRepository struct{}

var UserTokenRepository = newUserTokenRepository()

func newUserTokenRepository() *userTokenRepository {
	return new(userTokenRepository)
}

func (r *userTokenRepository) Create(db *gorm.DB, userToken *model.UserToken) error {
	return db.Create(userToken).Error
}

func (r *userTokenRepository) UpdateStatusInvalidByToken(db *gorm.DB, token string) error {
	return db.Model(&model.UserToken{}).Where("token = ?", token).Update("status", 1).Error
}

func (r *userTokenRepository) UserStatusByToken(db *gorm.DB, userId int64) (*model.UserToken, error) {
	result := &model.UserToken{}
	err := db.Where("user_id = ?", userId).Last(result).Error
	if err != nil {
		logs.Logger.Error("数据库查询token出错：", err)
		return nil, err
	}
	return result, err
}

func (r *userTokenRepository) GetUserIDByToken(db *gorm.DB, token string) (*model.UserToken, error) {
	if token == "" {
		return nil, errors.New("token不能为空")
	}
	return r.take(db, "token = ?", token)
}

func (r *userTokenRepository) take(db *gorm.DB, column string, value interface{}) (*model.UserToken, error) {
	result := &model.UserToken{}
	// err := db.Where(column, value).Take(result).Error
	err := db.Where(column, value).Find(&result).Error
	// logs.Logger.Info("token_repo query result:", result)
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return result, nil
}
