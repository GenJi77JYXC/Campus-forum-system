package service

import (
	"Campus-forum-system/database"
	"Campus-forum-system/model"
	"Campus-forum-system/repository"
	"Campus-forum-system/util"
	"errors"

	"gorm.io/gorm"
)

const (
	LikeArticle     = 1
	FavoriteArticle = 1
)

type lcService struct {
}

// LCService是入口 作为方便的接口
var LCService = newLCService()

func newLCService() *lcService {
	return &lcService{}
}

func (s *lcService) PostLikeArticle(userID, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserLikeOperation(database.GetDB(), userID, articleID)
	if err != nil {
		return errors.New("数据库查询失败")
	}
	// 已经点赞成功
	if opHis.Status == 1 {
		return nil
	}
	// Transaction:事务启动时将事务作为块处理，返回错误将回滚，否则要提交。事务在事务的fc中执行任意数量的命令。一旦成功，就会做出改变;如果发生错误，它们将被回滚。
	err = database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		if opHis.ID == 0 {
			err = repository.LCRepository.UpdateUserLikeOperation(database.GetDB(), userID, articleID, map[string]interface{}{"status": 1})
		} else {
			err = repository.LCRepository.CreateLike(database.GetDB(), &model.UserLikeArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				UpdateTime: util.NowTimestamp(),
			})
		}
		if err != nil {
			return err
		}
		// 更新文章的喜欢数量
		err = database.GetDB().Exec("update article set like_count = like_count+1 where id = ?", articleID).Error
		if err != nil {
			return err
		}
		article, _ := repository.ArticleRepository.GetArticleByID(database.GetDB(), articleID)
		err = database.GetDB().Exec("update user set be_liked_count = be_liked_count+1 where id = ?", article.UserID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// 判断用户是否已经点赞过文章
func (s *lcService) IsArticleLiked(article *model.Article, user *model.User) bool {
	if user == nil {
		return false
	}
	lcStatus, _ := repository.LCRepository.GetUserLikeOperation(database.GetDB(), user.ID, article.ID)
	return lcStatus.Status == LikeArticle
}

// 判断用户是否已经收藏过文章
func (s *lcService) IsArticleFavorited(article *model.Article, user *model.User) bool {
	if user == nil {
		return false
	}
	lcStatus, _ := repository.LCRepository.GetUserFavoriteOperation(database.GetDB(), user.ID, article.ID)
	return lcStatus.Status == FavoriteArticle
}

// 取消点赞文章
func (s *lcService) PostDelLikeArticle(userID, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserLikeOperation(database.GetDB(), userID, articleID)
	if err != nil || opHis.ID == 0 {
		return errors.New("数据库查询失败")
	}

	// 已经取消点赞
	if opHis.Status == 0 {
		return nil
	}

	err = database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		err = repository.LCRepository.UpdateUserLikeOperation(database.GetDB(), userID, articleID, map[string]interface{}{"status": 0})
		if err != nil {
			return err
		}
		err = database.GetDB().Exec("update article set like_count = like_count-1 where id = ?", articleID).Error
		if err != nil {
			return err
		}
		article, _ := repository.ArticleRepository.GetArticleByID(database.GetDB(), articleID)
		err = database.GetDB().Exec("update user set be_liked_count = be_liked_count-1 where id = ?", article.UserID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// 收藏文章
func (s *lcService) PostFavoriteArticle(userID, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserFavoriteOperation(database.GetDB(), userID, articleID)
	if err != nil {
		return errors.New("数据库查询失败")
	}
	// 已经收藏
	if opHis.Status == 1 {
		return nil
	}

	err = database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		if opHis.ID != 0 {
			err = repository.LCRepository.UpdateUserFavoriteOperation(database.GetDB(), userID, articleID, map[string]interface{}{"status": 1})
			database.GetDB().Exec("update user_favorite_article set update_time = ? where id = ?", util.NowTimestamp(), opHis.ID)
		} else {
			err = repository.LCRepository.CreateFavorite(database.GetDB(), &model.UserFavoriteArticle{
				UserID:     userID,
				ArticleID:  articleID,
				Status:     1,
				CreateTime: util.NowTimestamp(),
				UpdateTime: util.NowTimestamp(), // 每次收藏或取消收藏修改此时间
			})
		}
		if err != nil {
			return err
		}
		err = database.GetDB().Exec("update user set favourite_article_count = favourite_article_count+1 where id = ?", userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// 取消收藏
func (s *lcService) PostDelFavoriteArticle(userID, articleID int64) error {
	opHis, err := repository.LCRepository.GetUserFavoriteOperation(database.GetDB(), userID, articleID)
	if err != nil || opHis.ID == 0 {
		return errors.New("数据库查询失败")
	}

	// 已经取消收藏
	if opHis.Status == 0 {
		return nil
	}

	err = database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		err = repository.LCRepository.UpdateUserFavoriteOperation(database.GetDB(), userID, articleID, map[string]interface{}{"status": 0})
		if err != nil {
			return err
		}
		err = database.GetDB().Exec("update user set favourite_article_count = favourite_article_count-1 where id = ?", userID).Error
		return err
	})
	if err != nil {
		return errors.New("数据库操作出错")
	}
	return nil
}

// 获取用户收藏的文章列表
func (s *lcService) GetUserFavoriteArticleList(user *model.User, limit int, cursorTime int64, sortby, order string) (*model.FavoriteResponse, error) {
	// 获取用户收藏记录
	records := repository.LCRepository.GetFavoriteRecords(database.GetDB(), user.ID, cursorTime, limit, sortby, order)

	articles := make([]model.Article, 0, len(records))
	for _, record := range records {
		article, _ := repository.ArticleRepository.GetArticleByID(database.GetDB(), record.ArticleID)
		articles = append(articles, *article)
	}

	briefList, _ := ArticleService.BuildArticleList(user, articles)
	minCursorTime := cursorTime
	for i := range records {
		minCursorTime = util.MinInt64(minCursorTime, records[i].UpdateTime)
	}
	resp := &model.FavoriteResponse{}
	resp.TotalNum = len(briefList)
	resp.Cursor = minCursorTime
	resp.FavoriteList = briefList
	return resp, nil
}
