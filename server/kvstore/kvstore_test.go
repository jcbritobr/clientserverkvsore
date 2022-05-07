package kvstore

import (
	"testing"

	"github.com/jcbritobr/cstodo/model"
	"github.com/stretchr/testify/assert"
)

func TestNewKVStoreCantBeNil(t *testing.T) {
	store := NewKVStore()
	assert.NotNil(t, store, "The object store cant be nil")
}

func TestInsertItem(t *testing.T) {
	store := NewKVStore()
	store.InsertItem("1", model.Item{Title: "Item1", Description: "Simple Item", Done: false})
	assert.Equal(t, 1, store.Len(), "Len() must return 1, got %v", store.Len())
}

func TestLoadItemFromKVStore(t *testing.T) {
	store := NewKVStore()
	expected := model.Item{Title: "Item1", Description: "Simple Item", Done: false}
	store.InsertItem("1", expected)
	assert.Equal(t, 1, store.Len(), "Len() must return 1, got %v", store.Len())
	data, ok := store.LoadItem("1")
	assert.True(t, ok, "LoadItem() must return return ok=true, got %v", ok)
	assert.Equal(t, expected, data, "LoadItem() must return the item from store")
}

func TestLoadItemMustFail(t *testing.T) {
	store := NewKVStore()
	expected := model.Item{}
	data, ok := store.LoadItem("1")
	assert.False(t, ok, "LoadItem() must return false, got %v", ok)
	assert.Equal(t, expected, data, "LoadItem() must return empty Item, got %v", data)
}

func TestMustClearKVStoreData(t *testing.T) {
	store := NewKVStore()
	store.InsertItem("1", model.Item{})
	assert.Equal(t, 1, store.Len(), "Len() must return 1, got %v", store.Len())
	store.Clear()
	assert.Equal(t, 0, store.Len(), "Len() must return 0, got %v", store.Len())
}

func TestDoneundoneMustChangeInvertDoneData(t *testing.T) {
	store := NewKVStore()
	store.InsertItem("1", model.Item{Done: false})
	result := store.DoneUndone("1")
	assert.True(t, result, "result must be true. Got %v", result)
	data, _ := store.LoadItem("1")
	assert.True(t, data.Done, "data.Done must be true. Got %v", data.Done)
}
