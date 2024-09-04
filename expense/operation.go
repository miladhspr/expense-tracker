package expense

import (
	"fmt"
	"time"
)

func Add(filename string) error {
	description, amount, flagError := AddFlagsHandle()
	if flagError != nil {
		return fmt.Errorf("m Add 9: %v", flagError)
	}
	expenses, loadError := Load(filename)
	if loadError != nil {
		return fmt.Errorf("m Add 13: failed to load data db : %v", loadError)
	}
	newId := len(expenses) + 1
	newExpense := Expense{
		ID:        int64(newId),
		Amount:    *amount,
		Desc:      *description,
		CreatedAt: time.Now().Format(time.DateOnly),
	}
	expenses = append(expenses, newExpense)
	if err := Store(filename, expenses); err != nil {
		return fmt.Errorf("m Add: failed to store data into db : %v", err)
	}
	return nil
}

func Update(filename string) error {
	id, desc, amount, flagError := UpdateFlagsHandle()
	if flagError != nil {
		return flagError
	}
	expenses, loadErr := Load(filename)
	if loadErr != nil {
		return loadErr
	}
	for key, expense := range expenses {
		if expense.ID == *id {
			if *desc != "" {
				expense.Desc = *desc
			}
			if *amount != 0.0 {
				expense.Amount = *amount
			}
			expenses[key] = expense
		}

	}
	if storeErr := Store(filename, expenses); storeErr != nil {
		return storeErr
	}
	return nil
}
func Delete(filename string) error {
	id, err := DeleteFlagsHandle()
	if err != nil {
		return fmt.Errorf("m Delete : error of Flags : %v", err)
	}
	expenses, loadErr := Load(filename)
	if loadErr != nil {
		return loadErr
	}
	for key, expense := range expenses {
		if expense.ID == *id {
			expenses = append(expenses[:key], expenses[key+1:]...)
			break
		}
	}
	if storeErr := Store(filename, expenses); storeErr != nil {
		return storeErr
	}
	return nil
}
func List(filename string) error {
	expenses, loadErr := Load(filename)
	if loadErr != nil {
		return loadErr
	}
	fmt.Printf("# %-3s %-10s %-12s %s\n", "ID", "Date", "Description", "Amount")

	for _, expense := range expenses {
		fmt.Printf("# %-3d %-10s %-12s $%.2f\n", expense.ID, expense.CreatedAt, expense.Desc, expense.Amount)
	}
	return nil
}
func Summery(filename string) error {
	month, _ := SummeryFlagsHandle()
	var sum = 0.0
	expenses, loadErr := Load(filename)
	if loadErr != nil {
		return loadErr
	}

	if *month != 0 {
		m := time.Month(*month)
		for _, expense := range expenses {
			expenseDate, parseErr := time.Parse(time.DateOnly, expense.CreatedAt)
			if parseErr != nil {
				return fmt.Errorf("failed to parse date %v: %v", expense.CreatedAt, parseErr)
			}
			if expenseDate.Month() == m {
				sum += expense.Amount
			}
		}
		fmt.Printf("# Total expenses for %v : %0.1f", m, sum)
	} else {
		for _, expense := range expenses {
			sum += expense.Amount
		}
		fmt.Printf("# Total expenses: %0.1f", sum)
	}
	return nil
}
