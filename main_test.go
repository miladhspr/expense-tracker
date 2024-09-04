package main

import (
	"expense-tracker/expense"
	"os"
	"testing"
)

func TestLoadEmptyFile(t *testing.T) {
	// Create a temporary empty file
	tempFile, err := os.CreateTemp("", "temp-expense.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Test load function with the empty file
	expenses, err := expense.Load(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(expenses) != 0 {
		t.Fatalf("Expected empty slice, got %v", expenses)
	}
}

func TestStore(t *testing.T) {
	tempFile, err := os.CreateTemp("", "temp-expense.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	expenses := []expense.Expense{
		{
			ID:        1,
			Amount:    100.0,
			Desc:      "Groceries",
			CreatedAt: "2023-09-01",
		},
	}

	// Test store function
	if err := expense.Store(tempFile.Name(), expenses); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Reload the data to check if it was stored correctly
	loadedExpenses, err := expense.Load(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(loadedExpenses) != 1 || loadedExpenses[0].Desc != "Groceries" {
		t.Fatalf("Data was not stored correctly")
	}
}

func TestAddFlagsHandle(t *testing.T) {
	os.Args = []string{"expense-tracker", "add", "--description", "Lunch", "--amount", "15.0"}

	desc, amount, flagError := expense.AddFlagsHandle()
	if flagError != nil {
		t.Fatalf("Expected error on flag requires")
	}
	if *desc != "Lunch" {
		t.Fatalf("Expected description to be 'Lunch', got %v", *desc)
	}
	if *amount != 15.0 {
		t.Fatalf("Expected amount to be 15.0, got %v", *amount)
	}
}

func TestUpdateFlagsHandle(t *testing.T) {
	os.Args = []string{"expense-tracker", "update", "--id", "1", "--description", "Dinner"}

	id, desc, amount, err := expense.UpdateFlagsHandle()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if *id != 1 {
		t.Fatalf("Expected id to be 1, got %v", *id)
	}
	if *desc != "Dinner" {
		t.Fatalf("Expected description to be 'Dinner', got %v", *desc)
	}
	if *amount != 0.0 {
		t.Fatalf("Expected amount to be 0.0, got %v", *amount)
	}
}

func TestAddFunctionality(t *testing.T) {
	tempFile, createTempFileError := os.CreateTemp("", "temp-expense.txt")
	if createTempFileError != nil {
		t.Fatalf("Failed to create temp file: %v", createTempFileError)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	os.Args = []string{"expense-tracker", "add", "--description", "Books", "--amount", "50.0"}
	err := expense.Add(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expenses, err := expense.Load(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(expenses) != 1 || expenses[0].Desc != "Books" || expenses[0].Amount != 50.0 {
		t.Fatalf("Failed to add expense correctly")
	}
}

func TestUpdateFunctionality(t *testing.T) {
	tempFile, createTempFileError := os.CreateTemp("", "temp-expense.txt")
	if createTempFileError != nil {
		t.Fatalf("Failed to create temp file: %v", createTempFileError)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	os.Args = []string{"expense-tracker", "add", "--description", "Books", "--amount", "50.0"}
	err := expense.Add(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	os.Args = []string{"expense-tracker", "update", "--id", "1", "--description", "Updated Expense", "--amount", "45.2"}
	if err := expense.Update(tempFile.Name()); err != nil {
		t.Fatalf("Failed to Update : %v", err)
	}
	expenses, err := expense.Load(tempFile.Name())

	if err != nil {
		t.Fatalf("Failed to Load data: %v", err)
	}
	if len(expenses) != 1 || expenses[0].Desc != "Updated Expense" || expenses[0].Amount != 45.2 {
		t.Fatalf("ex: %v", expenses)
	}
}

func TestDeleteFlagsHandle(t *testing.T) {
	os.Args = []string{"expense-tracker", "delete", "--id", "1"}

	id, err := expense.DeleteFlagsHandle()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if *id != 1 {
		t.Fatalf("Expected id to be 1, got %v", *id)
	}
}

func TestDeleteFunctionality(t *testing.T) {
	tempFile, createTempFileError := os.CreateTemp("", "temp-expense.txt")
	if createTempFileError != nil {
		t.Fatalf("Failed to create temp file: %v", createTempFileError)
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	os.Args = []string{"expense-tracker", "add", "--description", "Books", "--amount", "50.0"}
	err := expense.Add(tempFile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	os.Args = []string{"expense-tracker", "delete", "--id", "1"}
	if err := expense.Delete(tempFile.Name()); err != nil {
		t.Fatalf("Failed to Update : %v", err)
	}
	expenses, err := expense.Load(tempFile.Name())

	if err != nil {
		t.Fatalf("Failed to Load data: %v", err)
	}
	if len(expenses) != 0 {
		t.Fatalf("assert to delete and file should be empty but : %v", expenses)
	}
}
