package service

import (
	"Campus-forum-system/database"
	"Campus-forum-system/model"
	"Campus-forum-system/repository"
	"Campus-forum-system/util"
	"errors"
	"unicode/utf8"

	"gorm.io/gorm"
)

type articleService struct {
}

func newArticleService() *articleService {
	return &articleService{}
}

var ArticleService = newArticleService()

// 上传文章
func (s *articleService) PostArticle(user *model.User, title string, content string) (*model.Article, error) {
	article := &model.Article{
		UserID:     user.ID,
		Title:      title,
		Content:    content,
		CreateTime: util.NowTimestamp(),
	}

	err := database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error

		err = repository.ArticleRepository.Create(database.GetDB(), article)
		if err != nil {
			return err
		}
		err = database.GetDB().Exec("update users set post_count = post_count+1 where id = ?", user.ID).Error
		return err
	})
	if err != nil {
		return nil, errors.New("post article failed")
	}
	return article, nil
}

// 获取文章列表
func (s *articleService) GetArticleList(currentUser *model.User, authorID int64, limit int, cursorTime int64, sortby string, order string) (*model.ArticleListResponse, error) {
	resp := &model.ArticleListResponse{}
	fields := []string{"id", "title", "create_time", "user_id", "view_count", "comment_count", "like_count", "content"}
	articles := repository.ArticleRepository.GetArticleFields(database.GetDB(), authorID, fields, cursorTime, limit, sortby, order)

	briefList, minCursorTime := s.BuildArticleList(currentUser, articles)

	resp.Cursor = minCursorTime
	for i := range briefList {
		// 最小创建时间
		if briefList[i].CreateTime < resp.Cursor {
			resp.Cursor = briefList[i].CreateTime
		}
	}

	resp.ArticleList = briefList
	resp.TotalNum = len(briefList)
	return resp, nil
}

func (s *articleService) GetArticleByID(currentUser *model.User, articleID int64) (*model.ArticleResponse, error) {
	articleInfo, err := repository.ArticleRepository.GetArticleByID(database.GetDB(), articleID)
	if err != nil {
		return nil, errors.New("get article by id failed")
	}
	userInfo, err := repository.UserRepository.GetUserByUserID(database.GetDB(), articleInfo.UserID)
	if err != nil {
		return nil, errors.New("get user by user id failed")
	}

	resp := &model.ArticleResponse{
		ArticleID:    articleInfo.ID,
		Title:        articleInfo.Title,
		User:         BuildUserBriefInfo(userInfo),
		Content:      util.MarkdownToHTML(articleInfo.Content),
		Liked:        LCService.IsArticleLiked(articleInfo, currentUser),
		Favortied:    LCService.IsArticleFavorited(articleInfo, currentUser),
		CommentCount: articleInfo.CommentCount,
		LikeCount:    articleInfo.LikeCount,
		CreateTime:   articleInfo.CreateTime,
	}
	return resp, nil
}

// 构建文章简要信息列表
func (s *articleService) BuildArticleList(currentUser *model.User, articles []model.Article) ([]*model.ArticleBriefInfo, int64) {
	var minCursorTime int64 = model.MAXCursorTime
	briefList := make([]*model.ArticleBriefInfo, len(articles))
	for i := range articles {
		minCursorTime = util.MinInt64(minCursorTime, articles[i].CreateTime) // 找到最小的创建时间
		// utf8.RuneCountInString(s string) (n int) 返回字符串的runes数量
		mkSummary := util.MarkdownToHTML(util.SubString(articles[i].Content, 0, util.MinInt(128, utf8.RuneCountInString(articles[i].Content))))
		briefList[i] = new(model.ArticleBriefInfo)                                                 // 为briefList[i]分配空间
		briefList[i].ArticleID = articles[i].ID                                                    // 文章id
		briefList[i].Title = articles[i].Title                                                     // 标题
		briefList[i].Summary = util.GetHTMLText(mkSummary)                                         // 简要信息
		briefList[i].CommentCount = articles[i].CommentCount                                       // 评论数
		briefList[i].LikeCount = articles[i].LikeCount                                             // 点赞数
		briefList[i].ViewCount = articles[i].ViewCount                                             // 浏览数
		briefList[i].CreateTime = articles[i].CreateTime                                           // 创建时间
		briefList[i].Liked = LCService.IsArticleLiked(&articles[i], currentUser)                   // 当前登录用户是否已点赞(如未登录为false)
		user, _ := repository.UserRepository.GetUserByUserID(database.GetDB(), articles[i].UserID) // 获取当前用户
		briefList[i].User = BuildUserBriefInfo(user)                                               // 作者信息
	}

	return briefList, minCursorTime
}
