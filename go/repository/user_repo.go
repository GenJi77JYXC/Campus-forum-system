package repository

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct{}

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return new(userRepository)
}

// Create ...
func (r *userRepository) Create(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	return r.take(db, "email = ?", email)
}

func (r *userRepository) GetUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	return r.take(db, "username = ?", username)
}

func (r *userRepository) GetUserByUserID(db *gorm.DB, userID int64) (*model.User, error) {
	return r.take(db, "id = ?", userID)
}

func (r *userRepository) UpdateOne(db *gorm.DB, userID int64, column string, value interface{}) error {
	return db.Model(model.User{}).Where("id = ?", userID).Update(column, value).Error
}

func (r *userRepository) UpdateMulti(db *gorm.DB, userID int64, kv map[string]interface{}) error {
	return db.Model(model.User{}).Where("id = ?", userID).Updates(kv).Error
}

func (r *userRepository) take(db *gorm.DB, column string, value interface{}) (*model.User, error) {
	result := &model.User{}
	// err := db.Where(column, value).Take(result).Error
	err := db.Where(column, value).Find(&result).Error
	if err != nil {
		logs.Logger.Errorf("query db error:", err)
		return nil, errors.New("用户不存在")
	}

	// userId := result.ID
	// token := &model.UserToken{}

	// err = db.Where("user_id = ?", userId).Last(&token).Error
	// if err != nil {
	// 	// 当用户没有登录过时，user_token表中没有记录, 查询时会报错record not found，这里需要捕获这个错误, 表示用户没有登录过，可以正常登录
	// 	if err.Error() == "record not found" {
	// 		return result, nil
	// 	}
	// 	logs.Logger.Errorf("query db error:", err)
	// 	return nil, err
	// }

	// if !token.Status {
	// 	// fmt.Println(result.Username)
	// 	logs.Logger.Errorf(result.Username, ":该用户在别处已经登录")

	// 	return nil, errors.New("该用户在别处已经登录")
	// }

	return result, nil
}
