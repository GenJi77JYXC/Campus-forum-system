package model

var modelList = []interface{}{
	&User{},
	&UserToken{},
	&Article{},
	&Comment{},
	&UserLikeArticle{},
	&UserFavoriteArticle{},
	&UserLikeComment{},
}

// Model ...
type Model struct {
	ID int64 `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id"`
}

// gorm tag 说明：
// column: 字段名 （mysql中字段名）
// type: 字段类型
// primaryKey: 主键
// autoIncrement: 自增
// unique: 唯一
// not null: 非空
// json: 序列化字段名
// string类型会自动转为varchar类型, 可以通过size指定varchar的长度
type User struct {
	Model
	Username              string `gorm:"unique;type:varchar(20);not null" json:"username"`
	Nickname              string `gorm:"column:nickname;type:varchar(30)" json:"nickname"`
	Password              string `gorm:"column:password;size:512;not null" json:"password"`
	AvatarURL             string `gorm:"column:avatar_url;size:255" json:"avatar_url"`                                     // 头像
	Gender                string `gorm:"column:gender;type:tinyint;default:2" json:"gender"`                               // 性别
	Email                 string `gorm:"column:email;type:varchar(50)" json:"email"`                                       // 邮箱
	EmailVerified         bool   `gorm:"column:email_verified;type:tinyint;default 0" json:"email_verified"`               // 邮箱是否已验证
	Description           string `gorm:"column:description;type:varchar(200)" json:"description"`                          // 个人描述
	AttentionCount        int    `gorm:"column:attention_count;type:int;default:0" json:"attention_count"`                 // 关注数
	FavouriteArticleCount int    `gorm:"column:favourite_article_count;type:int;default:0" json:"favourite_article_count"` // 收藏的文章数
	FansCount             int    `gorm:"column:fans_count;type:int;default:0" json:"fans_count"`                           // 粉丝数
	PostCount             int    `gorm:"column:post_count;type:int;default:0" json:"post_count"`                           // 发表数
	CommentCount          int    `gorm:"column:comment_count;type:int;default:0" json:"comment_count"`                     // 评论数
	BeLikedCount          int    `gorm:"column:be_liked_count;type:int;default:0" json:"be_liked_count"`                   // 获赞数
	Type                  int    `gorm:"column:type;type:int;default:0" json:"type"`                                       // 用户类型
	City                  string `gorm:"column:city;type:varchar(50)" json:"city"`                                         // 城市
	Province              string `gorm:"column:province;type:varchar(50)" json:"province"`                                 // 省份
	Country               string `gorm:"column:country;type:varchar(50)" json:"country"`                                   // 国家
	CreateTime            int64  `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime            int64  `gorm:"column:update_time;default:null" json:"update_time"`
	DeleteTime            int64  `gorm:"column:delete_time;default:null" json:"delete_time"`
}

// UserToken 用户token
type UserToken struct {
	Model
	UserID     int64  `gorm:"column:user_id;type:int" json:"user_id"`
	Token      string `gorm:"size:255;unique;not null" json:"token"`              //token
	ExpiredAt  int64  `gorm:"column:expired_at;type:int" json:"expired_at"`       //有效期至
	Status     bool   `gorm:"column:status;type:tinyint;default 0" json:"status"` // 0有效 1失效
	CreateTime int64  `gorm:"column:create_time;default:null" json:"create_time"`
}

// Article 文章
type Article struct {
	Model
	UserID       int64  `gorm:"column:user_id;type:int" json:"user_id"`                       //作者ID
	Title        string `gorm:"column:title;type:varchar(50);not null" json:"title"`          //标题
	Status       int    `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`  //文章状态  可以用于管理员审核
	Content      string `gorm:"column:content;type:text" json:"content"`                      //内容
	ViewCount    int    `gorm:"column:view_count;type:int;default:0" json:"view_count"`       //浏览数
	CommentCount int    `gorm:"column:comment_count;type:int;default:0" json:"comment_count"` //评论数
	LikeCount    int    `gorm:"column:like_count;type:int;default:0" json:"like_count"`       //点赞数
	CreateTime   int64  `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime   int64  `gorm:"column:update_time;default:null" json:"update_time"`
	DeleteTime   int64  `gorm:"column:delete_time;default:null" json:"delete_time"`
}

// Comment 评论
type Comment struct {
	Model
	UserID     int64  `gorm:"column:user_id;type:int" json:"user_id"`       //作者ID
	ArticleID  int64  `gorm:"column:article_id;type:int" json:"article_id"` //所属的文章ID
	Content    string `gorm:"column:content;type:text" json:"content"`
	ParentID   int64  `gorm:"column:parent_id;type:int" json:"parent_id"`                  //父评论ID（引用的评论）
	Status     int    `gorm:"column:status;type:tinyint;not null;default:0" json:"status"` //评论状态
	LikeCount  int    `gorm:"column:like_count;type:int;default:0" json:"like_count"`      //点赞数
	CreateTime int64  `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time;default:null" json:"update_time"`
	DeleteTime int64  `gorm:"column:delete_time;default:null" json:"delete_time"`
}

// UserLikeArticle 用户点赞文章
type UserLikeArticle struct {
	Model
	UserID     int64 `gorm:"column:user_id;type:int" json:"user_id"`
	ArticleID  int64 `gorm:"column:article_id;type:int" json:"article_id"`
	Status     int   `gorm:"column:status;type:tinyint;not null;default:0" json:"status"` //是否点赞 1已点赞，0未点赞
	UpdateTime int64 `gorm:"column:update_time;default:null" json:"update_time"`
}

// UserLikeComment 用户点赞评论
type UserLikeComment struct {
	Model
	UserID     int64 `gorm:"column:user_id;type:int" json:"user_id"`
	CommentID  int64 `gorm:"column:comment_id;type:int" json:"comment_id"`
	Status     int   `gorm:"column:status;type:tinyint;not null;default:0" json:"status"` //是否点赞 1已点赞，0未点赞
	UpdateTime int64 `gorm:"column:update_time;default:null" json:"update_time"`
}

// UserFavoriteArticle 用户收藏文章
type UserFavoriteArticle struct {
	Model
	UserID     int64 `gorm:"column:user_id;type:int" json:"user_id"`
	ArticleID  int64 `gorm:"column:article_id;type:int" json:"article_id"`
	Status     int   `gorm:"column:status;type:tinyint;not null;default:0" json:"status"` //是否已收藏 1已收藏，0未收藏
	CreateTime int64 `gorm:"column:create_time;default:null" json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time;default:null" json:"update_time"`
}

func GetModelList() []interface{} {
	return modelList
}
