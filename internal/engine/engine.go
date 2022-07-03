// engine 调度器相关逻辑及其实现

package engine

import (
	"context"

	"github.com/EAFA0/Toast/internal/engine/event"
	"github.com/EAFA0/Toast/internal/engine/slice"
	"github.com/EAFA0/Toast/internal/engine/task"
)

// Append 追加进度
func Append(ctx context.Context, tID task.ID, body event.Body) error {
	return nil
}

// ReBuild 重建进度
func ReBuild(ctx context.Context, tID task.ID, begin slice.ID) error {
	return nil
}
