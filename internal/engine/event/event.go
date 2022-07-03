package event

import "errors"

type ID uint64

type Body struct {
	Data      string // 事件主体
	TimeStamp int64  // 产生时间
}

var ErrIsLast = errors.New("is the last")

// First 获取指定时间点后第一个
func First(timestamp int64) (Body, error) {
	return Body{}, nil
}

// Next 获取下一个事件
func Next(id ID) (Body, error) {
	return Body{}, nil
}
