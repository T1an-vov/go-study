package module

type Record struct {
	Id uint
	Number string `json:"number" form:"number" `//唯一房间号
	Email string `json:"email" form:"email" binding:"required"`//
}
