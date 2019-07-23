package _linkedlist


type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(_value interface{}) *ListNode {
	return  &ListNode{
		next: nil,
		value: _value,
	}
}

/**
 * 获取当前节点下一个Node
 */
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

/**
 * 获取当前节点的值
 */
func (this *ListNode) GetValue() interface{}  {
	return  this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

/**
 *在某个节点后面插入节点
 */
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newListNode := NewListNode(v)
	oldNext := p.next  //链表插入操作
	p.next = newListNode
	newListNode.next = oldNext
	this.length++

	return true
}

/**
 * 在某个节点前面插入节点
 */
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}

	current := this.head.next
	pre := this.head

 	for nil != current {
 		if current == p {
 			break
		}

 		current = current.next
 		pre = current
	}

	if current == nil {
		return  false
	}

	newListNode := NewListNode(v)
	pre.next = newListNode
	newListNode.next = p

	this.length++
	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}


func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}

	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ { //链表的index查找，需要从表头开始next index次，才能获取到ListNode  和数组对比
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	currentNode := this.head.next
	preNode := this.head

	for nil != currentNode  {
		if currentNode ==p {
			break
		}

		currentNode = currentNode.next
		preNode = currentNode
	}

	if nil == currentNode {
		return false
	}

	preNode.next = currentNode.next
	p = nil
	this.length--
	return  true
}