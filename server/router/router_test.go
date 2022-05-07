package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jcbritobr/cstodo/model"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	e.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestPostInsertItem(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/insert", strings.NewReader(`{"title":"bake bread", "description":"we need bake bread", "done":false}`))
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	store.Clear()
}

func TestPostInsertItemMustFail(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/insert", strings.NewReader(`{`))
	req.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMustLoadItemFromKVStore(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	store.InsertItem("1", model.Item{Title: "first", Description: "first", Done: false})
	req, _ := http.NewRequest("GET", "/load?uuid=1", nil)
	assert.Equal(t, 1, store.Len(), "Len() must have value 1, got %v", store.Len())
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "w.code must have status 200, got %c", w.Code)
	expected := model.Item{Title: "first", Description: "first", Done: false}
	var item model.Item
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&item)
	assert.Equal(t, expected, item, "w.Body must have the same data of expected, got %v", item)
	store.Clear()
}

func TestLoadItemFromKVStoreMustFail(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/load?uuid=1", nil)
	assert.Equal(t, 0, store.Len(), "Len() must return 0. Got %v", store.Len())
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "w.Code must have status 200, Got %v", w.Code)
	var jsonErr model.ErrorMessage
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&jsonErr)
	assert.Equal(t, model.ErrorMessage{Message: model.ErrNotFound}, jsonErr, "message must have not found. Got %v", jsonErr)
}

func TestLoadItemFromKVStoreMustFailWithEmptyError(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/load?uuid", nil)
	assert.Equal(t, 0, store.Len(), "Len() must return 0. Got %v", store.Len())
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "w.Code must have status 200, Got %v", w.Code)
	var jsonErr model.ErrorMessage
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&jsonErr)
	assert.Equal(t, model.ErrorMessage{Message: model.ErrEmptyQuery}, jsonErr, "error must be ErrEmptyQuery. Got %v", jsonErr)
}

func TestListDataFromKVStore(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/list", nil)
	first := model.Item{Title: "first", Description: "first", Done: false}
	second := model.Item{Title: "second", Description: "second", Done: false}
	third := model.Item{Title: "third", Description: "third", Done: false}
	store.InsertItem("1", first)
	store.InsertItem("2", second)
	store.InsertItem("3", third)

	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "w.Code must be 200. Got %d", w.Code)
	var data map[string]model.Item
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&data)
	assert.Equal(t, 3, len(data), "len(data) must return 3. Got %v", len(data))

	store.Clear()
}

func TestDoneundoneMustSwitchItemDoneValue(t *testing.T) {
	e := Boostrap()
	w := httptest.NewRecorder()
	store.InsertItem("1", model.Item{Done: false})
	buffer := bytes.NewBufferString("")
	encoder := json.NewEncoder(buffer)
	dd := model.UuidMessage{Uuid: "1"}
	err := encoder.Encode(dd)
	assert.Nil(t, err, "encoder should not fail. Got %v", err)
	req, _ := http.NewRequest("POST", "/doneundone", buffer)
	e.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "w.Code must have status 200. Got %v", w.Code)

	var data model.DoneundoneResult
	decoder := json.NewDecoder(w.Body)
	decoder.Decode(&data)

	assert.True(t, data.Done, "data.Done must be true. Got %v", data.Done)

	store.Clear()
}
