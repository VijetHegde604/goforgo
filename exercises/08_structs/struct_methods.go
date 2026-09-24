// struct_methods.go
// Learn how to add methods to structs in Go

package main

import (
	"fmt"
	"math"
)

// Define a Rectangle struct with Width and Height fields
type Rectangle struct {
	// Define fields here
	Width  int
	Height int
}

// Define an Area method for Rectangle
// Hint: func (receiver ReceiverType) MethodName() ReturnType
func (r Rectangle) Area() float64 {
	// Calculate and return area
	return float64(r.Width) * float64(r.Height)
}

// Define a Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	// Calculate and return perimeter
	return 2 * float64(r.Width+r.Height)
}

// Define a Circle struct with Radius field
type Circle struct {
	// Define field here
	Radius int
}

// Define an Area method for Circle
func (c Circle) Area() float64 {
	// Calculate and return area (π * r²)
	return math.Pi * (float64(c.Radius) * float64(c.Radius))
}

// Define a BankAccount struct
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

// Define a Deposit method that increases the balance
// This should use a pointer receiver to modify the original struct
func (ba *BankAccount) Deposit(amount float64) {
	// Add amount to balance
	ba.Balance += amount
}

// Define a Withdraw method that decreases the balance
// Return true if successful, false if insufficient funds
func (ba *BankAccount) Withdraw(amount float64) bool {
	// Check if sufficient funds, then withdraw
	if ba.Balance < amount {
		return false
	}
	ba.Balance -= amount
	return true
}

// Define a GetBalance method (value receiver is fine here)
func (ba BankAccount) GetBalance() float64 {
	// Return current balance
	return ba.Balance
}

// Define a String method to customize how BankAccount prints
// This implements the fmt.Stringer interface
func (ba BankAccount) String() string {
	// Return formatted string representation
	return fmt.Sprintf("Account: %s, Balance: %.2f", ba.AccountNumber, ba.Balance)

}

func main() {
	// Create a Rectangle and call its methods
	rect := Rectangle{5, 3} // Create Rectangle with Width: 5, Height: 3

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// Create a Circle and call its method
	circle := Circle{4} // Create Circle with Radius: 4

	fmt.Printf("Circle area: %.2f\n", circle.Area())

	// Create a BankAccount and perform operations
	account := BankAccount{"12345", 1000.0} // Create BankAccount with AccountNumber: "12345", Balance: 1000.0

	fmt.Println("Initial account:", account) // This will use our String method

	// Make a deposit
	// Deposit $250
	account.Deposit(250)

	fmt.Printf("After deposit: Balance = %.2f\n", account.GetBalance())

	// Try to withdraw money
	success := account.Withdraw(500) // Try to withdraw $500
	if success {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient funds")
	}

	// Try to withdraw too much money
	success = account.Withdraw(2000) // Try to withdraw $2000
	if !success {
		fmt.Println("Cannot withdraw $2000 - insufficient funds")
	}

	fmt.Println("Final account:", account)
}
