package test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/handlers"
	"todo/store"
)

func TestTasksHandlerCRUD_Integration(t *testing.T) {

	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := handlers.NewTasksHandler(memStore)

	// Test data
	var newTaskJSON = []byte(`{"id":"1","description":"this is a task","complete":false}`)
	var updatedTaskJSON = []byte(`{"id":"1","description":"this is a task","complete":true}`)

	// CREATE - add a new task
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(newTaskJSON))
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	AssertEquals(t, res.StatusCode, 200)

	saved, _ := memStore.List()
	AssertEquals(t, len(saved), 1)

	// GET - find the task just added
	req = httptest.NewRequest(http.MethodGet, "/api/tasks/1", nil)
	w = httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	AssertEquals(t, res.StatusCode, 200)

	got, err := io.ReadAll(res.Body)

	AssertNil(t, err)
	AssertDeepEquals(t, got, newTaskJSON)

	// UPDATE - complete task
	req = httptest.NewRequest(http.MethodPut, "/api/tasks/1", bytes.NewBuffer(updatedTaskJSON))
	w = httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	AssertEquals(t, res.StatusCode, 200)

	gotUpdatedTask, err := memStore.Get("1")
	AssertNil(t, err)

	AssertTrue(t, gotUpdatedTask.Complete)

	// DELETE - remove task
	req = httptest.NewRequest(http.MethodDelete, "/api/tasks/1", nil)
	w = httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	AssertEquals(t, res.StatusCode, 200)

	saved, _ = memStore.List()
	AssertEquals(t, len(saved), 0)
}
