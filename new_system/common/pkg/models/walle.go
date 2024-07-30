package models

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
)

// WalletRecord 钱包记录
type WalletRecord struct {
	BasicModel
	UserID        int64   `gorm:"column:user_id" json:"user_id"`                                // 用户id
	WalletID      int64   `gorm:"column:wallet_id" json:"wallet_id"`                            // 钱包id
	Address       string  `gorm:"type:varchar(100);column:address" json:"address"`              // 冗余tron钱包地址
	Mobile        string  `gorm:"type:varchar(30);column:mobile" json:"mobile"`                 // 冗余
	Nickname      string  `gorm:"type:varchar(50);column:nickname" json:"nickname"`             // 冗余
	TokenID       int64   `gorm:"column:token_id" json:"token_id"`                              // token 表对应id
	TokenName     string  `gorm:"type:varchar(50);column:token_name" json:"token_name"`         // 冗余
	TokenKey      string  `gorm:"type:varchar(50);column:token_key" json:"token_key"`           // 冗余货币名称: USDT
	UsableChange  float64 `gorm:"type:decimal(18,6);column:usable_change" json:"usable_change"` // 变更值
	RtUsable      float64 `gorm:"type:decimal(18,6);column:rt_usable" json:"rt_usable"`         // 变更后的余额
	Remark        string  `gorm:"type:varchar(500);column:remark" json:"remark"`                // 备注
	Behavior      int     `gorm:"column:behavior" json:"behavior"`                              // ？ 给出 枚举 及其定义
	OrderID       int64   `gorm:"column:order_id" json:"order_id"`                              // 订单id
	TxHash        string  `gorm:"type:varchar(100);column:tx_hash" json:"tx_hash"`              // 交易hash
	FirstRecharge int8    `gorm:"column:first_recharge" json:"first_recharge"`                  // 干啥的？
}

func (w *WalletRecord) TableName() string {
	return "wallet_record"
}

// ExistWalletRecordByTxHash 根据txHash判断钱包记录是否存在
func (w *WalletRecord) ExistWalletRecordByTxHash(db *gorm.DB, txHash string) bool {
	var count int64
	db.Model(&WalletRecord{}).Where("tx_hash = ?", txHash).Count(&count)
	return count > 0
}

// GetWalletRecordTxHash 根据txHash获取钱包记录
func (w *WalletRecord) GetWalletRecordTxHash(db *gorm.DB, txHash string) (*WalletRecord, error) {
	var record WalletRecord
	err := db.Model(&WalletRecord{}).Where("tx_hash = ?", txHash).First(&record).Error
	return &record, err
}

// InsertWalletRecord 插入钱包记录
func (w *WalletRecord) InsertWalletRecord(db *gorm.DB, record *WalletRecord) error {
	return db.Create(record).Error
}

// GetWalletRecord 获取WalletRecord
func (w *WalletRecord) GetWalletRecord(db *gorm.DB, behavior int, id int64, userId int64) (*WalletRecord, error) {
	var wallet WalletRecord
	err := db.Model(&WalletRecord{}).Where("behavior = ?", behavior).
		Where("order_id = ?", id).Where("user_id = ?", userId).First(&wallet).Error
	return &wallet, err
}

// WalletAddress 用户钱包地址
type WalletAddress struct {
	BasicModel
	UserID   int64  `gorm:"column:user_id" json:"user_id"`                     // 用户id
	Address  string `gorm:"type:varchar(255);column:address" json:"address"`   // 钱包地址
	Password string `gorm:"type:varchar(255);column:password" json:"password"` // 私钥
	Remark   string `gorm:"type:varchar(255);column:remark" json:"remark"`     // 备注
}

func (w *WalletAddress) TableName() string {
	return "wallet_address"
}

// GetUserByAddress 根据地址获取用户
func (w *WalletAddress) GetUserByAddress(db *gorm.DB, address string) (*WalletAddress, error) {
	var wallet WalletAddress
	err := db.Model(&WalletAddress{}).Where("address = ?", address).First(&wallet).Error
	return &wallet, err
}

// Wallet 钱包  记录用户余额
type Wallet struct {
	BasicModel
	UserID int64   `gorm:"column:user_id" json:"user_id"`
	USDT   float64 `gorm:"type:decimal(18,6);column:usdt" json:"usdt"`     // USDT 余额
	WUSDT  float64 `gorm:"type:decimal(18,6);column:w_usdt" json:"w_usdt"` // 可 提取的usdt 余额 ?
	JYB    float64 `gorm:"type:decimal(18,6);column:jyb" json:"jyb"`       // EDC 系统内代币
}

func (w *Wallet) TableName() string {
	return "wallet"
}

// GetWalletByUserId 根据用户ID获取钱包
func (w *Wallet) GetWalletByUserId(db *gorm.DB, userId int64) (*Wallet, error) {
	var wallet Wallet
	err := db.Model(&Wallet{}).Where("user_id = ?", userId).FirstOrCreate(&wallet).Error
	return &wallet, err
}

// UpdateWalletUSDT 更新钱包余额USDT
func (w *Wallet) UpdateWalletUSDT(db *gorm.DB, userId int64, amount float64) error {
	return db.Model(&Wallet{}).Where("user_id = ?", userId).Update("usdt", gorm.Expr("usdt + ?", amount)).Error
}

// UpdateWalletUSDT 更新钱包余额WUSDT
func (w *Wallet) UpdateWalletWUSDT(db *gorm.DB, userId int64, amount float64) error {
	return db.Model(&Wallet{}).Where("user_id = ?", userId).Update("w_usdt", gorm.Expr("w_usdt + ?", amount)).Error
}

// USDTDeposit USDT充币
func (w *Wallet) USDTDeposit(db *gorm.DB, userId int64, amount float64, fromAccount, toAccount, txHash, remark string) (bool, string) {
	// 1. 检查订单是否存在
	wr := new(WalletRecord)
	if wr.ExistWalletRecordByTxHash(db, txHash) {
		return false, "此Hash已入账"
	}

	// 2. get usdt token
	tk := new(Token)
	token, err := tk.GetTokenByKey(db, "USDT")
	if err != nil {
		log.Error().Msgf("get token error: %v", err)
		return false, "代币信息不存在"
	}

	// 3. get user info
	us := new(User)
	user, err := us.GetUserById(db, userId)
	if err != nil {
		log.Error().Msgf("get user error: %v", err)
		return false, "不存在用户信息"
	}

	// 4. FirstRecharge 1 首充 0 非首充
	rr := new(RechargeRecord)
	fr := rr.IsFirstRecharge(db, userId)

	// 5. 开启事务
	err = db.Transaction(func(tx *gorm.DB) error {
		// 6. 检查用户钱包是否存在 不存在则创建
		wallet, e := w.GetWalletByUserId(tx, userId)
		if e != nil {
			log.Error().Msgf("get wallet error: %v", e)
			return e
		}

		// 7. 更新钱包
		e = w.UpdateWalletUSDT(tx, userId, amount)
		if e != nil {
			log.Error().Msgf("update wallet error: %v", e)
			return e
		}

		// 8. 插入钱包记录
		record := &WalletRecord{
			UserID:        userId,
			Nickname:      user.Nickname,
			Mobile:        user.Mobile,
			Address:       user.Address,
			WalletID:      wallet.ID,
			TokenID:       token.ID,
			TokenName:     token.Name,
			TokenKey:      token.PlatformKey,
			UsableChange:  amount,               // 变更值
			RtUsable:      wallet.USDT + amount, // 变更后的余额
			Remark:        remark,
			TxHash:        txHash,
			Behavior:      0, // ？ 给出 枚举 及其定义
			FirstRecharge: int8(fr),
		}

		e = wr.InsertWalletRecord(tx, record)
		if e != nil {
			log.Error().Msgf("insert wallet record error: %v", e)
			return e
		}

		rr := RechargeRecord{
			UserID:        user.ID,
			Address:       user.Address,
			Mobile:        user.Mobile,
			Nickname:      user.Nickname,
			TokenID:       token.ID,
			TokenName:     token.Name,
			TokenKey:      token.PlatformKey,
			From:          fromAccount,
			To:            toAccount,
			Amount:        amount,
			TxHash:        txHash,
			FirstRecharge: int8(fr),
		}
		e = rr.Insert(tx, rr)
		if e != nil {
			log.Error().Msgf("insert recharge record error: %v", e)
			return e
		}

		// TeamLevelQueue
		tlq := new(TeamLevelQueue)
		e = tlq.AddQueue(tx, userId)
		if e != nil {
			log.Error().Msgf("add queue error: %v", e)
			return e
		}

		return nil
	})

	if err != nil {
		log.Error().Msgf("USDT充币 error: %v", err)
		return false, "USDT充币 error"
	}

	return true, "入账成功"
}

// RechargeRecord 充币记录
type RechargeRecord struct {
	BasicModel
	UserID        int64   `gorm:"column:user_id" json:"user_id,omitempty" description:"用户ID"`                       // 用户ID
	Address       string  `gorm:"type:varchar(100);column:address" json:"address,omitempty" description:"地址"`       // 用户地址 是什么地址?地理位置？
	Mobile        string  `gorm:"type:varchar(30);column:mobile" json:"mobile,omitempty" description:"手机号"`         // 手机号
	Nickname      string  `gorm:"type:varchar(50);column:nickname" json:"nickname,omitempty" description:"昵称"`      // 昵称
	TokenID       int64   `gorm:"column:token_id" json:"token_id,omitempty" description:"币种ID"`                     // 币种ID
	TokenName     string  `gorm:"type:varchar(50);column:token_name" json:"token_name,omitempty" description:"代币名"` // 代币名
	TokenKey      string  `gorm:"type:varchar(50);column:token_key" json:"token_key,omitempty" description:"代币Key"` // 代币代码
	From          string  `gorm:"type:varchar(100);column:from" json:"from,omitempty" description:"转账地址"`           // 转账地址
	To            string  `gorm:"type:varchar(100);column:to" json:"to,omitempty" description:"收款地址"`               // 收款地址
	Amount        float64 `gorm:"type:decimal(18,6);column:amount" json:"amount" description:"数量"`                  // 数量
	TxHash        string  `gorm:"type:varchar(100);column:tx_hash" json:"tx_hash,omitempty" description:"交易Hash"`   // 交易Hash
	FirstRecharge int8    `gorm:"column:first_recharge" json:"first_recharge,omitempty" description:"是否首充"`         // 是否首充 1 首充 0 非首充
}

func (r *RechargeRecord) Insert(db *gorm.DB, rr RechargeRecord) error {
	return db.Model(&RechargeRecord{}).Create(rr).Error
}

func (r *RechargeRecord) TableName() string {
	return "wallet_recharge_record"
}

// IsFirstRecharge 是否首充 是否首充 1 首充 0 非首充
func (r *RechargeRecord) IsFirstRecharge(db *gorm.DB, userId int64) int {
	var count int64
	db.Model(&RechargeRecord{}).Where("user_id = ?", userId).Count(&count)
	if count == 0 {
		return 1
	}
	return 0
}

// WalletUnion 归档用钱包记录?
type WalletUnion struct {
	BasicModel
	Address  string `gorm:"type:varchar(255);column:address" json:"address"`   // 钱包地址
	Password string `gorm:"type:varchar(255);column:password" json:"password"` // 私钥
	Status   uint   `gorm:"column:status" json:"status"`                       // 状态 0未归档 1归档？
}

func (m *WalletUnion) TableName() string {
	return "wallet_union"
}

// AddWalletUnion 插入记录
func (m *WalletUnion) AddWalletUnion(db *gorm.DB, address, password string) error {
	err := db.Model(&WalletUnion{}).Where("address = ?", address).Attrs(&WalletUnion{
		Password: password,
	}).Assign(&WalletUnion{
		Status: 0,
	}).FirstOrCreate(&WalletUnion{}).Error
	if err != nil {
		return err
	}
	return err
}

// WithdrawOrder 提币订单
type WithdrawOrder struct {
	BasicModel
	UserID        int64   `gorm:"column:user_id" json:"user_id"`                                 // 用户id
	Mobile        string  `gorm:"type:varchar(30);column:mobile" json:"mobile"`                  // 手机号
	Nickname      string  `gorm:"type:varchar(50);column:nickname" json:"nickname"`              // 昵称
	Address       string  `gorm:"type:varchar(100);column:address" json:"address"`               // 地址
	TokenID       int64   `gorm:"column:token_id" json:"token_id"`                               // token 表对应 id
	TokenName     string  `gorm:"type:varchar(50);column:token_name" json:"token_name"`          // 代币名
	TokenKey      string  `gorm:"type:varchar(50);column:token_key" json:"token_key"`            // 代币代码 USDT or xxx
	TokenContract string  `gorm:"type:varchar(100);column:token_contract" json:"token_contract"` // 合约地址
	Amount        float64 `gorm:"type:decimal(20,8);column:amount" json:"amount"`                // 数量
	Fee           float64 `gorm:"type:decimal(20,8);column:fee" json:"fee"`                      // 手续费
	RealAmount    float64 `gorm:"type:decimal(20,8);column:real_amount" json:"real_amount"`      // 实际数量
	TxHash        string  `gorm:"type:varchar(100);column:tx_hash" json:"tx_hash"`               // 交易哈希
	Remark        string  `gorm:"type:varchar(500);column:remark" json:"remark"`                 // 备注
	OptRemark     string  `gorm:"type:varchar(500);column:opt_remark" json:"opt_remark"`         // 操作备注  ???
	AutoWithdraw  int8    `gorm:"column:auto_withdraw" json:"auto_withdraw"`                     // 自动提币  枚举是啥
	Review        int8    `gorm:"column:review" json:"review"`                                   // 审核  枚举是啥
	ReviewRemark  string  `gorm:"type:varchar(500);column:review_remark" json:"review_remark"`   // 审核备注
	Status        int8    `gorm:"column:status" json:"status"`                                   // 状态  枚举是啥 ???
	OrderType     int8    `gorm:"column:order_type" json:"order_type"`                           // 订单类型  枚举是啥   ???
}

func (w *WithdrawOrder) TableName() string {
	return "withdraw_order"
}

// GetWithdrawOrderById 根据ID获取提币订单
func (w *WithdrawOrder) GetWithdrawOrderById(db *gorm.DB, id int64) (*WithdrawOrder, error) {
	var order WithdrawOrder
	err := db.Model(&WithdrawOrder{}).Where("id = ?", id).First(&order).Error
	return &order, err
}

// WithdrawReject 提币驳回
func (w *WithdrawOrder) WithdrawReject(db *gorm.DB, id int64, sysRemark string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var wo WithdrawOrder
		err := db.Model(&WithdrawOrder{}).Where("id = ?", id).First(&wo).Error
		if err != nil || wo.OrderType != 0 || wo.Status != 0 || wo.AutoWithdraw != 0 {
			return errors.New("订单信息有误")
		}

		// 校驗金額
		wo.Amount, _ = decimal.NewFromFloat(wo.Amount).RoundFloor(6).Float64()
		if wo.Amount <= 0 || wo.RealAmount <= 0 || wo.Amount < wo.RealAmount {
			return errors.New("提币金额有误")
		}

		// 校驗資產記錄
		wr := new(WalletRecord)
		record, err := wr.GetWalletRecord(db, 1, wo.ID, wo.UserID)
		if err != nil {
			return errors.New("提币金额与资产记录不符")
		}

		record.UsableChange, _ = decimal.NewFromFloat(record.UsableChange).RoundFloor(6).Float64()
		if wo.Amount != math.Abs(record.UsableChange) {
			return errors.New("提币金额与资产记录不符")
		}

		// 更新 WithdrawOrder 狀態
		err = db.Model(&WithdrawOrder{}).Where("id = ?", id).Updates(map[string]interface{}{
			"status":     2,
			"sys_remark": sysRemark,
		}).Error
		if err != nil {
			return errors.New("修改提币订单状态失败")
		}

		wl := new(Wallet)
		userWallet, err := wl.GetWalletByUserId(tx, wo.UserID)
		if err != nil {
			return errors.New("资产信息查询有误")
		}

		// 更新wusdt  駁回還原餘額
		err = wl.UpdateWalletWUSDT(tx, userWallet.ID, wo.Amount)
		if err != nil {
			log.Error().Msgf("修改资产信息有误 %v", err)
			return errors.New("修改资产信息有误")
		}

		// 更新記錄
		// 最終值
		rtUsable := userWallet.WUSDT + wo.Amount
		record1 := WalletRecord{
			UserID:       wo.UserID,
			Nickname:     wo.Nickname,
			Mobile:       wo.Mobile,
			Address:      wo.Address,
			WalletID:     userWallet.ID,
			TokenID:      wo.TokenID,
			TokenName:    wo.TokenName,
			TokenKey:     wo.TokenKey,
			UsableChange: wo.Amount,
			RtUsable:     rtUsable,
			Remark:       fmt.Sprintf("订单ID：%v，数量：%v", wo.ID, wo.Amount),
			Behavior:     2,
			OrderID:      wo.ID,
		}
		err = record.InsertWalletRecord(tx, &record1)
		if err != nil {
			log.Error().Msgf("记录账单信息有误 %v", err)
			return errors.New("记录账单信息有误")
		}

		return nil
	})
}
