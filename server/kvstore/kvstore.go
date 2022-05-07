package kvstore

import (
	"sync"

	"github.com/jcbritobr/cstodo/model"
)

type KVStore struct {
	mutex sync.Mutex
	data  map[string]model.Item
}

func NewKVStore() *KVStore {
	return &KVStore{sync.Mutex{}, map[string]model.Item{}}
}

func (k *KVStore) Clear() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	for k2 := range k.data {
		delete(k.data, k2)
	}
}

func (k *KVStore) Len() int {
	return len(k.data)
}

func (k *KVStore) InsertItem(uuid string, item model.Item) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	k.data[uuid] = item
}

func (k *KVStore) LoadItem(uuid string) (model.Item, bool) {
	data, ok := k.data[uuid]
	return data, ok
}

func (k *KVStore) List() map[string]model.Item {
	return k.data
}

func (k *KVStore) DoneUndone(uuid string) bool {
	data, ok := k.data[uuid]
	if ok {
		data.Done = !data.Done
		k.data[uuid] = data
	}

	return ok
}
