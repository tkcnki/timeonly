package timeonly

import (
	"fmt"
	"time"
)

type TimeOnly struct {
	time.Time
}

// NewTimeOnly 関数を使用して TimeOnly インスタンスを作成
func NewTimeOnly(hour, minute, second int) TimeOnly {
	return TimeOnly{time.Date(0, 1, 1, hour, minute, second, 0, time.UTC)}
}

// NewTimeOnlyFromString 日付文字列からTimeOnlyインスタンスを作成
func NewTimeOnlyFromString(timeString, layout string) (TimeOnly, error) {
	parsedTime, err := time.Parse(layout, timeString)
	if err != nil {
		return TimeOnly{}, err
	}
	return TimeOnly{parsedTime}, nil
}

// NewTimeOnlyFromStringAndLocation 関数を使用して TimeOnly インスタンスを文字列とロケーションから作成
func NewTimeOnlyFromStringAndLocation(timeString, locationName string) (TimeOnly, error) {
	// ロケーションを取得
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return TimeOnly{}, err
	}

	// ロケーションを指定して日付文字列をパース
	parsedTime, err := time.ParseInLocation("15:04:05", timeString, loc)
	if err != nil {
		return TimeOnly{}, err
	}

	return TimeOnly{parsedTime}, nil
}

/*
String 文字列を返す。
  - フォーマットは"15:04:05"
  - hourは24時間以上の数値を取る
*/
func (s *TimeOnly) String() string {
	d := s.Sub(time.Time{})
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
