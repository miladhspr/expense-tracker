# Expense Tracker

The Expense Tracker is a simple command-line application designed to help you manage your finances. With this tool, you can easily add, update, delete, and view your expenses. Additionally, the application allows you to summarize your expenses, either for all time or for a specific month of the current year.

## Features

### Core Functionality

1. **Add an Expense**
    - Users can add an expense by providing a description and an amount.
    - Example:
      ```
      $ expense-tracker add --description "Lunch" --amount 20
      # Expense added successfully (ID: 1)
      ```

2. **Update an Expense**
    - Users can update an existing expense by specifying the expense ID, a new description, and/or a new amount.
    - Example:
      ```
      $ expense-tracker update --id 1 --description "Brunch" --amount 25
      # Expense updated successfully
      ```

3. **Delete an Expense**
    - Users can delete an expense by providing the expense ID.
    - Example:
      ```
      $ expense-tracker delete --id 1
      # Expense deleted successfully
      ```

4. **View All Expenses**
    - Users can list all expenses in a formatted table.
    - Example:
      ```
      $ expense-tracker list
      # ID  Date       Description  Amount
      # 1   2024-08-06 Lunch        $20.00
      # 2   2024-08-06 Dinner       $10.00
      ```

5. **View Expense Summary**
    - Users can view the total amount of all expenses.
    - Example:
      ```
      $ expense-tracker summary
      # Total expenses: $30.00
      ```

6. **View Monthly Expense Summary**
    - Users can view the total amount of expenses for a specific month of the current year.
    - Example:
      ```
      $ expense-tracker summary --month 8
      # Total expenses for August: $20.00
      ```

### Additional Features (Optional Enhancements)

1. **Expense Categories**
    - Add categories to expenses and allow filtering expenses by category.

2. **Budgeting**
    - Set a monthly budget and alert the user when the total expenses exceed the budget.

3. **Exporting Data**
    - Export the list of expenses to a CSV file for easy sharing and analysis.

## Implementation Details

### Command-line Arguments
- The application uses command-line arguments to perform different operations. You can use any module suitable for argument parsing in your chosen programming language.

### Data Storage
- Expenses are stored in a simple text file (e.g., JSON or CSV format). This makes it easy to load, update, and save expense records.

### Error Handling
- The application includes error handling for various scenarios, such as invalid inputs, non-existent expense IDs, and other edge cases.

### Modular Code
- The code is organized into functions to ensure modularity. This makes the code easier to maintain, test, and extend with new features.

## Example Workflow

1. **Add Expenses**
    - Add multiple expenses using the `add` command.
   ```
   $ expense-tracker add --description "Lunch" --amount 20
   $ expense-tracker add --description "Dinner" --amount 10
   ```

2. **View All Expenses**
    - List all expenses to see a summary.
   ```
   $ expense-tracker list
   ```

3. **View Expense Summary**
    - Get the total amount spent so far.
   ```
   $ expense-tracker summary
   ```

4. **Update an Expense**
    - Modify an existing expense.
   ```
   $ expense-tracker update --id 1 --description "Brunch" --amount 25
   ```

5. **Delete an Expense**
    - Remove an expense from the record.
   ```
   $ expense-tracker delete --id 1
   ```

6. **View Monthly Summary**
    - Get a summary of expenses for a specific month.
   ```
   $ expense-tracker summary --month 8
   ```

## Conclusion

The Expense Tracker project is a practical tool for managing personal finances through the command line. It offers basic functionalities for tracking expenses, with room for enhancement through additional features like categorization, budgeting, and data export. This project is ideal for improving your skills in command-line programming, file handling, and data management.

By working on this project, you will gain experience in:
- Building CLI applications.
- Managing and organizing data in files.
- Handling user input and providing meaningful feedback.
- Writing modular and maintainable code.
- https://roadmap.sh/projects/expense-tracker