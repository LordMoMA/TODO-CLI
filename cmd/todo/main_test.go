package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	todoFile = ".todo.json"
)

func TestMain(m *testing.M) {
	// Build a test version of our command
	fmt.Println("Building tool...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build binary for %s: %s", binName, err)
		os.Exit(1)
	}
	// Run the tests
	fmt.Println("Running tests...")
	retCode := m.Run()
	// Clean up the test version of our command
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(todoFile)
	// Exit
	os.Exit(retCode)
}

func TestTodoCLI(t *testing.T) {
	task := "unit test task"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := fmt.Sprintf("%s/%s", dir, binName)
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
