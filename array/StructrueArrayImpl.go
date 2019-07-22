package _array

import (
	"errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）存储数组中的数据是int类型的；
 *
 * Author: JBW
 *
 */
type Array struct {

	length  uint
	data 	[]int
}


/**
 * 返回当前数据的长度
 */
func (_this *Array) Len() uint {
	return _this.length
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity==0 {
		return  nil
	}

	array := Array{
		data: make([]int,capacity,capacity),
		length: 0,
	}
	return &array
}

//通过索引查找数组，索引范围[0,n-1]
func (_this *Array) Find(index uint) (int, error) {
	if index > uint(cap(_this.data)) {
		return 0, errors.New("out of index range")
	}
	return _this.data[index],nil
}

/**
 *	插入数值到索引index上
 * index 数组索引
 * value 插入的数值
 */
func (_this *Array) Insert(index uint, value int) error {

	if index > uint(cap(_this.data)) {
		return  errors.New("out of index range")
	}

	if _this.Len() == uint(cap(_this.data)) {
		return  errors.New("array full ")
	}

	for i := _this.length; i > index; i-- {
		_this.data[i] = _this.data[i-1]
	}
	_this.data[index] = value
	_this.length++

	return nil
}


func (_this *Array) Delete(index uint) (int, error) {
	if index > uint(cap(_this.data)) {
		return 0, errors.New("out of index range")
	}
	v := _this.data[index]
	for i := index; i < _this.Len()-1; i++ {
		_this.data[i] = _this.data[i+1]
	}
	_this.length--
	return v, nil
}

//打印数列
func (_this *Array) Print() {
	var format string
	for i := uint(0); i < _this.Len(); i++ {
		format += fmt.Sprintf("|%+v    ", _this.data[i])
	}
	fmt.Println(format)
}