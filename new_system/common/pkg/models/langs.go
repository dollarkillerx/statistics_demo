package models

import "gorm.io/gorm"

type LangPhrase struct {
	BasicModel
	Type        string `gorm:"type:varchar(100)" json:"type"` // 新增的类型字段
	LanguageKey string `gorm:"type:varchar(255);column:language_key" json:"language_key"`
	ZhCN        string `gorm:"type:varchar(255);column:zh_cn" json:"zh_cn"`   // 中文简体
	ZhHk        string `gorm:"type:varchar(255);column:zh_hk" json:"zh_hk"`   // 中文繁体
	EnUs        string `gorm:"type:varchar(255);column:en_us" json:"en_us"`   // 英文
	KoKr        string `gorm:"type:varchar(255);column:ko_kr" json:"ko_kr"`   // 韩文
	JaJp        string `gorm:"type:varchar(255);column:ja_jp" json:"ja_jp"`   // 日文
	Module      string `gorm:"type:varchar(255);column:module" json:"module"` // 模块
}

func (m *LangPhrase) GetByModular(modular string, db *gorm.DB) []LangPhrase {
	var list []LangPhrase
	db.Model(&LangPhrase{}).Where("module = ?", modular).Find(&list)
	return list
}
