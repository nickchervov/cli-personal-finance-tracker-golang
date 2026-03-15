package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"personal_finance_tracker/internal/finance"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func input(title string) string {
	fmt.Print(title)

	scanner.Scan()

	return scanner.Text()
}

func addTransactions(tracker *finance.FinanceTracker) {
	convertedAmount := 0
	var err error
	for {
		amount := input("Введите сумму:")
		if amount == "" {
			fmt.Println("Сумма не может быть пустой!")
			continue
		}
		convertedAmount, err = strconv.Atoi(amount)
		if err == nil {
			break
		}
		fmt.Println("Сумма должна быть в виде числа!")
	}

	transactionType := ""
	for {
		transactionType = input("Введите тип (доход/расход):")
		if strings.ToLower(transactionType) == "доход" || strings.ToLower(transactionType) == "расход" {
			break
		}
		fmt.Println("Необходимо ввести либо доход либо расход!")
	}

	category := ""
	for {
		category = input("Введите категорию:")
		if category != "" {
			break
		}
		fmt.Println("Необходимо ввести категорию!")
	}

	desc := input("Введите описание (можно оставить пустым):")

	tracker.AddTransaction(convertedAmount, transactionType, category, desc)

	fmt.Println("Транзакция добавлена!")
}

func getAllTransaction(tracker *finance.FinanceTracker) {
	outputTransactions := `Номер транзакции: %d
Сумма: %d
Тип: %s
Категория: %s
Описание: %s
Дата: %s
`
	transactions := tracker.GetAllTransactions()

	for i, transaction := range transactions {
		transactionDate := transaction.Date.Format("02.01.2006 15:04:05")
		fmt.Printf(outputTransactions, i+1, transaction.Amount, transaction.Type, transaction.Category, transaction.Description, transactionDate)
		fmt.Println()
	}
}

func getStatistic(tracker *finance.FinanceTracker) {
	fmt.Println("--- Статистика ---")

	transactions := tracker.GetAllTransactions()

	balance := 0
	income := 0
	expense := 0

	for i, transaction := range transactions {
		convertedType := strings.ToLower(transaction.Type)
		switch convertedType {
		case "доход":
			balance += transaction.Amount
			income += transaction.Amount
		case "расход":
			balance -= transaction.Amount
			expense += transaction.Amount
		default:
			log.Printf("транзакция под номером: %d имеет некорректный тип", i+1)
			continue
		}
	}

	fmt.Println("Общий баланс:", balance)
	fmt.Println("Доходы:", income)
	fmt.Println("Расходы:", expense)
}

func runProgramm(tracker *finance.FinanceTracker) bool {

	intro := `Выберите действие:
1. Добавить транзакцию
2. Показать все транзакции
3. Показать статистику
4. Выйти`

	fmt.Println(intro)

	choise := input("")
	convertedChoise, err := strconv.Atoi(choise)
	if err != nil {
		log.Println(err)
		return false
	}

	switch convertedChoise {
	case 1:
		fmt.Println()
		addTransactions(tracker)
		fmt.Println()

	case 2:
		fmt.Println()
		getAllTransaction(tracker)
		fmt.Println()
	case 3:
		fmt.Println()
		getStatistic(tracker)
		fmt.Println()

	case 4:
		fmt.Println()
		err = tracker.SaveToFile("transactions.json")
		if err != nil {
			log.Println(err)
		}
		fmt.Println("До свидания!")
		return false
	default:
		fmt.Println("Введено некорректное значение.")
	}
	return true
}
func main() {
	var tracker finance.FinanceTracker

	fmt.Println("Привет! Это трекер личных финансов.")
	fmt.Println()

	tracker.LoadFromFile("transactions.json")

	for runProgramm(&tracker) {
	}
}
