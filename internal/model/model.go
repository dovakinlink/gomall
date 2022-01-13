package model

type Member struct {
	Id         int    `gorm:"id" json:"id"`
	Username   string `gorm:"username" json:"username"` // 用户名（唯一标识）
	Password   string `gorm:"password" json:"password"` // 密码
	Nickname   string `gorm:"nickname" json:"nickname"` // 昵称
	Phone      string `gorm:"phone" json:"phone"`       // 手机号
	Status     int    `gorm:"status" json:"status"`     // 账号状态：0-启用 1-禁用
	CreateTime string `gorm:"create_time" json:"create_time"`
	UpdateTime string `gorm:"update_time" json:"update_time"`
}

func (*Member) TableName() string {
	return "member"
}
