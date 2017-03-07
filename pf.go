package pf

import (
	"os"
	"time"
)

//Date is structure of date.
type Date struct {
	year  int
	month int
	day   int
}

//NewDate returns Date instance
func NewDate(y, m, d int) Date {
	return Date{year: y, month: m, day: d}
}

//NewYearMonth returns Date instance
func NewYearMonth(y, m int) Date {
	return Date{year: y, month: m, day: 0}
}

//GetPremiumFriday returns day of premium friday
func (d Date) GetPremiumFriday() (int, error) {
	if d.day != 0 { //日付はセットされていないこと
		return 0, os.ErrInvalid
	}
	tm := d.toTime()
	if !validTime(tm) { //範囲チェック
		return 0, os.ErrInvalid
	}
	return tm.Day(), nil //計算済みの値を返す
}

//IsPremiumFriday returns true if Premium Friday
func (d Date) IsPremiumFriday() bool {
	if d.day == 0 { //日付はセットされていること
		return false
	}
	tm := d.toTime()
	if !validTime(tm) { //範囲チェック
		return false
	}
	if tm.Weekday() != time.Friday { //金曜日ではない
		return false
	}
	if tm.AddDate(0, 0, 7).Month() == tm.Month() { //1週間後も同じ月なら最終金曜日ではない
		return false
	}
	return true
}

func (d Date) toTime() time.Time {
	if d.day == 0 {
		//指定月末（翌月0日）で初期化する
		tm := time.Date(d.year, (time.Month)(d.month+1), 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で
		//最終金曜日の日付を求める
		w := tm.Weekday() - time.Friday
		if w < 0 {
			w += 7
		}
		return tm.AddDate(0, 0, -(int)(w))
	}
	return time.Date(d.year, (time.Month)(d.month), d.day, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で
}

func validTime(tm time.Time) bool {
	if tm.Year() < 2017 || (tm.Year() == 2017 && tm.Month() < time.February) { //2017年1月までは実施前なので無効
		return false
	}
	return true
}
