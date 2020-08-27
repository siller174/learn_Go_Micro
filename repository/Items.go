package repository

import (
	proto "grpcProject/proto/crud"
	"sync"
)

type Items struct {
	mx    *sync.Mutex
	items map[string]*proto.Employee
}

func NewItems() *Items {
	return &Items{
		items: make(map[string]*proto.Employee),
		mx:    &sync.Mutex{},
	}
}

func (items *Items) Put(id string, emp *proto.Employee) {
	items.mx.Lock()
	defer items.mx.Unlock()
	items.items[id] = emp
}

func (items *Items) Get(id string) *proto.Employee {
	items.mx.Lock()
	defer items.mx.Unlock()
	val, _ := items.items[id]
	return val
}

func (items *Items) Detele(id string) {
	items.mx.Lock()
	defer items.mx.Unlock()
	delete(items.items, id)
}
