package main

import (
    "fmt"
)

// Array - динамический массив
type Array struct {
    data []interface{}
    size int
}

// NewArray создает новый массив
func NewArray() *Array {
    return &Array{
        data: make([]interface{}, 0),
        size: 0,
    }
}

// Add добавляет элемент по индексу
func (a *Array) Add(index int, value interface{}) bool {
    if index < 0 || index > a.size {
        return false
    }
    
    a.data = append(a.data, nil)
    copy(a.data[index+1:], a.data[index:])
    a.data[index] = value
    a.size++
    return true
}

// Append добавляет элемент в конец массива
func (a *Array) Append(value interface{}) {
    a.data = append(a.data, value)
    a.size++
}

// Get возвращает элемент по индексу
func (a *Array) Get(index int) (interface{}, bool) {
    if index < 0 || index >= a.size {
        return nil, false
    }
    return a.data[index], true
}

// Remove удаляет элемент по индексу
func (a *Array) Remove(index int) bool {
    if index < 0 || index >= a.size {
        return false
    }
    
    a.data = append(a.data[:index], a.data[index+1:]...)
    a.size--
    return true
}

// Replace заменяет элемент по индексу
func (a *Array) Replace(index int, value interface{}) bool {
    if index < 0 || index >= a.size {
        return false
    }
    a.data[index] = value
    return true
}

// Length возвращает длину массива
func (a *Array) Length() int {
    return a.size
}

// String возвращает строковое представление массива
func (a *Array) String() string {
    return fmt.Sprintf("%v", a.data[:a.size])
}
