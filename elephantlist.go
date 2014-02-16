package elephantlist

import (
	"math/rand"
)

const (
	growRate       = 0.2
	maxLevel       = 24
	maxNodeKeySize = 32
)

type ElephantList struct {
	head     *node
	lessThan func(l, r interface{}) bool
}

type node struct {
	next   []*node
	keys   []interface{}
	values []interface{}
}

func NewList(lessThan func(l, r interface{}) bool) *ElephantList {
	return &ElephantList{
		head:     &node{},
		lessThan: lessThan,
	}
}

func NewIntList() *ElephantList {
	return NewList(func(l, r interface{}) bool {
		return l.(int) < r.(int)
	})
}

func (e *ElephantList) Get(key interface{}) (value interface{}, ok bool) {
	nd := e.selectNode(key)
	return e.get(nd, key)
}

func (e *ElephantList) Set(key, value interface{}) {
	nd := e.selectNode(key)
	e.set(nd, key, value)
	if len(nd.keys) >= maxNodeKeySize {
		e.split(nd)
	}
}

func (e *ElephantList) selectNode(key interface{}) *node {
	cur := e.head
	for i := len(cur.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && !e.lessThan(key, cur.next[i].keys[0]) {
			cur = cur.next[i]
		}
	}
	return cur
}

func (e *ElephantList) split(left *node) {
	rightSize := len(left.keys) / 2
	cut := len(left.keys) - rightSize

	right := &node{
		next:   make([]*node, randomLevel()),
		keys:   make([]interface{}, rightSize),
		values: make([]interface{}, rightSize),
	}

	copy(right.keys, left.keys[cut:])
	left.keys = left.keys[:cut]
	copy(right.values, left.values[cut:])
	left.values = left.values[:cut]

	e.fixLinks(left, right)
}

func (e *ElephantList) fixLinks(left, right *node) {
	if len(right.next) <= len(left.next) {
		for i := range right.next {
			right.next[i] = left.next[i]
			left.next[i] = right
		}
	} else {
		for i := len(e.head.next); i < len(right.next); i++ {
			e.head.next = append(e.head.next, nil)
		}

		cur := e.head
		for i := len(right.next) - 1; i >= 0; i-- {
			for cur.next[i] != nil && e.lessThan(cur.next[i].keys[0], right.keys[0]) {
				cur = cur.next[i]
			}
			right.next[i] = cur.next[i]
			cur.next[i] = right
		}
	}
}

func (e *ElephantList) get(nd *node, key interface{}) (value interface{}, ok bool) {
	for i, k := range nd.keys {
		switch {
		case e.lessThan(k, key):
			continue
		case e.lessThan(key, k):
			return nil, false
		default:
			return nd.values[i], true
		}
	}
	return nil, false
}

func (e *ElephantList) set(nd *node, key, value interface{}) {
	for i, k := range nd.keys {
		if !e.lessThan(k, key) {
			e.insert(nd, i, key, value)
			return
		}
	}

	nd.keys = append(nd.keys, key)
	nd.values = append(nd.values, value)
}

func (e *ElephantList) insert(nd *node, i int, key, value interface{}) {
	if e.lessThan(key, nd.keys[i]) {
		nd.keys = append(nd.keys, nil)
		copy(nd.keys[i+1:], nd.keys[i:])
		nd.keys[i] = key

		nd.values = append(nd.values, nil)
		copy(nd.values[i+1:], nd.values[i:])
		nd.values[i] = value
	} else {
		nd.values[i] = value
	}
}

func randomLevel() (level int) {
	for level = 1; rand.Float64() < growRate && level < maxLevel; level++ {
	}
	return
}
