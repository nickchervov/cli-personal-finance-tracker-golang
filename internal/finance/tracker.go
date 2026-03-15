package finance

import (
	"encoding/json"
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
