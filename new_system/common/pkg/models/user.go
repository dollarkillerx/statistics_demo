package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	BasicModel
	Mobile       string    `gorm:"type:varchar(100);column:mobile" json:"mobile"`             // 手机号
	Nickname     string    `gorm:"type:varchar(100);column:nickname" json:"nickname"`         // 昵称
	Avatar       string    `gorm:"type:varchar(255);column:avatar" json:"avatar"`             // 		头像
	Password     string    `gorm:"type:varchar(255);column:password" json:"password"`         // 密码
	PayPassword  string    `gorm:"type:varchar(255);column:pay_password" json:"pay_password"` // 支付密码
	Address      string    `gorm:"type:varchar(255);column:address" json:"address"`           // 地址 什么地址?
	ParentID     int       `gorm:"column:parent_id" json:"parent_id"`                         // 父级ID
	ParentTree   string    `gorm:"type:varchar(255);column:parent_tree" json:"parent_tree"`   // 父级树
	Level        int       `gorm:"column:level" json:"level"`                                 // 当前邀请层级
	LastIP       string    `gorm:"type:varchar(45);column:last_ip" json:"last_ip"`            // 最后登录IP
	InviteCode   string    `gorm:"type:varchar(100);column:invite_code" json:"invite_code"`   // 邀请码
	AuthStatus   int8      `gorm:"column:auth_status" json:"auth_status"`                     // 认证状态? 什么认证状态 请给出枚举
	ActiveStatus int8      `gorm:"column:active_status" json:"active_status"`                 // 活跃状态? 请给出枚举
	LeaderID     int64     `gorm:"column:leader_id" json:"leader_id"`                         // 顶级 tree 的id?
	InvitedAt    time.Time `gorm:"column:invited_at;autoCreateTime" json:"invited_at"`        // 被邀请时间
}

func (m *User) TableName() string {
	return "user"
}

func (m *User) GetUserById(db *gorm.DB, id int64) (*User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return &user, err
}
