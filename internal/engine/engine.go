// engine 调度器相关逻辑及其实现

package engine

import (
	"context"

	"github.com/EAFA0/Toast/internal/engine/event"
	"github.com/EAFA0/Toast/internal/engine/slice"
	"github.com/EAFA0/Toast/internal/engine/task"
)

// Append 追加一个事件
func Append(ctx context.Context, tID task.ID, body event.Data, sID ...slice.ID) error {
	return nil
}

// Insert 插入一个事件
func Insert(ctx context.Context, tID task.ID, body event.Data, sID ...slice.ID) error {
	return nil
}
