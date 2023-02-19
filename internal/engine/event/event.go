package event

import "context"

type ID uint64

type Data struct {
	ID        ID     // 主键 ID, 唯一自增
	TTL       int    // 每绑定一个 Slice + 1, 0 时删除
	Data      []byte // 事件主体
	TimeStamp int64  // 产生时间
}

// Store 存储一个事件信息
func Store(ctx context.Context, data Data) (ID, error) {
	return 0, nil
}

// Remove 删除一个事件
func Remove(ctx context.Context, id ID) error {
	return nil
}
