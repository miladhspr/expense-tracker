package expense

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(filename string) ([]Expense, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v", err)
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get info of the file : %v", err)
	}
	var expenses []Expense
	if info.Size() <= 0 {
		expenses = []Expense{}
	} else {
		if err := json.NewDecoder(file).Decode(&expenses); err != nil {
			return nil, fmt.Errorf("failed to decode expenses : %v", err)
		}
	}
	return expenses, nil
}

func Store(filename string, expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal expenses: %v", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write data into file : %v", err)
	}
	return nil
}
