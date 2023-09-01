package webapp

import (
	"sync"
)

type Db struct {
	cycles *CyclesDb
}

func InitDb() *Db {
	return &Db{
		cycles: &CyclesDb{
			m:  map[int]Cycle{},
			mu: &sync.RWMutex{},
		},
	}
}

type CyclesDb struct {
	m  map[int]Cycle
	mu *sync.RWMutex
}

func (d *CyclesDb) append(c Cycle) {
	d.mu.Lock()
	defer d.mu.Unlock()

	size := len(d.m)
	d.m[size] = c
}

func (d *CyclesDb) all() map[int]Cycle {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.m
}
