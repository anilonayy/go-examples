package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/anilonayy/go-examples/expense-tracker/internal/enums/command"
	"github.com/anilonayy/go-examples/expense-tracker/internal/models"
	"github.com/anilonayy/go-examples/expense-tracker/internal/services"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen == 0 {
		fmt.Println("Please provide a command")
		printHelp()
		os.Exit(1)
	}

	action := command.ParseCommand(args[0])

	switch action {
	case command.CommandAdd:
		if err := handleAddExpense(args); err != nil {
			fmt.Println("Error adding expense: ", err)
		}
		break
	case command.CommandList:
		if err := handleListExpense(); err != nil {
			fmt.Println("Error listing expenses: ", err)
		}
		break
	case command.CommandDelete:
		if err := handleDeleteExpense(args); err != nil {
			fmt.Println("Error removing expense ", err)
		}
		break
	case command.CommandSummary:
		if err := handleSummaryExpense(args); err != nil {
			fmt.Println("Error getting summary ", err)
		}
		break
	case command.CommandClear:
		if err := handleClearExpenses(); err != nil {
			fmt.Println("Error clearing expenses ", err)
		}
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("add -desc <description> -price <price>")
	fmt.Println("list")
	fmt.Println("delete -id <id>")
	fmt.Println("summary -month <month>")
	fmt.Println("clear")
}

func handleAddExpense(args []string) error {
	cmd := flag.NewFlagSet(command.CommandAdd.ToString(), flag.ExitOnError)
	desc := cmd.String("desc", "", "Description of the expense")
	price := cmd.Float64("price", 0.0, "Price of the expense")

	err := cmd.Parse(args[1:])
	if err != nil {
		return err
	}

	expense := models.Expense{
		Description: *desc,
		Price:       *price,
		Date:        time.Now().Format(time.RFC3339),
	}

	if err := services.AddExpense(&expense); err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID:%d)\n", expense.ID)

	return nil
}

func handleListExpense() error {
	expenses, err := services.GetExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return nil
	}

	fmt.Println("ID\t\tDescription\tPrice\tDate")

	for _, expense := range expenses {
		fmt.Printf("%d\t%s\t\t%.2f\t%s\n", expense.ID, expense.Description, expense.Price, expense.Date)
	}

	return nil
}

func handleDeleteExpense(args []string) error {
	cmd := flag.NewFlagSet(command.CommandDelete.ToString(), flag.ExitOnError)
	id := cmd.Int("id", 0, "ID of the expense to delete")

	err := cmd.Parse(args[1:])
	if err != nil {
		return err
	}

	err = services.DeleteExpenseByID(*id)
	if err != nil {
		return err
	}

	fmt.Printf("Expense with ID %d deleted successfully\n", *id)

	return nil
}

func handleSummaryExpense(args []string) error {
	cmd := flag.NewFlagSet(command.CommandDelete.ToString(), flag.ExitOnError)
	month := cmd.Int("month", 0, "ID of the expense to delete")

	err := cmd.Parse(args[1:])
	if err != nil {
		return err
	}

	if *month != 0 && (*month < 1 || *month > 12) {
		return fmt.Errorf("invalid month: %d", *month)
	}

	total, count, err := services.SummaryExpenses(*month)

	fmt.Printf("Total: $%.0f, Count: %d\n", total, count)
	return nil
}

func handleClearExpenses() error {
	err := services.ClearExpenses()
	if err != nil {
		return err
	}

	fmt.Println("Expenses cleared successfully")

	return nil
}
