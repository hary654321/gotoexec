package global

type LoginLog struct {
	Time int64
	Ip   string
}

// FixedSizeStack 固定大小的栈结构体
type FixedSizeStack struct {
	items    []LoginLog // 存储栈元素的切片
	capacity int        // 栈的最大容量
	top      int        // 栈顶索引
}

var FixedSizeStackInstance *FixedSizeStack

func init() {
	FixedSizeStackInstance = NewFixedSizeStack(50)
}

// NewFixedSizeStack 创建一个新的固定大小栈
func NewFixedSizeStack(capacity int) *FixedSizeStack {
	return &FixedSizeStack{
		items:    make([]LoginLog, capacity),
		capacity: capacity,
		top:      -1, // 初始化时栈顶索引为-1，表示栈为空
	}
}

// Push 入栈方法
func (s *FixedSizeStack) Push(log LoginLog) {
	if s.top >= s.capacity-1 {
		// 如果栈满，删除最老的数据（尾部）
		s.top = (s.top + 1) % s.capacity
	} else {
		// 如果栈未满，栈顶索引递增
		s.top++
	}
	// 将新元素放入栈顶位置
	s.items[s.top] = log
}

// Pop 出栈方法（可选，根据需要实现）
func (s *FixedSizeStack) Pop() *LoginLog {
	if s.top == -1 {
		return nil // 栈为空时返回nil
	}
	item := s.items[s.top]
	s.top-- // 栈顶索引递减
	return &item
}

func (s *FixedSizeStack) Get() []LoginLog {

	return s.items
}
