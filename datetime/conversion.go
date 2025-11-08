// Package datetime 时间处理
package datetime

import (
	"fmt"
	"math/rand"
	"time"
)

// IsTimeInRangeStr 判断 timeStr 是否在 [startStr, endStr) 范围内，仅比较时分秒
func IsTimeInRangeStr(timeStr, startStr, endStr, layout string) (bool, error) {
	// 若要支持秒，可改成 "15:04:05"

	parse := func(s string) (time.Time, error) {
		// 使用固定日期，保证只比较时间部分
		t, err := time.ParseInLocation(layout, s, time.Local)
		if err != nil {
			return time.Time{}, fmt.Errorf("解析时间失败 %s: %w", s, err)
		}
		return t, nil
	}

	t, err1 := parse(timeStr)
	s, err2 := parse(startStr)
	e, err3 := parse(endStr)
	if err1 != nil || err2 != nil || err3 != nil {
		return false, fmt.Errorf("时间解析错误: %v %v %v", err1, err2, err3)
	}

	// 正常范围，例如 08:00–12:00
	if e.After(s) {
		return t.Equal(s) || (t.After(s) && t.Before(e)), nil

	}

	// 跨午夜范围，例如 22:00–06:00
	return t.Equal(s) || t.After(s) || t.Before(e), nil
}

// RandomTimeInRange 在一个范围内生成随机时间
func RandomTimeInRange(start, end time.Time) (*time.Time, error) {
	if start.After(end) {
		return nil, fmt.Errorf("start时间不能晚于end时间")
	}

	if start.Equal(end) {
		return &start, nil
	}

	duration := end.Sub(start)
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	randomTime := start.Add(randomDuration)
	return &randomTime, nil
}
