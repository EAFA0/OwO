// task 一个 task 就是一个计算任务, 一个计算任务必须绑定一个切片才能正常计算
// 多个任务可以绑定同一个切片, 他们之间的计算进度可以不一样, 允许部分任务落后

package task

import (
	"context"

	"github.com/EAFA0/OwO/internal/engine/event"
	"github.com/EAFA0/OwO/internal/engine/slice"
)

type ID uint64

// DataID 任务数据对应的 ID
type DataID uint64

// Data 任务进度数据
type Data struct {
	ID       DataID   // 唯一 id
	Task     ID       // 任务 id
	Body     []byte   // 数据主体
	Event    event.ID // 当前目标事件
	Version  int      // 所属计算版本
	Previous DataID   // 前置数据 id
}

// 接受一个当前状态跟输入数据, 输出新的状态
type Task func(ctx context.Context, status, input []byte) (output []byte, err error)

// Bind 绑定切片跟任务
func Bind(ctx context.Context, tID ID, sID slice.ID) error {
	return nil
}
