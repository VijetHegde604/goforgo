// error_custom.go
// Learn how to create custom error types and implement the error interface

package main

import (
	"errors"
	"fmt"
	"time"
)

// Define a simple custom error type
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// Implement the error interface
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %v\n%s", e.Field, e.Value, e.Message)
}

// Define a more complex error type with multiple fields
type NetworkError struct {
	Operation string
	URL       string
	Timestamp time.Time
	Err       error // Wrapped error
}

func (e NetworkError) Error() string {
	return fmt.Sprintf(
		"Something went wrong with the network at %s during %s at %s\nErr: %v",
		e.URL,
		e.Operation,
		e.Timestamp.Format("2006-01-02 at 15:04"),
		e.Err,
	)
}

// Implement error unwrapping
func (e NetworkError) Unwrap() error {
	return e.Err
}

// Define sentinel errors
var (
	ErrInsufficientFunds = errors.New("insufficient balance")
	ErrAccountNotFound   = errors.New("account not found")
	ErrInvalidAmount     = errors.New("invalid amount")
)

// Custom error with additional methods
type AccountError struct {
	AccountID string
	Balance   float64
	Operation string
	Code      int
}

func (e AccountError) Error() string {
	return fmt.Sprintf(
		"account error [%d]: %s failed for account %s (balance: %.2f)",
		e.Code,
		e.Operation,
		e.AccountID,
		e.Balance,
	)
}

// Add method to check error severity
func (e AccountError) IsCritical() bool {
	return e.Code >= 500
}

// Add method to get user-friendly message
func (e AccountError) UserMessage() string {
	switch e.Code {
	case 404:
		return "Account not found"
	case 400:
		return "Invalid request"
	case 500:
		return "Internal server error"
	default:
		return "An error occurred"
	}
}

// Functions that return custom errors
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age cannot be negative",
		}
	}

	if age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "age cannot exceed 150",
		}
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "email cannot be empty",
		}
	}

	if !contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Value:   email,
			Message: "email must contain @",
		}
	}

	return nil
}

func fetchUserData(url string) error {
	// Simulate network operation that might fail
	if url == "" {
		return NetworkError{
			Operation: "GET",
			URL:       url,
			Timestamp: time.Now(),
			Err:       errors.New("URL cannot be empty"),
		}
	}

	if url == "timeout.com" {
		return NetworkError{
			Operation: "GET",
			URL:       url,
			Timestamp: time.Now(),
			Err:       errors.New("request timed out"),
		}
	}

	return nil
}

func withdraw(accountID string, amount float64) error {
	// Simulate account operations
	balance := 100.0

	if accountID == "" {
		return ErrAccountNotFound
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > balance {
		return ErrInsufficientFunds
	}

	// Simulate system error
	if amount > 1000 {
		return AccountError{
			AccountID: accountID,
			Balance:   balance,
			Operation: "withdraw",
			Code:      500,
		}
	}

	return nil
}

// Helper function for string contains
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("=== Custom Error Types ===")

	// Test validation errors
	testAges := []int{25, -5, 200, 0}

	for _, age := range testAges {
		err := validateAge(age)

		if err != nil {
			fmt.Printf("Age validation error: %v\n", err)

			// Extract ValidationError
			var valErr ValidationError

			if errors.As(err, &valErr) {
				fmt.Printf(
					"  Field: %s, Value: %v\n",
					valErr.Field,
					valErr.Value,
				)
			}
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}

	fmt.Println("\n=== Email Validation ===")

	testEmails := []string{
		"user@example.com",
		"",
		"invalid-email",
		"test@domain.org",
	}

	for _, email := range testEmails {
		err := validateEmail(email)

		if err != nil {
			fmt.Printf("Email validation error: %v\n", err)
		} else {
			fmt.Printf("Email '%s' is valid\n", email)
		}
	}

	fmt.Println("\n=== NetworkError with Wrapping ===")

	testURLs := []string{
		"https://api.example.com",
		"",
		"timeout.com",
		"https://valid.com",
	}

	for _, url := range testURLs {
		err := fetchUserData(url)

		if err != nil {
			fmt.Printf("Network error: %v\n", err)

			// Extract NetworkError
			var netErr NetworkError

			if errors.As(err, &netErr) {
				fmt.Printf(
					"  Operation: %s, URL: %s, Time: %s\n",
					netErr.Operation,
					netErr.URL,
					netErr.Timestamp.Format("15:04:05"),
				)

				// Unwrap the underlying error
				wrappedErr := errors.Unwrap(err)

				if wrappedErr != nil {
					fmt.Printf("  Wrapped error: %v\n", wrappedErr)
				}
			}
		} else {
			fmt.Printf("Successfully fetched data from %s\n", url)
		}
	}

	fmt.Println("\n=== Sentinel Errors ===")

	// Test account operations with sentinel errors
	testOperations := []struct {
		accountID string
		amount    float64
	}{
		{"ACC123", 50.0},
		{"", 25.0},
		{"ACC456", -10.0},
		{"ACC789", 150.0},
		{"ACC999", 2000.0},
	}

	for _, op := range testOperations {
		err := withdraw(op.accountID, op.amount)

		if err != nil {
			fmt.Printf("Withdraw error: %v\n", err)

			// Check for specific sentinel errors
			switch err {
			case ErrAccountNotFound:
				fmt.Println("  -> Please check the account ID")

			case ErrInvalidAmount:
				fmt.Println("  -> Amount must be positive")

			case ErrInsufficientFunds:
				fmt.Println("  -> Please add funds to your account")

			default:
				// Check for AccountError type
				var accErr AccountError

				if errors.As(err, &accErr) {
					fmt.Printf(
						"  -> User message: %s\n",
						accErr.UserMessage(),
					)

					fmt.Printf(
						"  -> Critical: %t\n",
						accErr.IsCritical(),
					)
				}
			}
		} else {
			fmt.Printf(
				"Successfully withdrew %.2f from %s\n",
				op.amount,
				op.accountID,
			)
		}
	}

	fmt.Println("\n=== Error Type Checking ===")

	// Create different types of errors
	testErrors := []error{
		ValidationError{
			Field:   "name",
			Value:   "",
			Message: "name is required",
		},
		NetworkError{
			Operation: "GET",
			URL:       "api.test.com",
			Timestamp: time.Now(),
			Err:       errors.New("connection failed"),
		},
		ErrInsufficientFunds,
		AccountError{
			AccountID: "TEST",
			Balance:   50.0,
			Operation: "transfer",
			Code:      400,
		},
		fmt.Errorf("generic error"),
	}

	for i, err := range testErrors {
		fmt.Printf("Error %d: %v\n", i+1, err)

		// Use type switch to handle different error types
		switch e := err.(type) {
		case ValidationError:
			fmt.Printf(
				"  -> Validation error on field: %s\n",
				e.Field,
			)

		case NetworkError:
			fmt.Printf(
				"  -> Network error during: %s\n",
				e.Operation,
			)

		case AccountError:
			fmt.Printf(
				"  -> Account error (code %d): %s\n",
				e.Code,
				e.UserMessage(),
			)

		default:
			fmt.Printf(
				"  -> Generic error type: %T\n",
				e,
			)
		}

		fmt.Println()
	}
}
