package repository

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"

	"gorm.io/gorm"
)

type commentRepository struct {
}

var CommentRepository = newCommentRepository()

func newCommentRepository() *commentRepository {
	return &commentRepository{}
}

// 创建评论
func (r *commentRepository) Create(db *gorm.DB, comment *model.Comment) error {
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// 获取评论列表 通过时间排序
func (r *commentRepository) GetCommentsByCursorTime(db *gorm.DB, articleID, cursorTime int64) ([]model.Comment, error) {
	var comments []model.Comment
	err := db.Where("create_time < ?", cursorTime).Where("article_id = ?", articleID).Limit(30).Find(&comments).Error
	if err != nil {
		logs.Logger.Errorf("GetCommentsByCursorTime: query db error: %v", err)
		return nil, err
	}
	return comments, err
}

// 通过评论id获取评论
func (r *commentRepository) GetCommentsByCommentID(db *gorm.DB, id int64) (*model.Comment, error) {
	return r.takeOne(db, "id = ?", id)
}

func (r *commentRepository) takeOne(db *gorm.DB, column string, value interface{}) (*model.Comment, error) {
	comment := &model.Comment{}
	err := db.Where(column, value).Find(&comment).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return comment, nil
}

// 通过文章id获取评论列表
func (r *commentRepository) GetCommentsByArticleID(db *gorm.DB, articleID int64) ([]model.Comment, error) {
	return r.takeList(db, "article_id = ?", articleID)
}

func (r *commentRepository) takeList(db *gorm.DB, column string, value interface{}) ([]model.Comment, error) {
	var comments []model.Comment
	err := db.Where(column, value).Find(&comments).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return comments, nil
}

// 删除评论
func (r *commentRepository) DeleteCommentByID(db *gorm.DB, id int64) error {
	err := db.Delete(&model.Comment{}, id).Error
	if err != nil {
		logs.Logger.Errorf("delete comment error: %v", err)
		return err
	}
	return nil
}
