package pf

import (
	"os"
	"time"
)

//GetPremiumFriday returns day of premium friday
func GetPremiumFriday(y int, m time.Month) (int, error) {
	//引数のチェック
	if y < 2017 || m < time.January || m > time.December {
		return 0, os.ErrInvalid
	}
	if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
		return 0, os.ErrInvalid
	}

	//指定月末（翌月0日）で初期化する
	tm := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で

	w := tm.Weekday() - time.Friday
	if w < 0 {
		w += 7
	}
	return tm.Day() - (int)(w), nil
}
