package main

import (
	"testing"
	"time"
)

func TestNewTodoApp(t *testing.T) {
	app := NewTodoApp()
	if app == nil {
		t.Fatal("NewTodoApp() returned nil")
	}
	if len(app.tasks) != 0 {
		t.Errorf("Expected empty tasks, got %d tasks", len(app.tasks))
	}
}

func TestCreateTask(t *testing.T) {
	app := NewTodoApp()
	
	task := app.create("Test task")
	
	if task == nil {
		t.Fatal("create() returned nil")
	}
	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}
	if task.Title != "Test task" {
		t.Errorf("Expected title 'Test task', got '%s'", task.Title)
	}
	if task.Completed {
		t.Error("Expected task to be incomplete")
	}
	if len(app.tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(app.tasks))
	}
}

func TestNextID(t *testing.T) {
	app := NewTodoApp()
	
	if app.nextID() != 1 {
		t.Errorf("Expected first ID to be 1, got %d", app.nextID())
	}
	
	app.create("Task 1")
	if app.nextID() != 2 {
		t.Errorf("Expected second ID to be 2, got %d", app.nextID())
	}
	
	app.create("Task 2")
	if app.nextID() != 3 {
		t.Errorf("Expected third ID to be 3, got %d", app.nextID())
	}
}

func TestCompleteTask(t *testing.T) {
	app := NewTodoApp()
	task := app.create("Test task")
	
	err := app.complete(task.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if !app.tasks[0].Completed {
		t.Error("Task should be completed")
	}
}

func TestCompleteNonExistentTask(t *testing.T) {
	app := NewTodoApp()
	
	err := app.complete(999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
	
	expectedMsg := "task 999 not found"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestRemoveTask(t *testing.T) {
	app := NewTodoApp()
	task := app.create("Test task")
	
	err := app.remove(task.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if len(app.tasks) != 0 {
		t.Errorf("Expected 0 tasks after removal, got %d", len(app.tasks))
	}
}

func TestRemoveNonExistentTask(t *testing.T) {
	app := NewTodoApp()
	
	err := app.remove(999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}

func TestGetTask(t *testing.T) {
	app := NewTodoApp()
	created := app.create("Test task")
	
	task, err := app.getTask(created.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if task.Title != "Test task" {
		t.Errorf("Expected title 'Test task', got '%s'", task.Title)
	}
}

func TestGetNonExistentTask(t *testing.T) {
	app := NewTodoApp()
	
	task, err := app.getTask(999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
	if task != nil {
		t.Error("Expected nil task for non-existent ID")
	}
}

func TestGetTasksNoFilter(t *testing.T) {
	app := NewTodoApp()
	app.create("Task 1")
	app.create("Task 2")
	app.complete(1)
	
	tasks := app.getTasks("")
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestGetTasksCompletedFilter(t *testing.T) {
	app := NewTodoApp()
	app.create("Task 1")
	app.create("Task 2")
	app.complete(1)
	
	tasks := app.getTasks("completed")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 completed task, got %d", len(tasks))
	}
	if !tasks[0].Completed {
		t.Error("Filtered task should be completed")
	}
}

func TestGetTasksPendingFilter(t *testing.T) {
	app := NewTodoApp()
	app.create("Task 1")
	app.create("Task 2")
	app.complete(1)
	
	tasks := app.getTasks("pending")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 pending task, got %d", len(tasks))
	}
	if tasks[0].Completed {
		t.Error("Filtered task should be pending")
	}
}

func TestTaskCreationTime(t *testing.T) {
	app := NewTodoApp()
	before := time.Now()
	task := app.create("Test task")
	after := time.Now()
	
	if task.CreatedAt.Before(before) || task.CreatedAt.After(after) {
		t.Error("Task creation time should be between before and after timestamps")
	}
}

func TestMultipleTasksIDSequence(t *testing.T) {
	app := NewTodoApp()
	
	task1 := app.create("Task 1")
	task2 := app.create("Task 2")
	task3 := app.create("Task 3")
	
	if task1.ID != 1 || task2.ID != 2 || task3.ID != 3 {
		t.Errorf("Expected IDs 1,2,3 got %d,%d,%d", task1.ID, task2.ID, task3.ID)
	}
}

func TestRemoveMiddleTask(t *testing.T) {
	app := NewTodoApp()
	
	app.create("Task 1")
	app.create("Task 2")
	app.create("Task 3")
	
	// Remove middle task
	err := app.remove(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if len(app.tasks) != 2 {
		t.Errorf("Expected 2 tasks after removal, got %d", len(app.tasks))
	}
	
	// Verify remaining tasks
	tasks := app.getTasks("")
	if tasks[0].ID != 1 || tasks[1].ID != 3 {
		t.Errorf("Expected remaining task IDs 1,3 got %d,%d", tasks[0].ID, tasks[1].ID)
	}
}