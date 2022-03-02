package module

type Room struct {
	ID uint
	Number string `json:"number" form:"number"  gorm:"unique"`//唯一房间号
	Name string `json:"name" form:"name" binding:"required"`//房间名
	Owner string `json:"owner" form:"owner" binding:"required"`//房主()
	People uint `json:"people" form:"people"`//现有人数
	Max uint `json:"max" form:"max" binding:"required"`//最大人数
	CreateAt int64 `json:"create_at" form:"create_at"`//创建时间
	Access string `json:"access" form:"access" binding:"required"`//允许进入
}
