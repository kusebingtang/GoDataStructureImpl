package _stack

type Stack interface {
	Push(v interface{}) 	/// 入栈
	Pop() interface{}   	/// 出栈
	IsEmpty() bool     	    /// 是否为空
	Top() interface{}		///返回队列头部元素
	Flush()					///清空堆栈
}
