// ListNode - узел односвязного списка
type ListNode struct {
    value interface{}
    next  *ListNode
}

// SinglyLinkedList - односвязный список
type SinglyLinkedList struct {
    head *ListNode
    size int
}

// NewSinglyLinkedList создает новый односвязный список
func NewSinglyLinkedList() *SinglyLinkedList {
    return &SinglyLinkedList{nil, 0}
}

// AddToHead добавляет элемент в начало списка
func (s *SinglyLinkedList) AddToHead(value interface{}) {
    newNode := &ListNode{value: value, next: s.head}
    s.head = newNode
    s.size++
}

// AddToTail добавляет элемент в конец списка
func (s *SinglyLinkedList) AddToTail(value interface{}) {
    newNode := &ListNode{value: value}
    
    if s.head == nil {
        s.head = newNode
    } else {
        current := s.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    s.size++
}

// AddAfter добавляет элемент после определенного значения
func (s *SinglyLinkedList) AddAfter(target interface{}, value interface{}) bool {
    current := s.head
    for current != nil {
        if current.value == target {
            newNode := &ListNode{value: value, next: current.next}
            current.next = newNode
            s.size++
            return true
        }
        current = current.next
    }
    return false
}

// AddBefore добавляет элемент перед определенным значением
func (s *SinglyLinkedList) AddBefore(target interface{}, value interface{}) bool {
    if s.head == nil {
        return false
    }
    
    if s.head.value == target {
        s.AddToHead(value)
        return true
    }
    
    prev := s.head
    current := s.head.next
    
    for current != nil {
        if current.value == target {
            newNode := &ListNode{value: value, next: current}
            prev.next = newNode
            s.size++
            return true
        }
        prev = current
        current = current.next
    }
    return false
}

// RemoveHead удаляет первый элемент
func (s *SinglyLinkedList) RemoveHead() bool {
    if s.head == nil {
        return false
    }
    s.head = s.head.next
    s.size--
    return true
}

// RemoveTail удаляет последний элемент
func (s *SinglyLinkedList) RemoveTail() bool {
    if s.head == nil {
        return false
    }
    
    if s.head.next == nil {
        s.head = nil
    } else {
        current := s.head
        for current.next.next != nil {
            current = current.next
        }
        current.next = nil
    }
    s.size--
    return true
}

// RemoveAfter удаляет элемент после определенного значения
func (s *SinglyLinkedList) RemoveAfter(target interface{}) bool {
    current := s.head
    for current != nil && current.next != nil {
        if current.value == target {
            current.next = current.next.next
            s.size--
            return true
        }
        current = current.next
    }
    return false
}

// RemoveBefore удаляет элемент перед определенным значением
func (s *SinglyLinkedList) RemoveBefore(target interface{}) bool {
    if s.head == nil || s.head.next == nil || s.head.value == target {
        return false
    }
    
    if s.head.next.value == target {
        s.head = s.head.next
        s.size--
        return true
    }
    
    prev := s.head
    current := s.head.next
    next := current.next
    
    for next != nil {
        if next.value == target {
            prev.next = next
            s.size--
            return true
        }
        prev = current
        current = next
        next = next.next
    }
    return false
}

// RemoveByValue удаляет элемент по значению
func (s *SinglyLinkedList) RemoveByValue(value interface{}) bool {
    if s.head == nil {
        return false
    }
    
    if s.head.value == value {
        s.head = s.head.next
        s.size--
        return true
    }
    
    prev := s.head
    current := s.head.next
    
    for current != nil {
        if current.value == value {
            prev.next = current.next
            s.size--
            return true
        }
        prev = current
        current = current.next
    }
    return false
}

// Find ищет элемент по значению
func (s *SinglyLinkedList) Find(value interface{}) (*ListNode, bool) {
    current := s.head
    for current != nil {
        if current.value == value {
            return current, true
        }
        current = current.next
    }
    return nil, false
}

// ReadForward читает список с начала
func (s *SinglyLinkedList) ReadForward() []interface{} {
    result := make([]interface{}, 0)
    current := s.head
    for current != nil {
        result = append(result, current.value)
        current = current.next
    }
    return result
}

// ReadBackward читает список с конца (через массив)
func (s *SinglyLinkedList) ReadBackward() []interface{} {
    forward := s.ReadForward()
    backward := make([]interface{}, len(forward))
    for i, j := 0, len(forward)-1; i < len(forward); i, j = i+1, j-1 {
        backward[i] = forward[j]
    }
    return backward
}

// Size возвращает размер списка
func (s *SinglyLinkedList) Size() int {
    return s.size
}

// String возвращает строковое представление списка
func (s *SinglyLinkedList) String() string {
    return fmt.Sprintf("%v", s.ReadForward())
}
