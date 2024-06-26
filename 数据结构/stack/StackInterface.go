package stack

// Stack 栈接口
type Stack interface {
	// Push 压栈
	Push(v interface{})
	// Pop 出栈
	Pop() interface{}
	// IsEmpty 是否为空
	IsEmpty() bool
	// Top 栈顶元素
	Top() interface{}
	// Flush 清空栈
	Flush()
	Print()
}
