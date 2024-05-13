package service

import (
	"Campus-forum-system/database"
	"Campus-forum-system/logs"
	"Campus-forum-system/middleware"
	"Campus-forum-system/model"
	"Campus-forum-system/repository"
	"Campus-forum-system/util"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

// GetCurrentUser ...
func (s *userService) GetCurrentUser(c *gin.Context) *model.User {
	token := s.GetToken(c)
	userToken, err := repository.UserTokenRepository.GetUserIDByToken(database.GetDB(), token)
	if err != nil {
		logs.Logger.Errorf("数据库查询token出错")
		return nil
	}
	if userToken == nil || userToken.Status || userToken.ExpiredAt < util.NowTimestamp() { // 不存在或者过期了
		return nil
	}
	user, err := repository.UserRepository.GetUserByUserID(database.GetDB(), userToken.UserID)
	if err != nil {
		logs.Logger.Errorf("数据库查询user出错")
		return nil
	}
	return user
}

// 从请求头中获取token
func (s *userService) GetToken(c *gin.Context) string {
	token := c.GetHeader("X-User-Token")
	return token
}

func (s *userService) SignUp(c *gin.Context) (*model.User, error) {
	req := getReqFromContext(c).(*model.RegisterRequest)

	// data verification
	if !util.CheckEmail(req.Email) {
		return nil, errors.New("邮箱格式错误")
	}
	userInfo, err := repository.UserRepository.GetUserByEmail(database.GetDB(), req.Email)
	if err != nil {
		return nil, errors.New("查询邮箱出错")
	}
	if userInfo.ID != 0 {
		return nil, errors.New("邮箱已被占用")
	}
	if !util.CheckUsername(req.Username) {
		return nil, errors.New("用户名格式错误")
	}
	userInfo, err = repository.UserRepository.GetUserByUsername(database.GetDB(), req.Username)
	if err != nil {
		return nil, errors.New("查询用户名出错")
	}
	if userInfo.ID != 0 {
		return nil, errors.New("用户名已被占用")
	}
	if !util.CheckPassword(req.Password) {
		return nil, errors.New("密码格式错误")
	}

	// 要哈希的密码（以字节数组形式）和一个表示密码强度的整数值（称为“cost”）。它返回一个字节数组，其中包含密码的 bcrypt 哈希值。这个哈希值通常是一个长字符串，可以保存在数据库中用于后续验证用户密码。
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) // 加密之后的密码

	// get avatar url
	var avatarURL string
	if req.Email != "" {
		avatarURL = util.RandomAvatarURL(req.Email)
	} else {
		avatarURL = util.RandomAvatarURL(req.Username)
	}

	user := &model.User{
		Username:   req.Username,
		Password:   string(encryptedPassword),
		Email:      req.Email,
		Nickname:   req.Username,
		AvatarURL:  avatarURL,
		CreateTime: util.NowTimestamp(),
	}

	if err := repository.UserRepository.Create(database.GetDB(), user); err != nil {
		logs.Logger.Error("db error:", err)
		return nil, errors.New("数据库操作出错")
	}
	return user, nil
	// util.DB().Transaction(func(tx *gorm.DB) error {

	// 	return nil
	// })
}

func (s *userService) Login(c *gin.Context) (user *model.User, err error) {
	req := getReqFromContext(c).(*model.LoginRequest)
	if req.Email == "" && req.Username == "" || req.Email != "" && req.Username != "" {
		return nil, errors.New("请使用用户名或邮箱二者之一登录")
	}
	if req.Email != "" {
		return s.loginByEmail(req.Email, req.Password)
	}
	return s.loginByUsername(req.Username, req.Password)
}

func (s *userService) loginByEmail(email string, password string) (*model.User, error) {
	user, err := repository.UserRepository.GetUserByEmail(database.GetDB(), email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) loginByUsername(username string, password string) (*model.User, error) {
	user, err := repository.UserRepository.GetUserByUsername(database.GetDB(), username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) Logout(c *gin.Context) error {
	token := s.GetToken(c)
	return repository.UserTokenRepository.UpdateStatusInvalidByToken(database.GetDB(), token)
}

func (s *userService) SetToken(userID int64) string {
	token, err := middleware.ReleaseToken(uint(userID))
	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(model.TokenExpireDays))
	if err != nil {
		logs.Logger.Error("生成token失败")
	}
	userToken := &model.UserToken{
		UserID:     userID,
		Token:      token,
		ExpiredAt:  util.Timestamp(expireTime),
		CreateTime: util.NowTimestamp(),
	}
	repository.UserTokenRepository.Create(database.GetDB(), userToken)
	return token
}

func (s *userService) GetUserInfo(userID int64) (*model.UserBriefInfo, error) {
	user, err := repository.UserRepository.GetUserByUserID(database.GetDB(), userID)
	if err != nil {
		return nil, errors.New("此用户不存在")
	}
	briefInfo := BuildUserBriefInfo(user)
	if briefInfo == nil {
		return nil, errors.New("此用户不存在")
	}
	return briefInfo, nil
}

func (s *userService) UpdateUserProfile(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.UpdateUserProfile)
	if user.ID != req.UserID {
		return errors.New("非当前登录用户")
	}
	mp := map[string]interface{}{
		"nickname":    req.Nickname,
		"description": req.Description,
	}
	err := repository.UserRepository.UpdateMulti(database.GetDB(), user.ID, mp)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetUsername(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetUsernameRequest)
	if !util.CheckUsername(req.Username) {
		return errors.New("用户名不合法！")
	}
	err := repository.UserRepository.UpdateOne(database.GetDB(), user.ID, "username", req.Username)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetEmail(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetEmailRequest)
	if !util.CheckEmail(req.Email) {
		return errors.New("邮箱不合法！")
	}
	err := repository.UserRepository.UpdateOne(database.GetDB(), user.ID, "email", req.Email)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) SetPassword(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.SetPasswordRequest)
	if !util.CheckPassword(req.Password) {
		return errors.New("密码不合法！")
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err := repository.UserRepository.UpdateOne(database.GetDB(), user.ID, "password", encryptedPassword)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
		return errors.New("操作失败")
	}
	return nil
}

func (s *userService) UpdatePassword(c *gin.Context) error {
	user := s.GetCurrentUser(c)
	if user == nil {
		return errors.New("当前未登录！")
	}
	req := getReqFromContext(c).(*model.UpdatePasswordRequest)
	if !util.CheckPassword(req.Password) {
		return errors.New("密码不合法！")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("原密码输入错误")
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err := repository.UserRepository.UpdateOne(database.GetDB(), user.ID, "password", encryptedPassword)
	if err != nil {
		logs.Logger.Errorf("数据库操作出错:%+v", err)
	}
	return err
}

// BuildUserBriefInfo ...
func BuildUserBriefInfo(user *model.User) *model.UserBriefInfo {
	if user == nil {
		return nil
	}
	userInfo := &model.UserBriefInfo{
		ID:                    user.ID,
		Username:              user.Username,
		Nickname:              user.Nickname,
		AvatarURL:             user.AvatarURL,
		Gender:                user.Gender,
		Description:           user.Description,
		AttentionCount:        user.AttentionCount,
		FavouriteArticleCount: user.FavouriteArticleCount,
		FansCount:             user.FansCount,
		PostCount:             user.PostCount,
		CommentCount:          user.CommentCount,
		Type:                  user.Type,
		City:                  user.City,
		Province:              user.Province,
		Country:               user.Country,
	}
	return userInfo
}
