// Slice 是 Event 的一个有序集合, 多个 Slice 可以包含同一个 Event
package slice

import (
	"context"
	"errors"

	"github.com/EAFA0/Toast/internal/engine/event"
)

type ID uint64

// Info Slice 的一些基础属性
type Info struct {
	ReComputable bool  // 是否可重算, 只有可重算的 Slice 会记录事件的顺序
	ExpireTime   int64 // 超出过期时间的事件绑定关系会被销毁, 单位 s
	Size         int64 // 最多只存储指定长度的事件绑定关系
}

var ErrIsFirst = errors.New("is the first")
var ErrIsLast = errors.New("is the last")

// Next 获取下一个事件
func Next(ctx context.Context, id ID, eventID event.ID) (event.Data, error) {
	return event.Data{}, nil
}

// Pervious 获取前一个事件
func Pervious(ctx context.Context, id ID, eventID event.ID) (event.Data, error) {
	return event.Data{}, nil
}

// Bind 将指定事件绑定到指定 slice 上
func Bind(ctx context.Context, data event.Data, id ...ID) error {
	return nil
}
