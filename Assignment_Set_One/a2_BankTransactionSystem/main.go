package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

const (
	DepositMoney       = 1
	WithdrawMoney      = 2
	ViewBalance        = 3
	TransactionHistory = 4
	Exit               = 5
)

func FindAccountByID(id int) (*Account, error) {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("error: Account not found")
}
func Deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than 0")
	}
	account.Balance += amount
	transaction := fmt.Sprintf("Deposited: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}
func Withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than 0")
	}
	if amount > account.Balance {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	transaction := fmt.Sprintf("Withdrawn: $%.2f", amount)
	account.TransactionHistory = append(account.TransactionHistory, transaction)
	return nil
}

func DisplayTransactionHistory(account *Account) {
	fmt.Println("Transaction History:")
	for _, transaction := range account.TransactionHistory {
		fmt.Println(transaction)
	}
}

func main() {
	accounts = append(accounts, Account{ID: 1, Name: "Luffy", Balance: 20000})
	accounts = append(accounts, Account{ID: 2, Name: "Zoro", Balance: 100})
	accounts = append(accounts, Account{ID: 3, Name: "Sanji", Balance: 1500})
	accounts = append(accounts, Account{ID: 4, Name: "Nami", Balance: 30000})
	accounts = append(accounts, Account{ID: 5, Name: "Robin", Balance: 25000})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		if choice == Exit {
			fmt.Println("Exiting.....")
			break
		}
		if choice < DepositMoney || choice > Exit {
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}
		var accountID int
		fmt.Print("Enter Account ID: ")
		fmt.Scan(&accountID)

		account, err := FindAccountByID(accountID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case DepositMoney:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := Deposit(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Deposited Successfully $%.2f. New Balance: %.2f\n", amount, account.Balance)
			}

		case WithdrawMoney:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := Withdraw(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Successfully withdrawn $%.2f. New Balance: %.2f\n", amount, account.Balance)
			}

		case ViewBalance:
			fmt.Printf("Account Balance: %.2f\n", account.Balance)

		case TransactionHistory:
			DisplayTransactionHistory(account)

			// default:
			// 	fmt.Println("Invalid choice.Please Select a different option.")
		}
	}
}
