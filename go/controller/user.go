package controller

import (
	"Campus-forum-system/database"
	"Campus-forum-system/middleware"
	"Campus-forum-system/model"
	"Campus-forum-system/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	DB := database.GetDB()

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// fmt.Println(username, "--------", password)

	// 判断密码合法性
	if len(password) < 6 {
		response.Response(ctx, 403, 403, "密码长度太低", "密码至少需要6位")
		return
	}
	var user model.User
	DB.Table("users").Where("username = ?", username).First(&user)
	if user.ID == 0 {
		response.Response(ctx, 403, 403, "用户不存在", "请重新输入信息")
		return
	}
	// 判断用户登录状态
	// if user.State {
	// 	response.Response(ctx, 403, 403, "用户在其他地方已经登录", "请勿重复登录")
	// 	return
	// } else {
	// 	DB.Model(&user).Update("State", true)
	// }

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(ctx, 403, 403, "密码错误", "请重新输入密码")
		return
	}
	// 发token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, 500, 500, "token发放错误", "token发放错误")
		return
	}

	response.Response(ctx, 200, 200, token, "登录成功")
}

func Regist(ctx *gin.Context) {
	DB := database.GetDB()

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// fmt.Println(username, "--------", password)
	// 判断密码合法性
	if len(password) < 6 {
		response.Response(ctx, 403, 403, "密码长度太低", "密码至少需要6位")
		return
	}
	var user model.User
	DB.Table("users").Where("username = ?", username).First(&user)
	if user.ID != 0 {
		response.Response(ctx, 403, 403, "用户已经存在", "请更换用户名")
		return
	}
	// 开始注册流程
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //生成哈希密码 GenerateFromPassword 不接受超过 72 字节的密码，这是 bcrypt 操作的最长密码。
	if err != nil {
		response.Response(ctx, 403, 403, "密码加密错误", "可能是密码超过72字节")
		return
	}
	newUser := model.User{
		Username: username,
		Password: string(hashPassword),
	}
	DB.Create(&newUser)

	////发放token
	//token, err := middleware.ReleaseToken(user)
	//if err != nil {
	//	response.Response(ctx, 403, 403, "发放token失败", "发放token失败")
	//	return
	//}

	response.Response(ctx, 200, 200, "注册成功", "注册成功")

}

// func LogOut(ctx *gin.Context) {
// 	user, err := ctx.Get("user")
// 	token, err1 := ctx.Get("token")
// 	// err == false 证明value不存在
// 	if !err && !err1 {
// 		response.Response(ctx, 500, 500, "ctx获取上下文错误", "服务器错误")
// 		fmt.Println("ctx获取上下文错误：", err)
// 		return
// 	}
// 	tokenString := fmt.Sprintf("%v", token)

// 	// 将token放入redis黑名单
// 	database.RedisSetKey(tokenString, tokenString)

// 	// 将user转换成model.User类型
// 	u := user.(model.User)
// 	DB := database.GetDB()
// 	DB.Table("users").Where("id = ?", u.ID).First(&u)
// 	u.State = false
// 	DB.Save(&u)

// 	response.Response(ctx, 200, 200, "用户登出成功", "登出成功")

// }

// func GetInfo(ctx *gin.Context) {
// 	user, _ := ctx.Get("user")
// 	token, _ := ctx.Get("token")
// 	if token == nil {
// 		ctx.JSON(500, gin.H{
// 			"msg": "ctx获取上下文错误",
// 		})
// 		fmt.Println("GetInfo():ctx获取上下文错误")
// 		return
// 	}
// 	tokenString := fmt.Sprintf("%v", token)

// 	if user == nil {
// 		ctx.JSON(500, gin.H{
// 			"msg": "ctx获取上下文错误",
// 		})
// 		fmt.Println("GetInfo()的user:ctx获取上下文错误")
// 		return
// 	}
// 	// 将user转换成model.User类型
// 	u := user.(model.User)
// 	ctx.JSON(200, gin.H{
// 		"UserId":   u.ID,
// 		"username": u.Username,
// 		"token":    tokenString,
// 	})
// }
