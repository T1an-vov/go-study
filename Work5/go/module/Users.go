package module

type User struct {
	ID        uint
	Username  string `form:"username" json:"username" binding:"required" gorm:"unique"`
	Password  string `form:"password" json:"password" binding:"required" `
	Credit    int    `form:"credit" json:"credit" ` // 目前学分
	Maxcredit int    `form:"maxcredit" json:"maxcredit" binding:"required"` //最大学分
}

