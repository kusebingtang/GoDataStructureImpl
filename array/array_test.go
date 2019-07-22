package _array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 20
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-5; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	arr.Insert(uint(87),100)
	arr.Print()

	arr.Insert(uint(18),100)
	arr.Print()

	fmt.Println(arr.Len())
}
