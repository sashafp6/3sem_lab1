// DoubleListNode - узел двусвязного списка
type DoubleListNode struct {
    value interface{}
    prev  *DoubleListNode
    next  *DoubleListNode
}

// DoublyLinkedList - двусвязный список
type DoublyLinkedList struct {
    head *DoubleListNode
    tail *DoubleListNode
    size int
}

// NewDoublyLinkedList создает новый двусвязный список
func NewDoublyLinkedList() *DoublyLinkedList {
    return &DoublyLinkedList{nil, nil, 0}
}

// AddToHead добавляет элемент в начало списка
func (d *DoublyLinkedList) AddToHead(value interface{}) {
    newNode := &DoubleListNode{value: value}
    
    if d.head == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        newNode.next = d.head
        d.head.prev = newNode
        d.head = newNode
    }
    d.size++
}

// AddToTail добавляет элемент в конец списка
func (d *DoublyLinkedList) AddToTail(value interface{}) {
    newNode := &DoubleListNode{value: value}
    
    if d.tail == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        newNode.prev = d.tail
        d.tail.next = newNode
        d.tail = newNode
    }
    d.size++
}

// AddAfter добавляет элемент после определенного значения
func (d *DoublyLinkedList) AddAfter(target interface{}, value interface{}) bool {
    current := d.head
    for current != nil {
        if current.value == target {
            newNode := &DoubleListNode{value: value, prev: current, next: current.next}
            
            if current.next != nil {
                current.next.prev = newNode
            } else {
                d.tail = newNode
            }
            current.next = newNode
            d.size++
            return true
        }
        current = current.next
    }
    return false
}

// AddBefore добавляет элемент перед определенным значением
func (d *DoublyLinkedList) AddBefore(target interface{}, value interface{}) bool {
    current := d.head
    for current != nil {
        if current.value == target {
            newNode := &DoubleListNode{value: value, prev: current.prev, next: current}
            
            if current.prev != nil {
                current.prev.next = newNode
            } else {
                d.head = newNode
            }
            current.prev = newNode
            d.size++
            return true
        }
        current = current.next
    }
    return false
}

// RemoveHead удаляет первый элемент
func (d *DoublyLinkedList) RemoveHead() bool {
    if d.head == nil {
        return false
    }
    
    if d.head == d.tail {
        d.head = nil
        d.tail = nil
    } else {
        d.head = d.head.next
        d.head.prev = nil
    }
    d.size--
    return true
}

// RemoveTail удаляет последний элемент
func (d *DoublyLinkedList) RemoveTail() bool {
    if d.tail == nil {
        return false
    }
    
    if d.head == d.tail {
        d.head = nil
        d.tail = nil
    } else {
        d.tail = d.tail.prev
        d.tail.next = nil
    }
    d.size--
    return true
}

// RemoveAfter удаляет элемент после определенного значения
func (d *DoublyLinkedList) RemoveAfter(target interface{}) bool {
    current := d.head
    for current != nil && current.next != nil {
        if current.value == target {
            nodeToRemove := current.next
            
            if nodeToRemove == d.tail {
                d.tail = current
                current.next = nil
            } else {
                current.next = nodeToRemove.next
                nodeToRemove.next.prev = current
            }
            d.size--
            return true
        }
        current = current.next
    }
    return false
}

// RemoveBefore удаляет элемент перед определенным значением
func (d *DoublyLinkedList) RemoveBefore(target interface{}) bool {
    if d.head == nil || d.head.next == nil {
        return false
    }
    
    current := d.head.next
    for current != nil {
        if current.value == target {
            nodeToRemove := current.prev
            
            if nodeToRemove == d.head {
                d.head = current
                current.prev = nil
            } else {
                current.prev = nodeToRemove.prev
                nodeToRemove.prev.next = current
            }
            d.size--
            return true
        }
        current = current.next
    }
    return false
}

// RemoveByValue удаляет элемент по значению
func (d *DoublyLinkedList) RemoveByValue(value interface{}) bool {
    current := d.head
    for current != nil {
        if current.value == value {
            if current == d.head {
                return d.RemoveHead()
            } else if current == d.tail {
                return d.RemoveTail()
            } else {
                current.prev.next = current.next
                current.next.prev = current.prev
                d.size--
                return true
            }
        }
        current = current.next
    }
    return false
}

// Find ищет элемент по значению
func (d *DoublyLinkedList) Find(value interface{}) (*DoubleListNode, bool) {
    current := d.head
    for current != nil {
        if current.value == value {
            return current, true
        }
        current = current.next
    }
    return nil, false
}

// ReadForward читает список с начала
func (d *DoublyLinkedList) ReadForward() []interface{} {
    result := make([]interface{}, 0)
    current := d.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}

// ReadBackward читает список с конца
func (d *DoublyLinkedList) ReadBackward() []interface{} {
    result := make([]interface{}, 0)
    current := d.tail
    for current != nil {
        result = append(result, current.value)
        current = current.prev
    }
    return result
}

// Size возвращает размер списка
func (d *DoublyLinkedList) Size() int {
    return d.size
}

// String возвращает строковое представление списка
func (d *DoublyLinkedList) String() string {
    return fmt.Sprintf("%v", d.ReadForward())
}
