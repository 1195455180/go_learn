package data

type Object interface {}

type Node struct {
	data Object
	next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}
func main()  {

}

//链表初始化
func (list *List) Init()  {
	(*list).size = 0  //此时链表是空的
	(*list).head = nil //没有车头
	(*list).tail = nil //没有车尾
}

func (list *List) append(node *Node) bool {
	if node ==nil{
		return false
	}
	(*node).next = nil
	if (*list).size == 0{
		(*list).head = node  //这是单向链表的第一个元素，也是链表的头部
	}else {
		oldTail := (*list).tail
		(*oldTail).next = node
	}
	(*list).tail = node
	(*list).size++

	return true
}

func (list *List) Insert(i int,node *Node) bool {
	// 空的节点、索引超出范围和空链表都无法做插入操作
	if node == nil || i > (*list).size || (*list).size == 0 {
		return false
	}

	if i==0{
		(*node).next = (*list).head
		(*list).head = node
	}else{
		// 找到前一个元素
		preItem := (*list).head
		for j := 1 ; j < i; j++ { // 数前面i个元素
			preItem = (*preItem).next
		}
		// 原来元素放到新元素后面,新元素放到前一个元素后面
		(*node).next = (*preItem).next
		(*preItem).next = preItem
	}

	(*list).size++
	return true
}

func (list *List) Remove(i int, node *Node) bool {
	if i >= (*list).size {
		return false
	}

	if i == 0 { // 删除头部
		node = (*list).head
		(*list).head = (*node).next
		if (*list).size == 1 { // 如果只有一个元素，那尾部也要调整
			(*list).tail = nil
		}
	} else {
		preItem := (*list).head
		for j := 1; j < i; j++ {
			preItem = (*preItem).next
		}

		node = (*preItem).next
		(*preItem).next = (*node).next

		if i == ((*list).size - 1) { // 若删除的尾部，尾部指针需要调整
			(*list).tail = preItem
		}
	}
	(*list).size--
	return true
}

func (list *List) Get(i int) *Node {
	if i >= (*list).size {
		return nil
	}

	item := (*list).head
	for j := 0; j < i ; j++ {    // 从head数i个
		item = (*item).next
	}

	return item
}


