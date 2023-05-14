// Package counter implements a [Counter] for comparable items.
package counter

import (
	"fmt"
	"sort"
)

// An ItemCount contains an item and its count.
type ItemCount[T comparable] struct {
	Item  T
	Count int
}

// A Counter keeps track of the counts of items (which must be comparable).
type Counter[T comparable] struct {
	data map[T]int
}

// New returns a new Counter.
func New[T comparable](items ...T) *Counter[T] {
	c := Counter[T]{data: make(map[T]int)}
	c.Update(items...)
	return &c
}

// FromMap returns a new Counter initialized from a map.
func FromMap[T comparable](m map[T]int) *Counter[T] {
	c := Counter[T]{data: make(map[T]int, len(m))}
	for k, v := range m {
		c.data[k] = v
	}
	return &c
}

// Add adds an item and returns the new item count.
func (c *Counter[T]) Add(item T) int {
	if cnt, ok := c.data[item]; ok {
		c.data[item] = cnt + 1
	} else {
		c.data[item] = 1
	}
	return c.data[item]
}

// Sub subtracts an item and returns the new item count.
func (c *Counter[T]) Sub(item T) int {
	if cnt, ok := c.data[item]; ok {
		c.data[item] = cnt - 1
	} else {
		c.data[item] = -1
	}
	return c.data[item]
}

// Remove removes an item from the Counter.
func (c *Counter[T]) Remove(item T) bool {
	if _, ok := c.data[item]; ok {
		delete(c.data, item)
		return true
	}
	return false
}

// Update adds items to the Counter.
func (c *Counter[T]) Update(items ...T) {
	for _, item := range items {
		c.Add(item)
	}
}

// Get returns the count for an item.
func (c *Counter[T]) Get(item T) int {
	return c.data[item]
}

// Contains returns true if the item is in the Counter (even when the count is 0).
func (c *Counter[T]) Contains(item T) bool {
	_, ok := c.data[item]
	return ok
}

// Len returns the number of items.
func (c *Counter[T]) Len() int {
	return len(c.data)
}

// Total returns the sum of the item counts.
func (c *Counter[T]) Total() int {
	total := 0
	for _, cnt := range c.data {
		total += cnt
	}
	return total
}

// MostCommon returns n items orderd by count (descending); n <= 0: all items.
func (c *Counter[T]) MostCommon(n int) []ItemCount[T] {
	result := make([]ItemCount[T], 0, len(c.data))
	for k, v := range c.data {
		result = append(result, ItemCount[T]{Item: k, Count: v})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Count > result[j].Count })
	if n > 0 && n < len(c.data) {
		result = result[:n]
	}
	return result
}

// Items returns a slice with items where each item is repeated according to their count.
// Items with negative count are ignored.
func (c *Counter[T]) Items() []T {
	result := make([]T, 0, c.Total())
	for _, ic := range c.MostCommon(0) {
		for i := 0; i < ic.Count; i++ {
			result = append(result, ic.Item)
		}
	}
	return result
}

// Map returns a map with items as keys and counts as values.
func (c *Counter[T]) Map() map[T]int {
	m := make(map[T]int, len(c.data))
	for k, v := range c.data {
		m[k] = v
	}
	return m
}

// Clone clones the Counter.
func (c *Counter[T]) Clone() *Counter[T] {
	return &Counter[T]{data: c.Map()}
}

// String returns a string representation of the Counter.
func (c *Counter[T]) String() string {
	return fmt.Sprintf("Counter{Items: %d, Total: %d}", len(c.data), c.Total())
}
