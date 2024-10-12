package services

import (
	"encoding/json"
	"errors"
	"github.com/anilonayy/go-examples/expense-tracker/internal/models"
	"os"
	"time"
)

const fileName = "expenses.json"

func AddExpense(e *models.Expense) error {
	expenses, err := GetExpensesFromFile()
	if err != nil {
		return err
	}

	e.ID = int(time.Now().Unix()) + len(expenses) + 1

	expenses = append(expenses, *e)

	return SaveExpenses(expenses)
}

func SaveExpenses(expenses []models.Expense) error {
	data, err := json.Marshal(expenses)
	if err != nil {
		return errors.New("error marshalling data")
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return errors.New("error writing to file")
	}

	return nil
}

func GetExpensesFromFile() ([]models.Expense, error) {
	data, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile(fileName, []byte{}, 0644)
		if err != nil {
			return nil, errors.New("error creating file: " + err.Error())
		}
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	if string(data) == "" {
		return nil, nil
	}

	var expenses []models.Expense
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, errors.New("error parsing file data")
	}

	return expenses, nil
}

func DeleteExpenseByID(id int) error {
	expenses, err := GetExpensesFromFile()
	if err != nil {
		return err
	}

	for i, e := range expenses {
		if e.ID == id {
			err := SaveExpenses(append(expenses[:i], expenses[i+1:]...))
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("expense not found")
}

func SummaryExpenses(month int) (float64, int, error) {
	expenses, err := GetExpensesFromFile()
	if err != nil {
		return 0, 0, err
	}

	if len(expenses) == 0 {
		return 0, 0, nil
	}

	total := 0.0
	totalCount := 0
	for _, expense := range expenses {
		expenseMonth, _ := time.Parse(time.RFC3339, expense.Date)
		if month != 0 && int(expenseMonth.Month()) != month {
			continue
		}

		total += expense.Price
		totalCount++
	}

	return total, totalCount, nil
}

func ClearExpenses() error {
	err := os.WriteFile(fileName, []byte{}, 0644)
	if err != nil {
		return errors.New("error clearing expenses")
	}

	return nil
}
