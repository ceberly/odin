package odin

import "testing"

import (
	"bytes"
)

type key []byte
type value []byte

type keyValue struct {
	key   key
	value value
}

func TestLevistorer(t *testing.T) {
	l := NewLeviStorer()

	tests := []struct {
		put   []keyValue
		del   []key
		get   []keyValue
		count int
	}{
		{
			[]keyValue{
				{[]byte("key1"), []byte("value1")},
				{[]byte("key2"), []byte("value2")},
				{[]byte("key3"), []byte("value3")},
			},
			[]key{
				[]byte("key2"),
			},
			[]keyValue{
				{[]byte("key1"), []byte("value1")},
				{[]byte("key3"), []byte("value3")},
			},
			2,
		},
	}

	for i, tt := range tests {
		for _, p := range tt.put {
			l.Put(p.key, p.value)
		}

		for _, d := range tt.del {
			l.Del(d)
		}

		c := l.Count()
		if c != tt.count {
			t.Errorf("%d. Expected count to equal %d, got %d instead.", i, tt.count, c)
		}

		for j, g := range tt.get {

			v, _ := l.Get(g.key)
			if !bytes.Equal(v, g.value) {
				t.Errorf("%d.%d Expected '%x' to equal '%x'\n", i, j, v, g.value)
			}
		}
	}
}
