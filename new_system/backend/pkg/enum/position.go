package enum

type Direction string // 方向

const (
	BUY  Direction = "BUY"  // 做多
	SELL Direction = "SELL" // 做空
)

type FollowDirection string // 跟随方向

const (
	POSITIVE FollowDirection = "POSITIVE" // 正
	NEGATIVE FollowDirection = "NEGATIVE" // 反
)
