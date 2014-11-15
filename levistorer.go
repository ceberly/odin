package odin

import (
	"github.com/jmhodges/levigo"
)

import "log"

type LeviStorer struct {
	ro *levigo.ReadOptions
	wo *levigo.WriteOptions
	db *levigo.DB
}

func NewLeviStorer() *LeviStorer {
	opts := levigo.NewOptions()
	opts.SetCreateIfMissing(true)

	db, err := levigo.Open("tmp/db", opts)
	if err != nil {
		log.Fatalln(err)
	}

	ro := levigo.NewReadOptions()
	wo := levigo.NewWriteOptions()

	l := LeviStorer{ro, wo, db}

	//defer ro.Close()
	//defer wo.Close()

	return &l
}

// This method is potentially slow.
func (l *LeviStorer) Count() int {
	it := l.db.NewIterator(l.ro)
	defer it.Close()

	c := 0
	for it.SeekToFirst(); it.Valid(); it.Next() {
		c += 1
	}

	return c
}

func (l *LeviStorer) Get(key []byte) ([]byte, error) {
	return l.db.Get(l.ro, key)
}

func (l *LeviStorer) Put(key, value []byte) error {
	return l.db.Put(l.wo, key, value)
}

func (l *LeviStorer) Del(key []byte) error {
	return l.db.Delete(l.wo, key)
}
