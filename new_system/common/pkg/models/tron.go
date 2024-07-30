package models

import (
	"gorm.io/gorm"
)

// TRON 相关

// TokenChain 区块链支付钱包相关
type TokenChain struct {
	BasicModel
	TokenID      int64  `gorm:"column:token_id" json:"token_id"`                         // token 表对应 id
	Contract     string `gorm:"type:varchar(255);column:contract" json:"contract"`       // 合约地址
	PayAddress   string `gorm:"type:varchar(255);column:pay_address" json:"pay_address"` // 支付账户
	PayPassword  string `gorm:"type:varchar(255);column:pay_password" json:"pay_password"`
	UnionAddress string `gorm:"type:varchar(255);column:union_address" json:"union_address"` // 归集账户
	GasLimit     uint64 `gorm:"column:gas_limit" json:"gas_limit"`                           // gas
	Status       uint   `gorm:"column:status" json:"status"`                                 // 1 启用 0 弃用 ?
}

func (t *TokenChain) TableName() string {
	return "token_chain"
}

// GetAll 获取所有的 TokenChain
func (t *TokenChain) GetAll(db *gorm.DB) []TokenChain {
	var list []TokenChain
	db.Model(&TokenChain{}).Where("status = 1").Find(&list)
	return list
}

// GetAdminToken 获取管理员账户
func (t *TokenChain) GetAdminToken(db *gorm.DB) (TokenChain, error) {
	var chain TokenChain
	err := db.Model(&TokenChain{}).Where("token_id = ?", 1).First(&chain).Error
	return chain, err
}

// Token 代币
type Token struct {
	BasicModel
	Name        string  `gorm:"type:varchar(255);column:name" json:"name"`                 // 代币名
	Logo        string  `gorm:"type:varchar(255);column:logo" json:"logo"`                 // 代币logo
	EqUSD       float64 `gorm:"type:decimal(18,6);column:eq_usd" json:"eq_usd"`            // 对美元价格
	PlatformKey string  `gorm:"type:varchar(255);column:platform_key" json:"platform_key"` // token: USDT or WUSDT
	QuoteSort   int     `gorm:"column:quote_sort" json:"quote_sort"`                       // 排序
	Status      int8    `gorm:"column:status" json:"status"`                               // 状态 请提供枚举定义!!!
}

func (t *Token) TableName() string {
	return "token"
}

// GetTokenByKey 根据 key 获取 token
func (t *Token) GetTokenByKey(db *gorm.DB, key string) (*Token, error) {
	var token Token
	err := db.Model(&Token{}).Where("platform_key = ?", key).First(&token).Error
	return &token, err
}

// ChainSend 提币记录
type ChainSend struct {
	BasicModel
	WithdrawID int64   `gorm:"column:withdraw_id" json:"withdraw_id"`                 // 提币 ID
	TokenID    int64   `gorm:"column:token_id" json:"token_id"`                       // token 表对应 id
	TokenName  string  `gorm:"type:varchar(255);column:token_name" json:"token_name"` // 代币名
	Amount     float64 `gorm:"type:decimal(18,6);column:amount" json:"amount"`        // 数量
	TxHash     string  `gorm:"type:varchar(255);column:tx_hash" json:"tx_hash"`       // 交易哈希
	Remark     string  `gorm:"type:varchar(255);column:remark" json:"remark"`         // 备注
}

func (c *ChainSend) TableName() string {
	return "chain_send"
}

// GetByWithdrawID 根据提币 ID 获取提币记录
func (c *ChainSend) GetByWithdrawID(db *gorm.DB, withdrawID int64) (*ChainSend, error) {
	var chainSend ChainSend
	err := db.Model(&ChainSend{}).Where("withdraw_id = ?", withdrawID).First(&chainSend).Error
	return &chainSend, err
}

func (c *ChainSend) AddChainSend(db *gorm.DB, chainSend *ChainSend) error {
	return db.Model(&ChainSend{}).Create(chainSend).Error
}
