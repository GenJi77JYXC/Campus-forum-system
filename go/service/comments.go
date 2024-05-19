package service

import (
	"Campus-forum-system/database"
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"Campus-forum-system/repository"
	"Campus-forum-system/util"
	"errors"
	"sort"

	"gorm.io/gorm"
)

type commentService struct {
}

var CommentService = newCommentService()

func newCommentService() *commentService {
	return &commentService{}
}

func (s *commentService) BuildComment(userID, articleID, parentID int64, content string) (*model.CommentInfo, error) {
	comment := &model.Comment{
		UserID:     userID,
		ArticleID:  articleID,
		ParentID:   parentID,
		Content:    content,
		CreateTime: util.NowTimestamp(),
	}

	err := database.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		err = repository.CommentRepository.Create(database.GetDB(), comment)
		if err != nil {
			return err
		}
		// update article comment count
		err = database.GetDB().Exec("update articles set comment_count = comment_count + 1 where id = ? and user_id = ?", articleID, userID).Error
		if err != nil {
			return err
		}
		// update user comment count
		err = database.GetDB().Exec("update users set comment_count = comment_count + 1 where id = ?", userID).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, errors.New("sql error : failed to build comment")
	}
	return buildCommentInfo(comment), nil
}

func (s *commentService) GetCommentList(articleID int64, cursorTime int64) (*model.CommentListResponse, error) {
	resp := new(model.CommentListResponse)
	comtList, err := repository.CommentRepository.GetCommentsByCursorTime(database.GetDB(), articleID, cursorTime)
	if err != nil {
		return nil, errors.New("查询评论信息出错")
	}

	commentList, minCursorTime := buildCommentInfoList(comtList)
	resp.ArticleID = articleID
	resp.TotalNum = len(comtList)
	resp.Cursor = minCursorTime
	resp.CommentList = commentList
	return resp, nil
}

func buildCommentInfo(comment *model.Comment) *model.CommentInfo {
	userInfo, err := repository.UserRepository.GetUserByUserID(database.GetDB(), comment.UserID)
	if err != nil {
		logs.Logger.Errorf("查询作者信息出错")
	}
	commentInfo := &model.CommentInfo{
		CommentID:      comment.ID,
		AuthorNickName: userInfo.Nickname,
		AuthorUserName: userInfo.Username,
		AuthorID:       userInfo.ID,
		AvatarURL:      userInfo.AvatarURL,
		Content:        util.MarkdownToHTML(comment.Content),
		LikeCount:      comment.LikeCount,
		CreateTime:     comment.CreateTime,
	}
	if comment.ParentID > 0 {
		parentComment, err := repository.CommentRepository.GetCommentsByCommentID(database.GetDB(), comment.ParentID)
		if err != nil {
			logs.Logger.Errorf("查询父评论信息出错")
		}
		commentInfo.ParentComment = buildCommentInfo(parentComment)
	}
	return commentInfo
}

func buildCommentInfoList(comtList []model.Comment) ([]*model.CommentInfo, int64) {
	var minCursorTime int64 = model.MAXCursorTime

	sortComments(comtList, func(p, q *model.Comment) bool {
		return p.ID < q.ID // 按照评论id排序
	})
	detailedCommentList := make([]*model.CommentInfo, len(comtList))
	for i := range comtList {
		minCursorTime = util.MinInt64(minCursorTime, comtList[i].CreateTime)
		userInfo, err := repository.UserRepository.GetUserByUserID(database.GetDB(), comtList[i].UserID)
		if err != nil {
			logs.Logger.Errorf("查询作者信息出错")
		}
		detailedCommentList[i] = &model.CommentInfo{
			CommentID:      comtList[i].ID,
			AuthorNickName: userInfo.Nickname,
			AuthorUserName: userInfo.Username,
			AuthorID:       userInfo.ID,
			AvatarURL:      userInfo.AvatarURL,
			Content:        util.MarkdownToHTML(comtList[i].Content),
			LikeCount:      comtList[i].LikeCount,
			CreateTime:     comtList[i].CreateTime,
		}
		detailedCommentList[i].ParentComment = findParentComment(i, comtList[i].ParentID, detailedCommentList)
	}
	return detailedCommentList, minCursorTime
}

func findParentComment(len int, parentID int64, detailedCommentList []*model.CommentInfo) *model.CommentInfo {
	var l, r int = 0, len
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if detailedCommentList[mid].CommentID == parentID {
			return detailedCommentList[mid]
		}
		if detailedCommentList[mid].CommentID > parentID {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return nil
}

// sort comments
type commentWrapper struct {
	comments []model.Comment
	by       func(p, q *model.Comment) bool
}

type sortBy func(p, q *model.Comment) bool

func (pw commentWrapper) Len() int { // rewrite Len()
	return len(pw.comments)
}
func (pw commentWrapper) Swap(i, j int) { // rewrite Swap()
	pw.comments[i], pw.comments[j] = pw.comments[j], pw.comments[i]
}
func (pw commentWrapper) Less(i, j int) bool { // rewrite Less()
	return pw.by(&pw.comments[i], &pw.comments[j])
}

// sortComments
func sortComments(comments []model.Comment, by sortBy) {
	sort.Sort(commentWrapper{comments, by}) // Sort按Less方法确定的升序对数据进行排序。它会调用一次数据。Len来确定n和O(n*log(n))对data的调用。少和数据交换。排序不能保证是稳定的。

}

// 点赞评论
func (s *commentService) LikeComment(commentID, userID int64) (*model.UserLikeCommentResponse, error) {
	// 先判断是否已经点赞过
	flag, err := repository.LCRepository.IsLikeComment(database.GetDB(), commentID, userID)
	if flag || err != nil {
		return nil, errors.New("已经点赞过了")
	}

	err = repository.LCRepository.CreateCommentLike(database.GetDB(), commentID, userID)
	if err != nil {
		logs.Logger.Errorf("点赞评论失败：%v", err)
		return nil, errors.New("点赞评论失败")
	}

	return &model.UserLikeCommentResponse{
		CommentID: commentID,
		UserID:    userID,
		Status:    true,
	}, nil
}

// 取消点赞评论
func (s *commentService) UnlikeComment(commentID, userID int64) (*model.UserLikeCommentResponse, error) {
	// 先判断是否已经点赞过
	flag, err := repository.LCRepository.IsLikeComment(database.GetDB(), commentID, userID)
	if !flag || err != nil {
		return nil, errors.New("未点赞过")
	}

	err = repository.LCRepository.CancelCommentLike(database.GetDB(), commentID, userID)
	if err != nil {
		logs.Logger.Errorf("取消点赞评论失败：%v", err)
		return nil, errors.New("取消点赞评论失败")
	}

	return &model.UserLikeCommentResponse{
		CommentID: commentID,
		UserID:    userID,
		Status:    false,
	}, nil

}

// 删除评论
func (s *commentService) DeleteComment(commentID, userID int64) error {
	comment, err := repository.CommentRepository.GetCommentsByCommentID(database.GetDB(), commentID)
	if err != nil {
		logs.Logger.Errorf("查询评论信息出错")
		return errors.New("查询评论信息出错")
	}
	if comment.UserID != userID {
		return errors.New("只能删除自己的评论")
	}
	err = database.GetDB().Transaction(func(tx *gorm.DB) error {
		// 删除评论
		err = repository.CommentRepository.DeleteCommentByID(tx, commentID)
		if err != nil {
			return err
		}
		// 更新文章评论数
		err = tx.Exec("update articles set comment_count = comment_count - 1 where id = ? and user_id = ?", comment.ArticleID, comment.UserID).Error
		if err != nil {
			return err
		}
		// 更新用户评论数
		err = tx.Exec("update users set comment_count = comment_count - 1 where id = ?", comment.UserID).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		logs.Logger.Errorf("删除评论失败：%v", err)
		return errors.New("删除评论失败")
	}
	return nil
}
