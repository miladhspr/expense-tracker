package expense

import (
	"flag"
	"fmt"
	"os"
)

func AddFlagsHandle() (*string, *float64, error) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// Define the flags within this flag set
	description := addCmd.String("description", "", "Description of the expense")
	amount := addCmd.Float64("amount", 0.0, "Amount of the expense")

	// Parse the flags starting from the 3rd argument
	addCmd.Parse(os.Args[2:])

	// Ensure the required flags are provided
	if *description == "" || *amount == 0.0 {
		return nil, nil, fmt.Errorf("usage: expense-tracker add --description <description> --amount <amount>")
	}
	return description, amount, nil
}

func UpdateFlagsHandle() (*int64, *string, *float64, error) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	// Define the flags within this flag set
	id := updateCmd.Int64("id", 0, "ID expense")
	description := updateCmd.String("description", "", "Description of the expense")
	amount := updateCmd.Float64("amount", 0.0, "Amount of the expense")

	// Parse the flags starting from the 3rd argument
	updateCmd.Parse(os.Args[2:])
	if *id == 0 {
		return nil, nil, nil, fmt.Errorf("usage: expense-tracker update --id <expense id>")
	}
	// Ensure one of them should be entered
	// required flags are provided
	if *description == "" && *amount == 0.0 {
		return nil, nil, nil, fmt.Errorf("usage: expense-tracker update --description <description> --amount <amount>")
	}
	return id, description, amount, nil
}

func DeleteFlagsHandle() (*int64, error) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	// Define the flags within this flag set
	id := deleteCmd.Int64("id", 0, "ID expense")

	// Parse the flags starting from the 3rd argument
	deleteCmd.Parse(os.Args[2:])
	if *id == 0 {
		return nil, fmt.Errorf("usage: expense-tracker update --id <expense id>")
	}
	return id, nil
}
func SummeryFlagsHandle() (*int64, error) {
	summaryCmd := flag.NewFlagSet("month", flag.ExitOnError)

	// Define the flags within this flag set
	month := summaryCmd.Int64("month", 0, "Total amount of specific month")

	// Parse the flags starting from the 3rd argument
	summaryCmd.Parse(os.Args[2:])

	return month, nil
}
