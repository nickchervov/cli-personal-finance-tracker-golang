package finance

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Transaction struct {
	Amount      int
	Type        string
	Category    string
	Description string
	Date        time.Time
}

type FinanceTracker struct {
	Transactions []Transaction
}

func (ft *FinanceTracker) SaveToFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal()
		}
	}()

	encoder := json.NewEncoder(file)

	return encoder.Encode(ft.Transactions)
}

func (ft *FinanceTracker) LoadFromFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal()
		}
	}()

	decoder := json.NewDecoder(file)

	return decoder.Decode(&ft.Transactions)
}

func (ft *FinanceTracker) AddTransaction(amount int, transactionType, category, desc string) error {
	transaction := Transaction{
		Amount:      amount,
		Type:        transactionType,
		Category:    category,
		Description: desc,
		Date:        time.Now(),
	}

	ft.Transactions = append(ft.Transactions, transaction)

	return nil
}

func (ft *FinanceTracker) GetAllTransactions() []Transaction {
	return ft.Transactions
}

func (ft *FinanceTracker) GetTransactionsByCategory(category string) []Transaction {
	if category == "" {
		fmt.Println("Необходимо ввести категорию!")
		return nil
	}
	transByCategory := []Transaction{}

	for _, trans := range ft.Transactions {
		if strings.EqualFold(trans.Category, category) {
			transByCategory = append(transByCategory, trans)
		}
	}

	if len(transByCategory) == 0 {
		fmt.Println("Транзакций по данной категории не найдено!")
	}
	return transByCategory
}

func (ft *FinanceTracker) GetMonthlySummary(month int) map[string]int {
	if month < 1 || month > 12 {
		fmt.Println("Такого месяца не существует!")
		return nil
	}
	result := make(map[string]int)

	for _, trans := range ft.Transactions {
		if trans.Date.Month() != time.Month(month) {
			continue
		}

		var label string
		if strings.ToLower(trans.Type) == "расход" {
			label = fmt.Sprintf("Расход_%s", trans.Date.Format("01.02.2006"))
		} else if strings.ToLower(trans.Type) == "доход" {
			label = fmt.Sprintf("Доход_%s", trans.Date.Format("01.02.2006"))
		} else {
			continue
		}
		result[label] += trans.Amount
	}

	if len(result) == 0 {
		fmt.Println("В данном месяце нет доходов или расходов.")
	}
	return result
}

func (ft *FinanceTracker) FindLargestExpense() (*Transaction, error) {
	largestExpense := (*Transaction)(nil)
	errNoExpense := errors.New("нет расходов")

	for _, trans := range ft.Transactions {
		if strings.ToLower(trans.Type) == "доход" {
			continue
		}
		if largestExpense == nil || trans.Amount > largestExpense.Amount {
			largestExpense = &trans
		}
	}

	if largestExpense != nil {
		return largestExpense, nil
	}
	return nil, errNoExpense
}

func (ft *FinanceTracker) GetBalance() int {
	balance := 0

	for _, trans := range ft.Transactions {
		balance += trans.Amount
	}

	return balance
}

// расходы
func (ft *FinanceTracker) GetTotalExpense() int {
	expense := 0

	for _, trans := range ft.Transactions {
		if strings.ToLower(trans.Type) != "расход" {
			continue
		}
		expense += trans.Amount
	}

	return expense
}

// доходы
func (ft *FinanceTracker) GetTotalIncome() int {
	income := 0

	for _, trans := range ft.Transactions {
		if trans.Type != "Доходы" {
			continue
		}
		income += trans.Amount
	}

	return income
}
