package errs

import "errors"

// Generic
var Generic400Err = errors.New("Invalid request")
var Generic401Err = errors.New("Unauthorized")
var Generic422Err = errors.New("Invalid or missing request params")
var GenericInvalidParams = errors.New("Invalid or missing request param")

// Budget
var BudgetNotFound = errors.New("Budget not found")
var BudgetInvalidBudgetId = errors.New("Invalid Budget ID")
var BudgetInvalidTitle = errors.New("Invalid Budget Title")
var BudgetInvalidDescription = errors.New("Invalid Budget Description")

// Auth
var AuthInvalidHeader = errors.New("Invalid Auth Header")
var AuthIvalidPass = errors.New("Invalid Password or Email")
var AuthUnathorized = errors.New("Unauthorized")

// User
var UserNotFound = errors.New("User not found")
var UserInvalidEmail = errors.New("Invalid User email")
var UserInvalidUsername = errors.New("Invalid User username")
var UserInvalidPassword = errors.New("Invalid User Password")

// Transaction
var TransactionNotFound = errors.New("Transaction not found")
var TransactionInvalidAmount = errors.New("Invalid Transaction amount")
var TransactionInvalidDescription = errors.New("Invalid Transaction Description")
