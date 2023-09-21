package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

func newList() *List {
	return &List{len: 0, firstNode: nil}
}

func (l *List) Add(data int64) (id int64) {
	newNode := &node{data: data}

	if l.firstNode == nil {
		l.firstNode = newNode
		l.firstNode.index = 0
		l.len++
		return 0
	}

	nod := l.firstNode
	id = 1
	for ; nod.next != nil; nod = nod.next {
		id++
	}
	nod.next = newNode
	nod.next.index = id
	l.len++
	return id
}

func (l *List) Print() {
	if l.firstNode == nil {
		fmt.Println("no data")
		return
	}
	for nod := l.firstNode; nod != nil; nod = nod.next {
		fmt.Println(nod.data)
	}
}

func (l *List) Print_All() {
	if l.firstNode == nil {
		fmt.Println("no data")
		return
	}
	for nod := l.firstNode; nod != nil; nod = nod.next {
		fmt.Println(nod.data, nod.index)
	}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if l.firstNode == nil {
		fmt.Println("no data")
		return
	}
	if id == 0 {
		if l.len > 1 {
			l.firstNode = l.firstNode.next
			l.refresh_indices()
			return
		}
		if l.len == 1 {
			l = &List{len: 0, firstNode: nil}
			l.refresh_indices()
			return
		}
	}
	if id >= l.len {
		fmt.Println("index out of range")
		return
	}
	if id < 0 {
		fmt.Println("give positive index")
		return
	}

	del_nod := l.find_node(id)
	prev_nod := l.find_node(id - 1)
	prev_nod.next = del_nod.next
	l.len--
	l.refresh_indices()
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) bool {
	if l.firstNode == nil {
		fmt.Println("no data")
		return false
	}
	var id int64 = 0
	for nod := l.firstNode; nod != nil; nod = nod.next {
		if nod.data == value {
			if id == 0 {
				l.firstNode = l.firstNode.next
				l.refresh_indices()
				return true
			}
			if id+1 < l.len {
				l.find_node(id - 1).next = l.find_node(id + 1)
				l.refresh_indices()
				l.len--
				return true
			}
			l.find_node(id - 1).next = nil
			l.refresh_indices()
			l.len--
			return true
		}
		id++
	}
	fmt.Println("not found")
	return false
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.firstNode == nil {
		fmt.Println("no data")
		return
	}
	for {
		res := l.RemoveByValue(value)
		if !res {
			return
		}
		// нужно узнать проблема ли, что выводится на экран 'not found'
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	nod := l.find_node(id)
	if nod != nil {
		return nod.data, true
	}
	return 0, false
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (index int64, ok bool) {
	if l.firstNode == nil {
		fmt.Println("no data")
		return 0, false
	}
	var id int64 = 0
	for nod := l.firstNode; nod != nil; nod = nod.next {
		if nod.data == value {
			return id, true
		}
		id++
	}
	fmt.Println("not found")
	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	if l.firstNode == nil {
		fmt.Println("no data")
		return nil, false
	}
	var id int64 = 0
	for nod := l.firstNode; nod != nil; nod = nod.next {
		if nod.data == value {
			ids = append(ids, id)
		}
		id++
	}
	if len(ids) > 0 {
		return ids, true
	}
	fmt.Println("not found")
	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.firstNode == nil {
		fmt.Println("no data")
		return nil, false
	}
	for nod := l.firstNode; nod != nil; nod = nod.next {
		values = append(values, nod.data)
	}
	if len(values) > 0 {
		return values, true
	}
	fmt.Println("not found")
	return nil, false
}

// Clear очищает список
func (l *List) Clear() {
	l.len = 0
	l.firstNode = nil
}

func (l *List) refresh_indices() {
	var id int64 = 0
	for nod := l.firstNode; nod != nil; nod = nod.next {
		nod.index = id
		id++
	}
}

func (l *List) find_node(index int64) *node {
	for nod := l.firstNode; nod != nil; nod = nod.next {
		if nod.index == index {
			return nod
		}
	}
	return nil
}
