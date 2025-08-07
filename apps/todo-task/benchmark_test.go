package main

import (
	"fmt"
	"testing"
)

func BenchmarkCreateTask(b *testing.B) {
	app := NewTodoApp()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		app.create(fmt.Sprintf("Task %d", i))
	}
}

func BenchmarkCompleteTask(b *testing.B) {
	app := NewTodoApp()
	// Pre-create tasks
	for i := 0; i < b.N; i++ {
		app.create(fmt.Sprintf("Task %d", i))
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.complete(i + 1)
	}
}

func BenchmarkGetTasks(b *testing.B) {
	app := NewTodoApp()
	// Pre-create 1000 tasks
	for i := 0; i < 1000; i++ {
		task := app.create(fmt.Sprintf("Task %d", i))
		if i%2 == 0 {
			app.complete(task.ID)
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.getTasks("")
	}
}

func BenchmarkGetTasksFiltered(b *testing.B) {
	app := NewTodoApp()
	// Pre-create 1000 tasks
	for i := 0; i < 1000; i++ {
		task := app.create(fmt.Sprintf("Task %d", i))
		if i%2 == 0 {
			app.complete(task.ID)
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.getTasks("completed")
	}
}