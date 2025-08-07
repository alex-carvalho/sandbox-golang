package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

type TodoApp struct {
	tasks []Task
}

func NewTodoApp() *TodoApp {
	return &TodoApp{tasks: []Task{}}
}

func (app *TodoApp) nextID() int {
	return len(app.tasks) + 1
}



func (app *TodoApp) create(title string) *Task {
	task := Task{
		ID:        app.nextID(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	app.tasks = append(app.tasks, task)
	return &task
}

func (app *TodoApp) complete(id int) error {
	for i := range app.tasks {
		if app.tasks[i].ID == id {
			app.tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (app *TodoApp) remove(id int) error {
	for i, task := range app.tasks {
		if task.ID == id {
			app.tasks = append(app.tasks[:i], app.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (app *TodoApp) getTask(id int) (*Task, error) {
	for _, task := range app.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task %d not found", id)
}

func (app *TodoApp) getTasks(filter string) []Task {
	var filtered []Task
	for _, task := range app.tasks {
		if filter == "completed" && !task.Completed {
			continue
		}
		if filter == "pending" && task.Completed {
			continue
		}
		filtered = append(filtered, task)
	}
	return filtered
}

func (app *TodoApp) showHelp() {
	fmt.Println("\n📝 Todo CLI Commands:")
	fmt.Println("  list [completed|pending] - List tasks")
	fmt.Println("  create <title>           - Create new task")
	fmt.Println("  complete <id>            - Mark task as complete")
	fmt.Println("  remove <id>              - Remove task")
	fmt.Println("  details <id>             - Show task details")
	fmt.Println("  help                     - Show this help")
	fmt.Println("  exit                     - Exit application")
}

func (app *TodoApp) processCommand(input string) bool {
	parts := strings.Fields(strings.TrimSpace(input))
	if len(parts) == 0 {
		return true
	}
	
	command := parts[0]
	
	switch command {
	case "list":
		filter := ""
		if len(parts) > 1 {
			filter = parts[1]
		}
		tasks := app.getTasks(filter)
		if len(tasks) == 0 {
			fmt.Println("📭 No tasks found")
		} else {
			fmt.Println("\n📋 Tasks:")
			for _, task := range tasks {
				status := "⭕"
				if task.Completed {
					status = "✅"
				}
				fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
			}
		}
		
	case "create":
		if len(parts) < 2 {
			fmt.Println("Usage: create <title>")
			return true
		}
		title := strings.Join(parts[1:], " ")
		task := app.create(title)
		fmt.Printf("✓ Created task: %s\n", task.Title)
		
	case "complete":
		if len(parts) < 2 {
			fmt.Println("Usage: complete <id>")
			return true
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("❌ Invalid task ID")
			return true
		}
		if err := app.complete(id); err != nil {
			fmt.Printf("❌ %v\n", err)
		} else {
			fmt.Printf("✓ Task completed\n")
		}
		
	case "remove":
		if len(parts) < 2 {
			fmt.Println("Usage: remove <id>")
			return true
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("❌ Invalid task ID")
			return true
		}
		if err := app.remove(id); err != nil {
			fmt.Printf("❌ %v\n", err)
		} else {
			fmt.Printf("🗑️  Task removed\n")
		}
		
	case "details":
		if len(parts) < 2 {
			fmt.Println("Usage: details <id>")
			return true
		}
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("❌ Invalid task ID")
			return true
		}
		task, err := app.getTask(id)
		if err != nil {
			fmt.Printf("❌ %v\n", err)
		} else {
			status := "Pending"
			if task.Completed {
				status = "Completed"
			}
			fmt.Printf("ID: %d\n", task.ID)
			fmt.Printf("Title: %s\n", task.Title)
			fmt.Printf("Status: %s\n", status)
			fmt.Printf("Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
		}
		
	case "help":
		app.showHelp()
		
	case "exit":
		fmt.Println("👋 Goodbye!")
		return false
		
	default:
		fmt.Printf("❌ Unknown command: %s (type 'help' for commands)\n", command)
	}
	
	return true
}

func main() {
	app := NewTodoApp()
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("📝 Welcome to Todo CLI!")
	app.showHelp()
	
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		
		if !app.processCommand(scanner.Text()) {
			break
		}
	}
}