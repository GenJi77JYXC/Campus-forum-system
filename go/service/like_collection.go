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
