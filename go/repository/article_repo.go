package repository

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"fmt"

	"gorm.io/gorm"
)

type articleRepository struct {
}

// ArticleRepository is the entrance as a convenient interface
var ArticleRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return new(articleRepository)
}

func (r *articleRepository) Create(db *gorm.DB, article *model.Article) error {
	return db.Create(article).Error
}

func (r *articleRepository) GetArticleFields(db *gorm.DB, authorID int64, fields []string, cursorTime int64, limit int, sortby string, order string) []model.Article {
	var articles []model.Article
	if authorID == 0 {
		db.Where("create_time < ?", cursorTime).Select(fields).Order(fmt.Sprintf("%s %s", sortby, order)).Limit(limit).Find(&articles)
	} else {
		db.Where("user_id = ? and create_time < ?", authorID, cursorTime).Select(fields).Order(fmt.Sprintf("%s %s", sortby, order)).Limit(limit).Find(&articles)
	}
	return articles
}

func (r *articleRepository) GetArticleByID(db *gorm.DB, id int64) (*model.Article, error) {
	return r.take(db, "id = ?", id)
}

func (r *articleRepository) DeleteArticleByID(db *gorm.DB, id int64) error {
	return db.Where("id = ?", id).Delete(&model.Article{}).Error
}

func (r *articleRepository) UpdateArticleByID(db *gorm.DB, articleID int64, title, content string) error {
	err := db.Exec("update articles set title = ?, content = ? where id = ?", title, content, articleID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) take(db *gorm.DB, column string, value interface{}) (*model.Article, error) {
	result := new(model.Article)
	if err := db.Where(column, value).Find(&result).Error; err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, err
	}
	return result, nil
}
