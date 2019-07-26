package _stack


/**
* 基于数组作为底层数据结构实现的栈的功能
*/

type ArrayStack struct {
	//数据
	data [] interface{}
	//栈顶指针
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (_this *ArrayStack) IsEmpty() bool {
	if _this.top < 0 {
		return true
	}
	return false
}

func (_this *ArrayStack) Push(value interface{}) {

	if _this.top<0 {
		_this.top = 0
	}else {
		_this.top += 1
	}

	if _this.top > len(_this.data)-1 {
		_this.data = append(_this.data,value)
	}else  {
		_this.data[_this.top] = value
	}
}

func (_this *ArrayStack) Pop() interface{} {
	if _this.IsEmpty() {
		return nil
	}

	value := _this.data[_this.top]  //数据获取栈定元素
	_this.top -=1 //栈顶指针下移
	return  value
}

func (_this *ArrayStack) Top() interface{} {
	if _this.IsEmpty() {
		return nil
	}
	value := _this.data[_this.top]  //数据获取栈定元素
	return  value
}

/**
* 清空堆栈数据
*
*只要把 TOP指向-1，数据中的数据暂时不用清空
 */
func (_this *ArrayStack) Flush() {
	_this.top = -1
}