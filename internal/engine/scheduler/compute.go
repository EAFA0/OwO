// scheduler 调度器相关逻辑及其实现

package scheduler

import (
	"context"

	"github.com/EAFA0/toast/internal/do/user"
	"github.com/EAFA0/toast/internal/engine/task"
)

// Compute 推进任务进度
func Compute(ctx context.Context, tID task.ID, uID user.ID) error {
	return nil
}