package models

import (
	"gorm.io/gorm"
	"strings"
)

type TeamLevelQueue struct {
	BasicModel
	UserID     int64  `gorm:"column:user_id" json:"user_id"`
	ChangeTree string `gorm:"type:varchar(600);column:change_tree" json:"change_tree"`
	DealStatus string `gorm:"type:varchar(255);column:deal_status" json:"deal_status"` // 0 未处理 1 已处理 ?
}

func (t *TeamLevelQueue) TableName() string {
	return "product_team_level_queue"
}

// AddQueue 加入队列
func (t *TeamLevelQueue) AddQueue(db *gorm.DB, userId int64) error {
	us := new(User)
	user, err := us.GetUserById(db, userId)
	if err != nil {
		return err
	}

	// 查询是否有存在 当前userId 的任务
	var inQueue int64
	err = db.Model(&TeamLevelQueue{}).Where("FIND_IN_SET(?,change_tree)", userId).Count(&inQueue).Error
	if err != nil {
		return err
	}

	if inQueue == 0 {
		// 拼接 changeTree
		changeTree := user.ParentTree + "," + string(user.ID)
		changeTree = strings.Trim(changeTree, ",")
		queue := TeamLevelQueue{
			UserID:     userId,
			ChangeTree: changeTree,
			DealStatus: "0",
		}

		err = db.Create(&queue).Error
		if err != nil {
			return err
		}
	}

	return nil
}
