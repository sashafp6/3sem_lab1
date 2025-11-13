// Stack - структура стека
type Stack struct {
    data []interface{}
    size int
}

// NewStack создает новый стек
func NewStack() *Stack {
    return &Stack{
        data: make([]interface{}, 0),
        size: 0,
    }
}

// Push добавляет элемент в стек
func (s *Stack) Push(value interface{}) {
    s.data = append(s.data, value)
    s.size++
}

// Pop удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (interface{}, bool) {
    if s.size == 0 {
        return nil, false
    }
    
    value := s.data[s.size-1]
    s.data = s.data[:s.size-1]
    s.size--
    return value, true
}

// Peek возвращает верхний элемент без удаления
func (s *Stack) Peek() (interface{}, bool) {
    if s.size == 0 {
        return nil, false
    }
    return s.data[s.size-1], true
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
    return s.size == 0
}

// Size возвращает размер стека
func (s *Stack) Size() int {
    return s.size
}

// String возвращает строковое представление стека
func (s *Stack) String() string {
    return fmt.Sprintf("%v", s.data)
}
