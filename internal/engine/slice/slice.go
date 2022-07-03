package slice

import (
	"github.com/EAFA0/Toast/internal/engine/event"
	"github.com/EAFA0/Toast/internal/engine/task"
)

type ID uint64

// Body 中间态数据
type Body struct {
	Event event.ID // 对应 event id
	Scene string   // 计算现场
	Task  task.ID  // 对应 task id
	Data  string   // 主体数据
	Pre   ID       // 前一切片
}

// Next 构造下一块切片
func Next(base Body, item event.Body) (Body, error) {
	return Body{}, nil
}
