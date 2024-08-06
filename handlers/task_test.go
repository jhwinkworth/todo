package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/models"
	"todo/store"
	"todo/test"
)

func getSimpleTask() models.Task {
	return models.Task{
		ID:          "1",
		Description: "do the dishes",
		Complete:    false,
	}
}

func TestCreateNewTask(t *testing.T) {
	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := NewTasksHandler(memStore)

	newTask := getSimpleTask()

	newTaskJSON, err := json.Marshal(newTask)
	test.AssertNil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewReader(newTaskJSON))
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	test.AssertEquals(t, res.StatusCode, 200)

	saved, _ := memStore.List()
	test.AssertEquals(t, len(saved), 1)

	got := saved["1"]

	test.AssertNil(t, err)
	test.AssertEquals(t, got, newTask)
}

func TestGetTasks(t *testing.T) {
	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := NewTasksHandler(memStore)

	newTask := getSimpleTask()
	newTaskMap := map[string]models.Task{"1": newTask}

	err := memStore.Add(newTask)
	test.AssertNil(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	test.AssertEquals(t, res.StatusCode, 200)

	got, err := io.ReadAll(res.Body)
	test.AssertNil(t, err)

	want, err := json.Marshal(newTaskMap)
	test.AssertDeepEquals(t, got, want)
}

func TestGetTasksByID(t *testing.T) {
	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := NewTasksHandler(memStore)

	newTask := getSimpleTask()

	err := memStore.Add(newTask)
	test.AssertNil(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/tasks/1", nil)
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	test.AssertEquals(t, res.StatusCode, 200)

	got, err := io.ReadAll(res.Body)
	test.AssertNil(t, err)

	want, err := json.Marshal(newTask)
	test.AssertDeepEquals(t, got, want)
}

func TestUpdateTask(t *testing.T) {
	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := NewTasksHandler(memStore)

	newTask := getSimpleTask()
	updatedTask := getSimpleTask()
	updatedTask.Complete = true

	err := memStore.Add(newTask)
	test.AssertNil(t, err)

	updatedTaskJSON, err := json.Marshal(updatedTask)
	test.AssertNil(t, err)

	req := httptest.NewRequest(http.MethodPut, "/api/tasks/1", bytes.NewReader(updatedTaskJSON))
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	test.AssertEquals(t, res.StatusCode, 200)

	saved, _ := memStore.List()
	got := saved["1"]

	test.AssertNil(t, err)
	test.AssertEquals(t, got, updatedTask)
}

func TestDeleteTask(t *testing.T) {
	// Create a MemStore and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := NewTasksHandler(memStore)

	newTask := getSimpleTask()

	err := memStore.Add(newTask)
	test.AssertNil(t, err)

	req := httptest.NewRequest(http.MethodDelete, "/api/tasks/1", nil)
	w := httptest.NewRecorder()
	tasksHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	test.AssertEquals(t, res.StatusCode, 200)

	saved, _ := memStore.List()
	test.AssertEquals(t, len(saved), 0)
}
