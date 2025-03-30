package account

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Phone    string `gorm:"type:varchar(32);unique;default:null;" json:"phone"` // 手机号,次登录方式
	Email    string `gorm:"type:varchar(32);unique;default:null;" json:"email"` // 邮箱,主登录方式
	Password string `gorm:"type:varchar(255);not null;" json:"password"`        // 加密密码
	Nickname string `gorm:"type:varchar(64);not null;" json:"nickname"`         // 昵称
	Avatar   string `gorm:"type:varchar(255);default:null;" json:"avatar"`      //用户头像

}

func (Account) TableName() string {
	return "accounts"
}
