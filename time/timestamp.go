package time

import (
	"time"
)

func GetLastMonthTimestamp() (int64, int64) {
	// 获取当前时间
	now := time.Now()

	// 计算上个月的开始时间
	lastMonthStart := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.UTC)
	lastMonthStartUnix := lastMonthStart.Unix()

	// 计算上个月的结束时间
	lastMonthEnd := time.Date(now.Year(), now.Month()-1, 1, 23, 59, 59, 0, time.UTC).AddDate(0, 1, -1)
	lastMonthEndUnix := lastMonthEnd.Unix()

	return lastMonthStartUnix, lastMonthEndUnix
}
