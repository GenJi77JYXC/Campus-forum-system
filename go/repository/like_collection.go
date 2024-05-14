package repository

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"fmt"

	"gorm.io/gorm"
)

type lcRepository struct{}

var LCRepository = newLCRepository()

func newLCRepository() *lcRepository {
	return &lcRepository{}
}

func (r *lcRepository) CreateLike(db *gorm.DB, op *model.UserLikeArticle) error {
	return db.Create(op).Error
}

// 通过用户ID和文章ID获取用户点赞操作记录
func (r *lcRepository) GetUserLikeOperation(db *gorm.DB, userID, articleID int64) (*model.UserLikeArticle, error) {
	return r.takeLikeOne(db, map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	})
}

func (r *lcRepository) takeLikeOne(db *gorm.DB, params map[string]interface{}) (*model.UserLikeArticle, error) {
	opHis := &model.UserLikeArticle{}
	err := db.Where(params).Limit(1).Find(&opHis).Error
	if err != nil {
		logs.Logger.Error("query db error:", err)
		return nil, err
	}
	return opHis, nil
}

// 通过用户ID和文章ID更新用户点赞操作记录
func (r *lcRepository) UpdateUserLikeOperation(db *gorm.DB, userID int64, articleID int64, params map[string]interface{}) error {
	return db.Model(&model.UserLikeArticle{}).Where("user_id = ? and article_id = ?", userID, articleID).Updates(params).Error
}

// 创建收藏记录
func (r *lcRepository) CreateFavorite(db *gorm.DB, op *model.UserFavoriteArticle) error {
	return db.Create(op).Error
}

// 通过用户ID和文章ID获取用户收藏操作记录
func (r *lcRepository) GetUserFavoriteOperation(db *gorm.DB, userID int64, articleID int64) (*model.UserFavoriteArticle, error) {
	return r.takeFavoriteOne(db, map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	})
}

func (r *lcRepository) takeFavoriteOne(db *gorm.DB, params map[string]interface{}) (*model.UserFavoriteArticle, error) {
	opHis := &model.UserFavoriteArticle{}
	err := db.Where(params).Limit(1).Find(&opHis).Error
	if err != nil {
		logs.Logger.Error("query db error:", err)
		return nil, err
	}
	return opHis, nil
}

// 通过用户ID和文章ID更新用户收藏操作记录
func (r *lcRepository) UpdateUserFavoriteOperation(db *gorm.DB, userID int64, articleID int64, params map[string]interface{}) error {
	return db.Model(&model.UserFavoriteArticle{}).Where("user_id = ? and article_id = ?", userID, articleID).Updates(params).Error
}

// 获取用户收藏记录
func (r *lcRepository) GetFavoriteRecords(db *gorm.DB, userID int64, cursorTime int64, limit int, sortby string, order string) []model.UserFavoriteArticle {
	var records []model.UserFavoriteArticle

	db.Where("update_time < ? and user_id = ?", cursorTime, userID).Order(fmt.Sprintf("%s %s", sortby, order)).Where("status = ?", 1).Limit(limit).Find(&records)
	return records
}
