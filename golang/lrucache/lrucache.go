package lrucache

import "container/list"

type Node[T any] struct {
	Data    T
	Element *list.Element
}

type LRUCache[T any] struct {
	Keys *list.List
	Data map[string]Node[T]
	Size int
}

func New[T any](size int) *LRUCache[T] {
	return &LRUCache[T]{
		Keys: list.New(),
		Data: make(map[string]Node[T]),
		Size: size,
	}
}

func (c *LRUCache[T]) Get(key string) (T, bool) {
	if node, ok := c.Data[key]; ok {
		c.Keys.MoveToFront(node.Element)
		return node.Data, true
	}
	var empty T
	return empty, false
}

func (c *LRUCache[T]) Put(key string, value T) {
	if node, ok := c.Data[key]; ok {
		node.Data = value
		c.Keys.MoveToFront(node.Element)
		return
	}

	if c.Keys.Len() == c.Size {
		last := c.Keys.Back()
		c.remove(last)
	}

	element := c.Keys.PushFront(key)
	c.Data[key] = Node[T]{Data: value, Element: element}
}

func (c *LRUCache[T]) Del(key string) {
	if node, ok := c.Data[key]; ok {
		c.remove(node.Element)
	}
}

func (c *LRUCache[T]) remove(element *list.Element) {
	if element == nil {
		return
	}
	c.Keys.Remove(element)
	delete(c.Data, element.Value.(string))
}
