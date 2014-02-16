package elephantlist

import (
	"fmt"
	"math/rand"
	"testing"
)

func assert(t *testing.T, condition bool, args ...interface{}) {
	if !condition {
		t.Fatal(args)
	}
}

func dumpList(e *ElephantList) {
	fmt.Println("--------")
	fmt.Printf("%p => %+v\n", e.head, e.head)
	for cur := e.head; cur.next != nil; {
		cur = cur.next[0]
		if cur == nil {
			break
		}
		fmt.Printf("%p => %+v\n", cur, cur)
	}
}

func TestNew(t *testing.T) {
	list := NewIntList()
	assert(t, list.head != nil)
}

func TestInternalSet(t *testing.T) {
	list := NewIntList()

	list.Set(3, "three")
	assert(t, len(list.head.keys) == 1)

	list.Set(4, "???")
	assert(t, len(list.head.keys) == 2)
	list.Set(4, "four")
	assert(t, len(list.head.keys) == 2)

	list.Set(1, "one")
	assert(t, len(list.head.keys) == 3)

	list.Set(2, "two")
	assert(t, len(list.head.keys) == 4)
	assert(t, len(list.head.values) == 4)

	var (
		value interface{}
		ok    bool
	)

	value, ok = list.Get(0)
	assert(t, !ok)

	value, ok = list.Get(1)
	assert(t, ok)
	assert(t, value.(string) == "one")

	value, ok = list.Get(2)
	assert(t, ok)
	assert(t, value.(string) == "two")

	value, ok = list.Get(3)
	assert(t, ok)
	assert(t, value.(string) == "three")

	value, ok = list.Get(4)
	assert(t, ok)
	assert(t, value.(string) == "four")
}

func TestSequentialSet(t *testing.T) {
	list := NewIntList()
	for i := 0; i < 10000; i += 2 {
		list.Set(i, i+1)
	}
	for i := 0; i < 10000; i++ {
		value, ok := list.Get(i)

		if i%2 == 0 {
			assert(t, ok)
			assert(t, value.(int) == i+1)
		} else {
			assert(t, !ok)
		}
	}
}

func TestRandomSet(t *testing.T) {
	list := NewIntList()
	n := 10000
	perm := rand.Perm(n)
	for i := 0; i < n/2; i++ {
		list.Set(perm[i], perm[i]+1)
		//dumpList(list)
	}

	for i := 0; i < n; i++ {
		key := perm[i]
		value, ok := list.Get(key)

		if i < n/2 {
			assert(t, ok, "key not found:", key)
			assert(t, value.(int) == key+1)
		} else {
			assert(t, !ok, "found strange key:", key)
		}
	}
}
